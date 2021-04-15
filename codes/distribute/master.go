package mr

import (
	"errors"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
)

type SrcFiles map[string]bool
type MiddleFiles map[string]bool
type WorkerInfo map[string]int

type Master struct {
	// Your definitions here.
	srcs    SrcFiles    //源文件
	middles MiddleFiles //中间文件
	workers WorkerInfo  //workers名称
}

func (p *Master) Regist(request RegistReq) error {
	if _, exist := p.workers[request.workerName]; exist {
		return errors.New("worker already registed!")
	}

	p.workers[request.workerName] = 1
	return nil
}

func (p *Master) AskTask(request AskTaskReq, response *AskTaskRsp) int, error {
	if _, exist := p.workers[request.workerName]; !exist {
		return -1, errors.New("worker not registed!")
	}

	for k, v := range p.srcs {
		if v {
			continue
		}

		response.fileName = k
		return 0, nil
	}

	return 0, errors.New("no more srcs")
}

// Your code here -- RPC handlers for the worker to call.

//
// an example RPC handler.
//
// the RPC argument and reply types are defined in rpc.go.
//
func (m *Master) Example(args *ExampleArgs, reply *ExampleReply) error {
	reply.Y = args.X + 1
	return nil
}

//
// start a thread that listens for RPCs from worker.go
//
func (m *Master) server() {
	rpc.Register(m)
	rpc.HandleHTTP()
	//l, e := net.Listen("tcp", ":1234")
	sockname := masterSock()
	os.Remove(sockname)
	l, e := net.Listen("unix", sockname)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
}

//
// main/mrmaster.go calls Done() periodically to find out
// if the entire job has finished.
//
func (m *Master) Done() bool {
	ret := false

	// Your code here.

	return ret
}

//
// create a Master.
// main/mrmaster.go calls this function.
// nReduce is the number of reduce tasks to use.
//
func MakeMaster(files []string, nReduce int) *Master {
	m := Master{}

	// Your code here.

	m.server()
	return &m
}
