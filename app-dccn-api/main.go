package main

import (
	"context"
	"errors"
	"log"
	"os"

	grpc "github.com/micro/go-grpc"
	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"

	ankr_default "github.com/Ankr-network/dccn-common/protos"
	common_proto "github.com/Ankr-network/dccn-common/protos/common"
	dcmgr "github.com/Ankr-network/dccn-common/protos/dcmgr/v1/micro"
	mail "github.com/Ankr-network/dccn-common/protos/email/v1/micro"
	taskmgr "github.com/Ankr-network/dccn-common/protos/taskmgr/v1/micro"
	"github.com/Ankr-network/dccn-hub/app-dccn-api/apihandler"

	dbservice "github.com/Ankr-network/dccn-hub/app-dccn-dcmgr/db_service"
	"github.com/Ankr-network/dccn-hub/app-dccn-dcmgr/subscriber"

	usermgr "github.com/Ankr-network/dccn-common/protos/usermgr/v1/micro"

	"github.com/Ankr-network/dccn-hub/app-dccn-dcmgr/handler"
	"github.com/Ankr-network/dccn-hub/app-dccn-usermgr/config"

	_ "github.com/micro/go-plugins/broker/rabbitmq"
)

var (
	srv        micro.Service
	conf       config.Config
	db         dbservice.DBService
	err        error
	authList   map[string]struct{}
	userClient *apihandler.ApiUser
	taskClient *apihandler.ApiTask
)

func main() {
	Init()

	if db, err = dbservice.New(conf.DB); err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	startHandler()
}

// Init starts handler to listen.
func Init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	if conf, err = config.Load(); err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("Load config %+v\n", conf)

	authList = map[string]struct{}{
		"TaskMgr.CreateTask": struct{}{},
		"TaskMgr.TaskList":   struct{}{},
		"TaskMgr.CancelTask": struct{}{},
		"TaskMgr.PurgeTask":  struct{}{},
		"TaskMgr.TaskDetail": struct{}{},
		"TaskMgr.UpdateTask": struct{}{},
	}
}

func startHandler() {
	// New Service
	srv = grpc.NewService(
		micro.WrapHandler(AuthWrapper),
	)

	// Initialise service
	srv.Init()

	// Register User Handler
	userClient = apihandler.NewApiUser(srv.Client())
	if err := usermgr.RegisterUserMgrHandler(srv.Server(), userClient); err != nil {
		log.Fatal(err.Error())
	}

	// Register Task Handler
	taskClient = apihandler.NewApiTask(srv.Client())
	if err := taskmgr.RegisterTaskMgrHandler(srv.Server(), taskClient); err != nil {
		log.Fatal(err.Error())
	}

	// Dc Manager register handler
	// New Publisher to deploy new task action.
	taskFeedback := micro.NewPublisher(ankr_default.MQFeedbackTask, srv.Client())

	dcHandler := handler.New(db, taskFeedback)

	// Register Function as TaskStatusFeedback to update task by data center manager's feedback.
	if err := micro.RegisterSubscriber(ankr_default.MQDeployTask, srv.Server(), subscriber.New(dcHandler.DcStreamCaches)); err != nil {
		log.Fatal(err.Error())
	}

	// Register Dc Manager Handler
	if err := dcmgr.RegisterDCStreamerHandler(srv.Server(), dcHandler); err != nil {
		log.Fatal(err.Error())
	}

	defer dcHandler.Cleanup()

	// Register Email Handler
	if err := mail.RegisterMailHandler(srv.Server(), apihandler.NewApiEmail(srv.Client())); err != nil {
		log.Fatal(err.Error())
	}

	// Run srv
	if err := srv.Run(); err != nil {
		log.Println(err.Error())
	}
}

func needAuth(method string) bool {
	_, ok := authList[method]
	return ok
}

func AuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		meta, ok := metadata.FromContext(ctx)
		// Note this is now uppercase (not entirely sure why this is...)
		var token string
		if ok {
			token = meta["token"]
		}

		if os.Getenv("DISABLE_AUTH") == "true" || !needAuth(req.Method()) {
			if ok && token != "" {
				if err := userClient.RefreshToken(ctx, &usermgr.Token{Token: token}, &common_proto.Error{}); err != nil {
					log.Println(err.Error())
					return err
				}
			}
			return fn(ctx, req, resp)
		}

		if !ok {
			log.Println("no auth meta-data found in request")
			return errors.New("no auth meta-data found in request")
		}

		log.Println("Authenticating with token: ", token)
		// Auth here
		// Really shouldn't be using a global here, find a better way
		// of doing this, since you can't pass it into a wrapper.
		userMgrService := usermgr.NewUserMgrService(ankr_default.UserMgrRegistryServerName, srv.Client())
		if _, err := userMgrService.VerifyAndRefreshToken(context.Background(), &usermgr.Token{Token: token}); err != nil {
			log.Println(err.Error())
			return err
		}

		return fn(ctx, req, resp)
	}
}
