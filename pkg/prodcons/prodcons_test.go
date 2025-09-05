package prodcons

import (
	"reflect"
	"sync"
	"testing"
)

func TestProducerConsumer(t *testing.T) {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)

	go producer(ch, &wg)

	consumedCh := make(chan []int)
	go func() {
		consumedCh <- consumer(ch, &wg)
	}()

	wg.Wait()

	consumed := <-consumedCh

	expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	if !reflect.DeepEqual(consumed, expected) {
		t.Errorf("Expected %v, got %v", expected, consumed)
	}

	close(consumedCh)
}
