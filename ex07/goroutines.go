package goroutines

func Process(c chan string) chan string {
	out := make(chan string)
	state := make(chan bool)

	go func() {
		out <- "(" + <-c + ")"
		state <- true
	}()

	go func() {
		<-state
		close(out)
	}()

	return out
}
