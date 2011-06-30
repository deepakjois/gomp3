package main

import "mp3"
import "os"
import "time"
import "runtime"

func main() {
	go func() {
		runtime.LockOSThread()
		mp3.DecodeTrack(os.Args[1])
	}()

	<- time.Tick(5e9)
}
