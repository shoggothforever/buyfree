package mrpc

func practice() {
	PlatFormService = NewWorkerPool(WORKERNUMS)
	for i := 0; i < WORKERNUMS; i++ {
		worker := NewWorker()
		worker.Run()
	}
	go func() {
		for {
			select {
			case req := <-PlatFormService.ReqChan:
				{
					worker := <-PlatFormService.WorkerChan
					worker.ReqChan <- req
					if val := <-worker.ReplyChan; val == true {
						PlatFormService.WorkerChan <- worker
					} else {
						close(worker.ReqChan)
						close(worker.ReplyChan)
						worker = nil
						newworker := NewWorker()
						PlatFormService.WorkerChan <- newworker
						newworker.Run()

					}
				}

			}
		}
	}()
}

type CountRequest struct {
	Iterator int64 `json:"iterator"`
	Communicator
}

func NewCountRequest(it int64) *CountRequest {
	req := &CountRequest{Iterator: it, Communicator: NewCommunicator()}
	return req
}

func (o *CountRequest) Handle() {
	o.ReplyChan <- true
	//fmt.Println("管道大小", len(o.ReplyChan))
}
