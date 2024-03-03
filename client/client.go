// clearCommitHistory
package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	args := os.Args
	if len(args) < 3 {
		panic("порт кол-во запросов")
	}

	port := args[1]
	requests, err := strconv.Atoi(args[2])
	if err != nil {
		panic("ошибка конвертации кол-во запросов")
	}

	var wg sync.WaitGroup
	start := time.Now()
	for range requests {
		wg.Add(1)
		go func() {
			SendRequest(port)
			wg.Done()
		}()
	}
	wg.Wait()
	duration := time.Since(start)
	fmt.Println(duration)
}
