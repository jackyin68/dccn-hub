package taskmgr

import (
	"fmt"
	"os"

	ankr_const "github.com/Ankr-network/dccn-common"
	"github.com/Ankr-network/dccn-hub/util"
)

type server struct {
	name string
}

//Listen RambbitMQ queue (ankr_const.TaskManagerQueueName), handle the event from api_listener, router the event to DataCetner
func (s server) Handle(e util.Event) {
	logStr := fmt.Sprintf("this is handle message %s\n", e.Type)
	util.WriteLog(logStr)
	task := util.GetTask(e.TaskID)

	// add router
	logStr2 := fmt.Sprintf("task %d in DataCenter id %d  datacenter name %s  \n", task.ID, int(task.Datacenterid), task.Datacenter)
	util.WriteLog(logStr2)

	util.Send(ankr_const.K8sAdopterQueueName, e)
}

func StartService() {
	if len(os.Args) == 3 {
		util.MongoDBHost = os.Args[1]

	}

	if len(os.Args) == 3 {
		util.RabbitMQHost = os.Args[2]
	}
	ss := server{}
	util.Receive(ankr_const.TaskManagerQueueName, ss)
}
