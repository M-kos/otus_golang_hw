package hw05parallelexecution

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	tCh := make(chan Task, len(tasks))
	done := make(chan struct{})
	errCh := make(chan struct{})
	var err error
	var wg sync.WaitGroup
	// var tasksCounter int64
	// errCounter := int64(m)

	wg.Add(n)

	go func() {
		for i := 0; i < n; i++ {
			go func() {
				defer wg.Done()

				select {
				case t := <-tCh:
					if err := t(); err != nil {
						fmt.Println(1, err)
						errCh <- struct{}{}
						// atomic.AddInt64(&errCounter, -1)
					}

					// atomic.AddInt64(&tasksCounter, 1)
				case <-done:
					return
				}

				// for task := range tCh {
				// 	if err := task(); err != nil {
				// 		atomic.AddInt64(&errCounter, -1)
				// 	}

				// 	atomic.AddInt64(&tasksCounter, 1)
				// }
			}()

		}

		wg.Wait()
		// close(errCh)
	}()

	go func() {
		for i := 0; i < len(tasks); i++ {
			fmt.Println(2, i)
			tCh <- tasks[i]
		}
		close(tCh)
	}()

	for {
		fmt.Println("atomic.LoadInt64(&errCounter)", atomic.LoadInt64(&errCounter))
		fmt.Println("atomic.LoadInt64(&tasksCounter)", atomic.LoadInt64(&tasksCounter))
		if atomic.LoadInt64(&errCounter) <= 0 && m > 0 || atomic.LoadInt64(&errCounter) < 0 {
			done <- struct{}{}
			err = ErrErrorsLimitExceeded
			break
		}

		if atomic.LoadInt64(&tasksCounter) >= int64(len(tasks)) {
			done <- struct{}{}
			break
		}
	}

	return err
}

// // Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
// func Run(tasks []Task, n, m int) error {
// 	tCh := make(chan Task)
// 	var err error
// 	var wg sync.WaitGroup
// 	var tasksCounter int64
// 	errCounter := int64(m)

// 	wg.Add(n)

// 	for i := 0; i < n; i++ {
// 		go func() {
// 			defer wg.Done()

// 			for task := range tCh {
// 				if err := task(); err != nil {
// 					atomic.AddInt64(&errCounter, -1)
// 				}

// 				atomic.AddInt64(&tasksCounter, 1)
// 			}
// 		}()
// 	}

// 	for i := 0; i < len(tasks); i++ {
// 		tCh <- tasks[i]

// 		if atomic.LoadInt64(&errCounter) <= 0 && m > 0 || atomic.LoadInt64(&errCounter) < 0 {
// 			close(tCh)
// 			err = ErrErrorsLimitExceeded
// 			break
// 		}
// 	}

// 	if err == nil {
// 		close(tCh)
// 	}

// 	wg.Wait()

// 	return err
// }
