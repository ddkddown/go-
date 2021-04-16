package mr

import (
	"errors"
	"fmt"
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
	srcs      SrcFiles    //源文件
	middles   MiddleFiles //中间文件
	workers   WorkerInfo  //workers名称
	workerMu  sync.Mutex  //mutex
	middlesMu sync.Mutex
	srcsMu    sync.Mutex
}

//Regist worker regist
func (p *Master) Regist(request RegistReq, i *int) error {
	p.workerMu.Lock()
	defer p.workerMu.Unlock()

	if _, exist := p.workers[request.WorkerName]; exist {
		return errors.New("worker already registed!")
	}

	p.workers[request.WorkerName] = 1
	fmt.Println("regist worker: ", request.WorkerName)
	return nil
}

//AskSrcTask worker fetch task
func (p *Master) AskSrcTask(request AskSrcTaskReq, response *AskSrcTaskRsp) error {
	p.srcsMu.Lock()
	defer p.srcsMu.Unlock()

	/*
		if _, exist := p.workers[request.workerName]; !exist {
			return -1, errors.New("worker not registed!")
		}
	*/
	fmt.Println("0")
	for k, v := range p.srcs {
		if v.flag {
			continue
		}

		response.FileName = k
		v.flag = true
		v.worker = request.WorkerName
		p.srcs[k] = v
		fmt.Println("1")
		return nil
	}

	fmt.Println("2")
	return errors.New("no more task")
}

//AskMiddleTask worker fetch task
func (p *Master) AskMiddleTask(request AskMiddleTaskReq, response *AskMiddleTaskRsp) error {
	p.middlesMu.Lock()
	defer p.middlesMu.Unlock()

	for k, v := range p.middles {
		if v.flag {
			continue
		}

		response.FileName = k
		v.flag = true
		v.worker = request.WorkerName
		p.middles[k] = v
		return nil
	}

	return errors.New("no more task")
}

func (p *Master) FinishSrc(request FinishSrcReq, i *int) error {
	p.srcsMu.Lock()
	delete(p.srcs, request.SrcName)
	p.srcsMu.Unlock()

	var tmp check
	tmp.flag = false
	tmp.worker = ""

	p.middlesMu.Lock()
	p.middles[request.MiddleName] = tmp
	p.middlesMu.Unlock()
	return nil
}

func (p *Master) FinishMiddle(request FinishMiddleReq, i *int) error {
	p.middlesMu.Lock()
	defer p.middlesMu.Unlock()

	delete(p.middles, request.FileName)
	return nil
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
	if 0 == len(m.srcs) && 0 == len(m.middles) {
		ret = true
		return ret
	}

	m.workerMu.Lock()
	defer m.workerMu.Unlock()

	for w, c := range m.workers {

		c++
		if c >= 10 {
			// worker down

			m.srcsMu.Lock()
			for k, v := range m.srcs {
				if v.worker == w {
					v.flag = false
					v.worker = ""
					m.srcs[k] = v
				}
			}
			m.srcsMu.Unlock()

			m.middlesMu.Lock()
			for k, v := range m.middles {
				if v.worker == w {
					v.flag = false
					v.worker = ""
					m.middles[k] = v
				}
			}
			m.middlesMu.Unlock()

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

	for _, file := range files {
		var tmp check
		tmp.flag = false
		tmp.worker = ""
		m.srcs[file] = tmp
	}

	m.server()
	return &m
}
