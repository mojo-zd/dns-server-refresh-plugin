package job

import "github.com/Sirupsen/logrus"

var (
	stop = make(chan string) // the jobs to be stop
)

type Worker struct {
	unRegistry chan string
	jobs       []*HealthJob
}

// Run ...
func (w *Worker) Run(tasks map[string]interface{}) {
	for _, url := range tasks {
		w.AddJob(url.(string))
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

// Watch
func (w *Worker) Watch() {

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

// remove  remove job from w.jobs
func (w *Worker) remove(url string) {
	for i, j := range w.jobs {
		if j.URL == url {
			j = nil
			w.jobs = append(w.jobs[:j], w.jobs[i+1:]...)
		}
	}
}
