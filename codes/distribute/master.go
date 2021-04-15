package mr

import (
	"errors"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"sync"
)

type check struct {
	flag   bool
	worker string
}

type SrcFiles map[string]check
type MiddleFiles map[string]check
type WorkerInfo map[string]int

type Master struct {
	// Your definitions here.
	srcs    SrcFiles    //源文件
	middles MiddleFiles //中间文件
	workers WorkerInfo  //workers名称
	mu      sync.Mutex  //mutex
}

//Regist worker regist
func (p *Master) Regist(request RegistReq) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if _, exist := p.workers[request.workerName]; exist {
		return errors.New("worker already registed!")
	}

	p.workers[request.workerName] = 1
	return nil
}

//AskTask worker fetch task
func (p *Master) AskTask(request AskTaskReq, response *AskTaskRsp) (int, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if _, exist := p.workers[request.workerName]; !exist {
		return -1, errors.New("worker not registed!")
	}

	for k, v := range p.srcs {
		if v.flag {
			continue
		}

		response.fileName = k
		v.flag = true
		v.worker = request.workerName
		p.srcs[k] = v
		return 0, nil
	}

	for k, v := range p.middles {
		if v.flag {
			continue
		}

		response.fileName = k
		v.flag = true
		v.worker = request.workerName
		p.middles[k] = v
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
	m.mu.Lock()
	defer m.mu.Unlock()

	ret := false

	// Your code here.
	if 0 == len(m.srcs) && 0 == len(m.middles) {
		ret = true
		return ret
	}

	for w, c := range m.workers {
		c++
		if c >= 10 {
			// worker down
			for k, v := range m.srcs {
				if v.worker == w {
					v.flag = false
					v.worker = ""
					m.srcs[k] = v
				}
			}

			for k, v := range m.middles {
				if v.worker == w {
					v.flag = false
					v.worker = ""
					m.middles[k] = v
				}
			}
		} else {
			m.workers[w] = c
		}
	}
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
	m.srcs = make(map[string]check)
	m.middles = make(map[string]check)
	m.workers = make(map[string]int)

	m.server()
	return &m
}
