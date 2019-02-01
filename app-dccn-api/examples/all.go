package main

import (
	"context"
	"log"
	"time"

	taskmgr "github.com/Ankr-network/dccn-common/protos/taskmgr/v1/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	usermgr "github.com/Ankr-network/dccn-common/protos/usermgr/v1/grpc"

	common_proto "github.com/Ankr-network/dccn-common/protos/common"
	apiCommon "github.com/Ankr-network/dccn-hub/app-dccn-api/examples/common"
)

// var addr = "localhost:50051"

var addr = "client-dev.dccn.ankr.network:50051"

// func init() {
// 	addr = os.Getenv("API_ADDRESS")
// 	log.Println("Get Addr: ", addr)
// }

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err.Error())
	}
	defer func(conn *grpc.ClientConn) {
		if err := conn.Close(); err != nil {
			log.Println(err.Error())
		}
	}(conn)

	taskClient := taskmgr.NewTaskMgrClient(conn)
	userClient := usermgr.NewUserMgrClient(conn)

	user := &usermgr.User{
		Name:     "user_test",
		Nickname: "test",
		Email:    `123@Gmail.com`,
		Password: "1234567890",
		Balance:  99,
	}
	if _, err := userClient.Register(context.Background(), user); err != nil {
		log.Fatal(err.Error())
	} else {
		log.Println("Register Ok")
	}

	var token string
	var userId string
	if rsp, err := userClient.Login(context.TODO(), &usermgr.LoginRequest{Email: user.Email, Password: user.Password}); err != nil {
		log.Fatal(err.Error())
	} else {
		log.Printf("login Success: %s\n", rsp.Token)
		token = rsp.Token
		userId = rsp.UserId
	}

	md := metadata.New(map[string]string{
		"token": token,
	})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	tokenContext, cancel := context.WithTimeout(ctx, 180*time.Second)
	defer cancel()

	task := apiCommon.MockTasks()[0]
	log.Println("Test CreateTask")
	if rsp, err := taskClient.CreateTask(tokenContext, &taskmgr.CreateTaskRequest{UserId: userId, Task: &task}); err != nil {
		log.Fatal(err.Error())
	} else {
		task.Id = rsp.TaskId
		log.Println("CreateTask ok", *rsp)
	}

	userTasks := make([]*common_proto.Task, 0)
	if rsp, err := taskClient.TaskList(tokenContext, &taskmgr.ID{UserId: userId}); err != nil {
		log.Fatal(err.Error())
	} else {
		userTasks = append(userTasks, rsp.Tasks...)
		if len(userTasks) == 0 {
			log.Fatalf("no tasks belongs to %s", userId)
		} else {
			log.Println(len(userTasks), "tasks belongs to ", userId)
			log.Println(userTasks[0])
		}
	}

	log.Println("Test CancelTask")
	if _, err := taskClient.CancelTask(tokenContext, &taskmgr.Request{UserId: userId, TaskId: task.Id}); err != nil {
		log.Fatal(err.Error())
	} else {
		log.Println("CancelTask Ok")

		// Verify Canceled task
		log.Println("Test TaskDetail")
		if _, err := taskClient.TaskDetail(tokenContext, &taskmgr.Request{UserId: userId, TaskId: task.Id}); err != nil {
			log.Fatal(err.Error())
		} else {
			log.Println("TaskDetail Ok")
		}
	}

	if _, err := userClient.Logout(tokenContext, &usermgr.LogoutRequest{}); err != nil {
		log.Fatal(err.Error())
	} else {
		log.Println("Logout ok")
	}

	// UpdateTask
	task.Name = "updateTask"
	if _, err := taskClient.UpdateTask(tokenContext, &taskmgr.UpdateTaskRequest{UserId: userId, Task: &task}); err != nil {
		log.Println(err.Error())
	}

	log.Println("END")
}
