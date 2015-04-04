package main

import (
	"sync"
)

// initAll takes the input data, runs totalRoutine (+1) which receive a piece of
// the initial array.
// chanDati is a channel used for producer and consumer routines
func initAll(rows []string, totalRoutine int) {

	// all data is in memory, so it is possible create a channel with the right size
	var chanDati chan *ToElaborate = make(chan *ToElaborate, len(rows))

	waitConsumer := &sync.WaitGroup{}
	waitConsumer.Add(totalRoutine)

	for i := 0; i < totalRoutine; i++ {
		go consumerAll(chanDati, waitConsumer)
	}

	divide := len(rows) / totalRoutine
	lim1 := 0
	lim2 := divide

	waitProducer := &sync.WaitGroup{}
	if len(rows)%totalRoutine == 0 {
		waitProducer.Add(totalRoutine)
	} else {
		waitProducer.Add(totalRoutine + 1)
	}

	for lim1 < len(rows) {
		if lim2 <= len(rows) {
			go producerAll(rows[lim1:lim2], chanDati, waitProducer)
		} else {
			go producerAll(rows[lim1:], chanDati, waitProducer)
		}
		lim1 += divide
		lim2 += divide
	}

	waitProducer.Wait()
	//	fmt.Println("All data produced")
	close(chanDati)

	waitConsumer.Wait()
	//	fmt.Println("All data consumed")

}

// producerAll sends to channel *ToElaborate struct
func producerAll(row []string, chanDati chan *ToElaborate, waitProducer *sync.WaitGroup) {

	defer waitProducer.Done()

	var value string

	for _, value = range row {
		chanDati <- newToElaborate(value)
	}

}

// consumerAll reads data from channel, and then elaborates one data at the time
func consumerAll(chanDati chan *ToElaborate, waitConsumer *sync.WaitGroup) {

	defer waitConsumer.Done()

	var value *ToElaborate
	var ok bool

	for {

		select {
		case value, ok = <-chanDati:

			if !ok {
				return
			}

			value.Run()

		}

	}

}
