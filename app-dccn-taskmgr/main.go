package main

import (
	"log"

	grpc "github.com/micro/go-grpc"
	micro "github.com/micro/go-micro"

	ankr_default "github.com/Ankr-network/dccn-common/protos"
	pb "github.com/Ankr-network/dccn-common/protos/taskmgr/v1/micro"
	"github.com/Ankr-network/dccn-hub/app-dccn-taskmgr/config"
	dbservice "github.com/Ankr-network/dccn-hub/app-dccn-taskmgr/db_service"
	"github.com/Ankr-network/dccn-hub/app-dccn-taskmgr/handler"
	"github.com/Ankr-network/dccn-hub/app-dccn-taskmgr/subscriber"

	_ "github.com/micro/go-plugins/broker/rabbitmq"
)

var (
	conf config.Config
	db   dbservice.DBService
	err  error
)

func main() {
	Init()

	if db, err = dbservice.New(conf.DB); err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	startHandler(db)
}

// Init starts handler to listen.
func Init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	if conf, err = config.Load(); err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("load config: %#v\n", conf)
}

// StartHandler starts handler to listen.
func startHandler(db dbservice.DBService) {
	// New Service
	srv := grpc.NewService(
		micro.Name(ankr_default.TaskMgrRegistryServerName),
		// micro.WrapHandler(wrapper.AuthWrapper),
	)

	// Initialise srv
	srv.Init()

	// New Publisher to deploy new task action.
	deployTask := micro.NewPublisher(ankr_default.MQDeployTask, srv.Client())

	// Register Function as TaskStatusFeedback to update task by data center manager's feedback.
	if err := micro.RegisterSubscriber(ankr_default.MQFeedbackTask, srv.Server(), subscriber.New(db)); err != nil {
		log.Fatal(err.Error())
	}

	// Register Handler
	if err := pb.RegisterTaskMgrHandler(srv.Server(), handler.New(db, deployTask)); err != nil {
		log.Fatal(err.Error())
	}

	// Run srv
	if err := srv.Run(); err != nil {
		log.Println(err.Error())
	}
}