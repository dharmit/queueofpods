package main

import (
	"fmt"
	"k8s.io/client-go/util/workqueue"
	"sync"
)

const concurrentMG = 4

func main() {
	var wg sync.WaitGroup
	queue := workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())

	for i := 0; i < 10; i++ {
		queue.Add(i)
	}
	queue.ShutDownWithDrain()

	wg.Add(concurrentMG)
	for i := 0; i < concurrentMG; i++ {
		go func() {
			defer wg.Done()
			for {
				i, quit := queue.Get()
				if quit {
					return
				}
				defer queue.Done(i)
				fmt.Printf("%v\n", i)
			}
		}()
	}
	wg.Wait()
}
