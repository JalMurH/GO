package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

// definicion de clases(estructuras)
type Job struct {
	Name   string
	Delay  time.Duration
	Number int
}

type Worker struct {
	Id         int
	JobQueue   chan Job
	WorkerPool chan chan Job
	QuitChan   chan bool
}

type Dispatcher struct {
	WorkerPool chan chan Job
	MaxWorkers int
	JobQueue   chan Job
}

// constructor para los workers
func NetWorker(id int, WorkerPool chan chan Job) *Worker {
	return &Worker{
		Id:         id,
		JobQueue:   make(chan Job),
		WorkerPool: WorkerPool,
		QuitChan:   make(chan bool),
	}
}

// metodo para iniciar los workers
func (w Worker) Start() {
	//gorutine funciona anonima que se ejecutara para uso de la concurrencia
	go func() {
		for {
			//recibe en el atributo workerpool los trabajos
			w.WorkerPool <- w.JobQueue

			select {
			//en caso que se reciba el trabajo a uno de los trabajadores se inicia y se calcula el fibonacci
			case job := <-w.JobQueue:
				fmt.Printf("Worker ID: %d started\n", w.Id)
				resFib := Fibonacci(job.Number)
				//se pausa el proseso hasta que acaba
				time.Sleep(job.Delay)
				fmt.Printf("Worker ID: %d Finished, Result: %d\n", w.Id, resFib)
			case <-w.QuitChan:
				//detiene el trabajardor
				fmt.Printf("Worker ID: %d, Stopped\n", w.Id)
			}
		}
	}()
}

// metodo para detener los workers
func (w Worker) Stop() {
	go func() {
		w.QuitChan <- true
	}()
}

// constructor del despachador que coordina los trabajos y los workers
func NewDispatcher(jobQueue chan Job, maxWorkers int) *Dispatcher {
	worker := make(chan chan Job, maxWorkers)
	return &Dispatcher{
		JobQueue:   jobQueue,
		MaxWorkers: maxWorkers,
		WorkerPool: worker,
	}
}

// metodo del despachador
func (d *Dispatcher) Dispatch() {
	for {
		select {
		case job := <-d.JobQueue:
			go func() {
				workerJobQueue := <-d.WorkerPool
				workerJobQueue <- job
			}()

		}
	}
}

// funcion que calcula el fibonacci de forma recursiba
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func (d *Dispatcher) Run() {
	for i := 0; i < d.MaxWorkers; i++ {
		worker := NetWorker(i, d.WorkerPool)
		worker.Start()
	}

	go d.Dispatch()
}

func requestHandler(w http.ResponseWriter, r *http.Request, jobQueue chan Job) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	delay, err := time.ParseDuration(r.FormValue("delay"))
	if err != nil {
		http.Error(w, "Invalid delay", http.StatusBadRequest)
		return
	}

	value, err := strconv.Atoi(r.FormValue("value"))
	if err != nil {
		http.Error(w, "Invalid value", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	if name == "" {
		http.Error(w, "Invalid name", http.StatusBadRequest)
	}

	job := Job{Name: name, Delay: delay, Number: value}
	jobQueue <- job
	w.WriteHeader(http.StatusCreated)
}

func main() {
	const (
		maxWorkers   = 4
		maxQueueSize = 20
		port         = ":8081"
	)
	jobQueue := make(chan Job, maxQueueSize)
	dispatcher := NewDispatcher(jobQueue, maxWorkers)

	dispatcher.Run()
	http.HandleFunc("/fib", func(w http.ResponseWriter, r *http.Request) {
		requestHandler(w, r, jobQueue)
	})
	log.Fatal(http.ListenAndServe(port, nil))
}
