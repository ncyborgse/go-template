package prodcons

import (
	//"fmt"
	"sync"
)

func producer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(ch)

	for i := 0; i < 10; i++ {
		ch <- i
	}

}

func consumer(ch chan int, wg *sync.WaitGroup) []int {
	defer wg.Done()

	var recieved_array = []int{}

	for item := range ch {
		//fmt.Println("Consumed:", item)
		recieved_array = append(recieved_array, item)
	}
	return recieved_array
}

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)

	go producer(ch, &wg)
	go consumer(ch, &wg)

	wg.Wait()
}
