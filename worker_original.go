// Original code with Dispatcher
package main

import (
	_ "expvar"
	"flag"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

// Job 拥有执行工作单元所需的属性
type Job struct {
	Name  string
	Delay time.Duration //  Job 处理所需要的时间
}

// use id mark worker
func NewWorker(id int, workerPool chan chan Job) Worker {
	return Worker{
		id:         id,
		jobQueue:   make(chan Job),
		workerPool: workerPool,
		quitChan:   make(chan bool),
	}
}

type Worker struct {
	id         int
	jobQueue   chan Job
	workerPool chan chan Job
	quitChan   chan bool
}

func (w Worker) start() {
	go func() {
		for {
			w.workerPool <- w.jobQueue

			select {
			case job := <-w.jobQueue:
				// Dispatcher has added a job to worker jobQueue
				fmt.Printf("worker%d: started %s, blocking for %f seconds\n", w.id, job.Name, job.Delay.Seconds())
				time.Sleep(job.Delay)
				fmt.Printf("worker%d: completed %s!\n", w.id, job.Name)
			case <-w.quitChan:
				fmt.Printf("worker%d stopping\n", w.id)
				return
			}
		}
	}()
}

func (w Worker) stop() {
	go func() {
		w.quitChan <- true
	}()
}

func NewDispatcher(jobQueue chan Job, maxWorkers int) *Dispatcher {
	workerPool := make(chan chan Job, maxWorkers)

	return &Dispatcher{
		jobQueue:   jobQueue,
		maxWorkers: maxWorkers,
		workerPool: workerPool,
	}
}

type Dispatcher struct {
	workerPool chan chan Job
	maxWorkers int
	jobQueue   chan Job
}

func (d *Dispatcher) run() {
	for i := 0; i < d.maxWorkers; i++ {
		worker := NewWorker(i+1, d.workerPool)
		worker.start()
	}

	go d.dispatch()
}

func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-d.jobQueue:
			go func() {
				fmt.Printf("fetching workerJobQueue for: %s\n", job.Name)
				workerJobQueue := <-d.workerPool
				fmt.Printf("adding %s to workerJobQueue\n", job.Name)
				workerJobQueue <- job
			}()
		}
	}
}

func requestHandler(w http.ResponseWriter, r *http.Request, jobQueue chan Job) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	delay, err := time.ParseDuration(r.FormValue("delay"))
	if err != nil {
		http.Error(w, "Bad delay value: "+err.Error(), http.StatusBadRequest)
		return
	}

	// 有效处理延迟在 1s~10s
	if delay.Seconds() < 1 || delay.Seconds() > 10 {
		http.Error(w, "The delay must be between 1 and 10 seconds, inclusively.", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	if name == "" {
		http.Error(w, "You must specify a name.", http.StatusBadRequest)
		return
	}
	// 创建新 Job 并将 Job 推送至 jobQueue （由分发者分发）
	job := Job{Name: name, Delay: delay}
	jobQueue <- job

	// Render Success 渲染成功状态
	w.WriteHeader(http.StatusCreated)
}

func main() {
	var (
		maxWorkers   = flag.Int("max_workers", 5, "the number of workers to start")
		maxQueueSize = flag.Int("max_queue_size", 100, "The size of job queue")
		port         = flag.String("port", "8080", "the server port")
	)
	flag.Parse()

	// create the job queue
	jobQueue := make(chan Job, *maxQueueSize)
	// start the dispatcher
	dispatcher := NewDispatcher(jobQueue, *maxWorkers)
	// (go dispatcher())
	dispatcher.run()

	// start the http handler
	http.HandleFunc("/work", func(w http.ResponseWriter, r *http.Request) {
		requestHandler(w, r, jobQueue)
	})
	// start http server which listening on $port
	log.Fatal(http.ListenAndServe(":"+*port, nil))
} // Original code with Dispatcher
