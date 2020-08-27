package parallelexecuter

import "time"

func ParallelExec(a []func(int) error, n int, m int, t int) []error { // function that was made to execute functions parallely using goroutines the function receives a slice
	chErr := make(chan error, m) // creating buffered channel chErr	  // of functions and three variablrs as input, as output function returns slice of errors
	for i := 0; i < n; i++ {     // using cycle to create n parallel goroutines
		x := a[i]
		go func() { // goroutine calls closure that in turn performs the rest of actions
			if err := x(t); err != nil {
				if len(chErr) != m {
					chErr <- err // sending error into the channel
				} else {
					return // if channel is full program stops processing
				}
			}
		}()
	}
	time.Sleep(5 * time.Second) // timer is used to cynchronize goroutines
	close(chErr)                //closing channel chErr
	var slerr []error
	for er := range chErr {
		slerr = append(slerr, er)
	}
	return slerr
}
