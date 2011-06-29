package main

import "mp3"
import "os"

func main() {
	mp3.DecodeTrack(os.Args[1])
	mp3.TestLibAo()
}
