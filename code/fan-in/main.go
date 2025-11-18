package main

import (
	"context"
	"sync"
	"time"
)

func main() {
	var chans []<-chan int

	for range 5 {
		ch := make(chan int)
		chans = append(chans, ch)

		go func() {
			for i := range 100 {
				ch <- i
			}
			close(ch)
		}()
	}

	ctx, cancel := context.WithTimeout(context.Background(), 400*time.Microsecond)
	defer cancel()

	start := time.Now()

	//res := fanin1(chans...)
	//res := fanin2(chans...)
	res := fanin3(ctx, chans...)

	for val := range res {
		println(val)
	}

	println("duration:", time.Since(start).Microseconds())
}

/*
fanin1 позволяет слить каналы при условии, что нам не важен простой из-за
последовательного перебора каналов.
*/
func fanin1(in ...<-chan int) <-chan int {
	merged := make(chan int)

	go func() {
		for _, ch := range in {
			for val := range ch {
				merged <- val
			}
		}

		close(merged)
	}()

	return merged
}

/*
fanin2 позволяет слить каналы асинхронно, запуская горутины для каждого из каналов.
Горутины будут закрываться только при закрытии каналов-источников.
*/
func fanin2(in ...<-chan int) <-chan int {
	merged := make(chan int)
	wg := sync.WaitGroup{}

	for _, ch := range in {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for val := range ch {
				merged <- val
			}
		}()
	}

	go func() {
		wg.Wait()
		close(merged)
	}()

	return merged
}

/*
fanin3 позволяет слить каналы асинхронно, как и fanin2. Отличием является возможность
завершить горутины по контексту.
*/
func fanin3(ctx context.Context, in ...<-chan int) <-chan int {
	merged := make(chan int)
	wg := sync.WaitGroup{}

	for _, ch := range in {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for {
				select {
				case val, ok := <-ch:
					if !ok {
						return
					}

					merged <- val
				case <-ctx.Done():
					return
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(merged)
	}()

	return merged
}
