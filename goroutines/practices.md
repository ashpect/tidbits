1. Don't create goroutines in library
    1. Let consumer control concurrency
    2. If channel returned through function call, then okay
2. Know when goroutine ends
    1. Avoid memory leaks
3. Check for race conditions at compile time
4. Parallelism
    1. By default,Go uses cpu threads equal to available cores
    2. Change with runtime.GOMAXPROCS
    3. More threads increase performance, too many slows it down. (Have performance tests in production)
