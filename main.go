package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func printFib(i int, n int, wg *sync.WaitGroup, ch chan string) {
	defer wg.Done()
	result := fib(n)
	ch <- fmt.Sprintf("Result %d, Fibonacci of %d is %d", i, n, result)
}

func main() {
	RESULT_COUNT := 20
	FIBONACCI_NUMBER := 42

	NUM_THREADS := runtime.NumCPU()
	// NUM_THREADS := 1
	runtime.GOMAXPROCS(NUM_THREADS)

	fmt.Printf("USING %v THREADS\n", NUM_THREADS)

	fmt.Println("STARTING...")

	var wg sync.WaitGroup
	ch := make(chan string, RESULT_COUNT)

	for i := 1; i <= RESULT_COUNT; i++ {
		wg.Add(1)
		go printFib(i, FIBONACCI_NUMBER, &wg, ch)
		time.Sleep(10 * time.Millisecond)
	}

	fmt.Printf("CREATED %v WORKERS\n", RESULT_COUNT)

	startTime := time.Now()

	go func() {
		wg.Wait()
		close(ch)
	}()

	fmt.Println("WAITING FOR RESULTS...")

	for result := range ch {
		fmt.Println(result)
	}

	elapsedTime := time.Since(startTime)
	fmt.Printf("ELAPSED TIME: %v\n", elapsedTime)

	fmt.Println("EXITING...")
}
