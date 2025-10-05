package main

func main() {
	ready := make(chan bool)
	close(ready)
	// panic
	close(ready)
}
