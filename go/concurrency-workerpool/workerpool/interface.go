package workerpool

import "context"

// WorkerLauncher is an interface for launching workers
type WorkerLauncher interface {
	LaunchWorker(in chan Request, stopCh chan struct{})
}

// Dispatcher is an interface for managing worker pool.
type Dispatcher interface {
	AddWorker(w WorkerLauncher)
	RemoveWorker(minWorkers int)
	LaunchWorker(id int, w WorkerLauncher)
	ScaleWorkers(minWorkers, maxWorkers int, loadThreshold int)
	MakeRequest(Request)
	Stop(ctx context.Context)
}
