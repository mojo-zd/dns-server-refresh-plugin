package job

import (
	"sync"

	"github.com/Sirupsen/logrus"
)

type WatchCallbackFunc func(args ...interface{})

type Worker struct {
	stop chan string // the jobs to be stop
	jobs []*HealthJob
	sync.Mutex
}

// NewWorker ...
func NewWorker() *Worker {
	worker := &Worker{stop: make(chan string, 10)}
	return worker
}

// Run ...
func (w *Worker) Run(tasks map[string]interface{}) {
	for _, url := range tasks {
		w.AddJob(url.(string))
	}
}

// Stop stop worker
func (w *Worker) Stop() {
	for _, j := range w.jobs {
		w.Remove(j.URL)
		j = nil
	}
}

// AddJob ...
func (w *Worker) AddJob(url string) {
	if w.exist(url) {
		logrus.Infof("job %s has exist ", url)
		return
	}

	job := &HealthJob{URL: url, worker: w}
	w.jobs = append(w.jobs, job)
	job.HealthCheck()
}

// Watch ...
func (w *Worker) Watch(callback WatchCallbackFunc, args ...interface{}) {
	for {
		url := <-w.stop
		w.Remove(url)
		if callback != nil {
			callback(args)
		}
	}
}

func (w *Worker) exist(url string) (exist bool) {
	for _, j := range w.jobs {
		if j.URL == url {
			exist = true
			break
		}
	}
	return
}

// Remove  Remove job from w.jobs
func (w *Worker) Remove(url string) {
	w.Lock()
	defer w.Unlock()
	for i, j := range w.jobs {
		if j.URL == url {
			j = nil
			w.jobs = append(w.jobs[:i], w.jobs[i+1:]...)
		}
	}
}
