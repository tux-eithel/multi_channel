package main

import (
	"sync"
)

// initSequential takes the input data, reads data sequentially and sends to the producer
// channel. Producers goroutine, read input data, and prepare the structs for consumer goroutine
func initSequential(rows []string, totalRoutine int) {

	// in real case, you don't know the length of rows, but you can "calculate" based on size of file
	var chanProcess chan string = make(chan string, len(rows))
	var chanDati chan *ToElaborate = make(chan *ToElaborate, len(rows))

	waitConsumer := &sync.WaitGroup{}
	waitConsumer.Add(totalRoutine)

	waitProducer := &sync.WaitGroup{}
	waitProducer.Add(totalRoutine)

	for i := 0; i < totalRoutine; i++ {
		go consumerAll(chanDati, waitConsumer)
		go producerSequential(chanProcess, chanDati, waitProducer)
	}

	var srt string
	for _, srt = range rows {
		chanProcess <- srt
	}

	close(chanProcess)
	//	fmt.Println("All data produced")
	waitProducer.Wait()

	close(chanDati)
	waitConsumer.Wait()
	//	fmt.Println("All data consumed")

}

// producerSequential takes data from a channel, prepares it, and sends the new struct
// to chanDati
func producerSequential(chanProcess chan string, chanDati chan *ToElaborate, waitProducer *sync.WaitGroup) {

	defer waitProducer.Done()

	var value string
	var ok bool

	for {
		select {
		case value, ok = <-chanProcess:

			if !ok {
				return
			}

			chanDati <- newToElaborate(value)

		}
	}

}
