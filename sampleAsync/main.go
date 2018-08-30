package sampleAsync

import (
	"fmt"
	"strconv"
	"sync"
)

type Thing struct {
	Name string
}

var Title map[string]int

func main() {
	doConcurrently()
}

func doConcurrently() {

	// title := make(map[string]int)

	var (
		things   = make(chan Thing)
		finished = make(chan struct{})
		wg       sync.WaitGroup
	)

	for i := 0; i < 4; i++ {
		wg.Add(1)
		iAsStr := strconv.Itoa(i)

		go func(wg *sync.WaitGroup, j string) {
			defer wg.Done()
			Title[j] = 2
			fmt.Println(Title[j])
		}(&wg, iAsStr)
	}
	wg.Wait()

	go func() {
		// will consume until close
		consume(things)
		// signal consumption has finished
		close(finished)
	}()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()

			things <- Thing{
				Name: fmt.Sprintf("Paijo %v", j),
			}
		}(i)
	}

	// wait until all producers have stopped
	wg.Wait()

	// then you can close
	close(things)

	// wait until finished consuming
	<-finished
}

func consume(things <-chan Thing) {
	// will do work until close
	for thing := range things {
		// do work
		fmt.Println(thing.Name)
	}
}
