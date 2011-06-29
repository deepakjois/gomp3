package main

import "mp3"
import "os"

func main() {
	mp3.PrintInfo(os.Args[1])
	mp3.ListDecoders()
	mp3.TestLibAo()
}
