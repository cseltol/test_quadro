package main

import "sync"

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		Server()
		Client()

		defer wg.Done()
	}()
	wg.Wait()
}