package main

import (
	"log"
	"sync"
	"time"
)

type Processor struct {
	queue    chan int
	workers  sync.WaitGroup
	shutdown chan struct{}
}

func (p *Processor) Run(count int) {
	p.shutdown = make(chan struct{})
	for i := 0; i < count; i++ {
		p.workers.Add(1)
		go p.worker(i)
	}
}

func (p *Processor) Stop() {
	close(p.shutdown)
	p.workers.Wait()
}

func (p *Processor) Submit(i int) {
	p.queue <- i
}

func (p *Processor) worker(id int) {
	defer p.workers.Done()
	for {
		select {
		case item := <-p.queue:
			log.Printf("%3d: (%d) WORKING", item, id)
			if item%2 == 0 {
				time.Sleep(time.Second * 3)
			} else {
				time.Sleep(time.Second * 5)
			}
			log.Printf("%3d: (%d) WORKED", item, id)

		case <-p.shutdown:
			log.Printf("Shutdown %d", id)
			return
		}
	}
}

func main() {
	processor := &Processor{
		queue:   make(chan int, 1),
		workers: sync.WaitGroup{},
	}
	processor.Run(4) // start with 4 workers

	go func() {
		var i = 0
		for {
			if i < 100 {
				log.Printf("%3d: ADDING", i)
				processor.Submit(i)
				log.Printf("%3d: ADDED", i)
			}
			if i == 10 {
				processor.Stop() // waits until existing workers have completed
				processor.Run(2) // restart with only 2 workers
			}
			if i == 15 {
				processor.Stop()
				processor.Run(20) // restart with 20 workers
			}
			i++
		}
	}()

	forever := make(chan struct{})
	<-forever
}
