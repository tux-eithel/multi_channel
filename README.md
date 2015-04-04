# Big source of data, different channel implementations
===
Suppose you have a lot of data to elaborate, like for example a big file,
which is the best way to subdivide the work in multiple goroutine?

Three approaches:
* All data in memory, (like after `ReadAll` function, not really suitable for big file), **one channel**, *n* goroutines take a piece of initial data array and prepare for *n* consumer goroutines.
* Read data sequentially, *n* goroutines prepare data reading data from one channel and send to another, n goroutines will consume data, overall **two channel** will be used.
* Read data sequentially, two goroutine prepare *n* channel for producer and consumer. Every channel has its own goroutine which produce/consume. Data will be spread along the channels. **Multiple channels** are involved.

### Results
```
testing: warning: no tests to run
PASS
BenchmarkAll	             30	  38862375 ns/op
BenchmarkSequential          30   40441853 ns/op
BenchmarkMulti               30   40387717 ns/op
BenchmarkAllCPU	            100	  17431929 ns/op
BenchmarkSequentialCPU	    100	  18935215 ns/op
BenchmarkMultiCPU            30   51210445 ns/op
```