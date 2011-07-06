Play an MP3 file in Go using libmpg123 and libao

* For now, make sure you are on OS X. This is because one of the files pulls in CoreFoundation as a CGO dependendency
* Make sure libmpg123 and libao are installed. You could use `brew install mpg123 libao`
* `goinstall github.com/deepakjois/gomp3`
* $GOPATH/bin/play demo.mp3