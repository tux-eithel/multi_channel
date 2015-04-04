package main

import (
	"sync"
)

// initMultiChannel takes the input data, prepare totalRoutine channel (for producer
// and consumer), where data will be spread. Every channel has only one routine
// to produce/consume data
func initMultiChannel(rows []string, totalRoutine int) {

	// in real case, you don't know the length of rows, but you can "calculate" based on size of file
	var chanProcess []chan string = make([]chan string, totalRoutine)
	var chanDati []chan *ToElaborate = make([]chan *ToElaborate, totalRoutine)

	smallSize := len(rows) / totalRoutine
	for i := 0; i < totalRoutine; i++ {
		chanProcess[i] = make(chan string, smallSize)
		chanDati[i] = make(chan *ToElaborate, smallSize)
	}

	waitConsumerMulti := &sync.WaitGroup{}
	waitConsumerMulti.Add(1)

	waitProducerMulti := &sync.WaitGroup{}
	waitProducerMulti.Add(1)

	go consumerMultiChannelMulti(chanDati, waitConsumerMulti)
	go producerMultiChannelMulti(chanProcess, chanDati, waitProducerMulti)

	var srt string
	var i int = 0
	for _, srt = range rows {

		if len(chanProcess) == i {
			i = 0
		}

		chanProcess[i] <- srt
	}

	for _, value := range chanProcess {
		close(value)
	}
	waitProducerMulti.Wait()

	for _, value := range chanDati {
		close(value)
	}
	waitConsumerMulti.Wait()

}

// producerMultiChannelMulti prepares the routine which preapre the data
func producerMultiChannelMulti(chanProcess []chan string, chanDati []chan *ToElaborate, waitProducerMulti *sync.WaitGroup) {

	defer waitProducerMulti.Done()

	waitProducer := &sync.WaitGroup{}
	waitProducer.Add(len(chanProcess))

	for _, value := range chanProcess {
		func(chanProcess chan string, waitProducer *sync.WaitGroup) {
			go producerMultiChannel(value, chanDati, waitProducer)
		}(value, waitProducer)
	}

	waitProducer.Wait()

}

// producerMultiChannel takes input data from a channel and spread all over the channels
func producerMultiChannel(chanProcess chan string, chanDati []chan *ToElaborate, waitProducer *sync.WaitGroup) {

	defer waitProducer.Done()

	var i int = 0
	var value string
	var ok bool

	for {
		select {
		case value, ok = <-chanProcess:

			if !ok {
				return
			}

			if len(chanDati) == i {
				i = 0
			}

			chanDati[i] <- newToElaborate(value)
		}
	}

}

// consumerMultiChannelMulti prepares the routines for consume data, one for channel
func consumerMultiChannelMulti(chanDati []chan *ToElaborate, waitConsumerMulti *sync.WaitGroup) {

	defer waitConsumerMulti.Done()

	waitConsumer := &sync.WaitGroup{}
	waitConsumer.Add(len(chanDati))

	for _, value := range chanDati {
		func(chanDati chan *ToElaborate, waitConsumer *sync.WaitGroup) {
			go consumerAll(chanDati, waitConsumer)
		}(value, waitConsumer)
	}

	waitConsumer.Wait()

}
