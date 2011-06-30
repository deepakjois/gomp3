package main

import "mp3"
import "os"

func main() {
	mp3.Init()
	defer mp3.Shutdown()

	var control = make(chan int)

	go func() {
		mp3.DecodeTrack(os.Args[1],control)
	}()

	<-control
}
