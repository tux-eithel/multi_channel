// Multi Channel test some possible solution to elaborate a lot data
package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"runtime"
)

// nTime define the number of string to elaborate
const nTime int = 10000

// totalRoutine is circa the number of goroutine it is going to use. This value may be
// boosted using -x flag
var totalRoutine int
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
var multi bool
var typeTest string
var per int

// init prepares the variabiles for parsing
func init() {

	flag.BoolVar(&multi, "mcpu", false, "Scale on all avaible cpu")
	flag.StringVar(&typeTest, "type", "all", "Type of test")
	flag.IntVar(&per, "x", 1, "nCPu * x")

}

// ToElaborate is a struct wich contains an input string, e some funciton which
// simulate some execution over this data
type ToElaborate struct {
	input  string
	toExec func()
}

// Run executes the function who simulated some complex elaboration over the data.
// For every input data the function will be the same
func (te *ToElaborate) Run() {
	te.toExec()
}

// newToElaborate returns a new *ToElaborate struct
func newToElaborate(srt string) *ToElaborate {
	return &ToElaborate{
		srt,
		manyFor,
	}
}

func main() {

	flag.Parse()

	// if per < 1, don't raise an error, but instead set=1 and continue the execution
	if per < 1 {
		per = 1
	}
	totalRoutine = runtime.NumCPU() * per

	if multi {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	switch typeTest {
	case "all":
		// read all data, separate in multiple go routine and use one channel for communicate
		fmt.Println("All")
		initAll(prepareData(), totalRoutine)

	case "seq":
		// read data sequentially, two channel, one for prepare the data, the other
		// for read
		fmt.Println("Seq")
		initSequential(prepareData(), totalRoutine)

	case "multi":
		// read data sequentially, separate the data in multiple channel
		fmt.Println("Multi")
		initMultiChannel(prepareData(), totalRoutine)

	default:
		log.Fatal("Wrong -type options")
	}

}

// manyFor is a fake function to simulate some elaboration
func manyFor() {
	cont := 0
	for i := 1; i < nTime/2; i++ {
		cont++
	}
}

// prepareData creates an array of data
func prepareData() []string {
	var obj = make([]string, nTime)
	for i := 0; i < nTime; i++ {
		obj[i] = randSeq(rand.Int31n(int32(len(letters))))
	}
	return obj
}

// randSeq generates a string of length n
func randSeq(n int32) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
