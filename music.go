package mp3

// #cgo LDFLAGS: -lmpg123
/*
#include <stdio.h>
#include <unistd.h>
#include <stdlib.h>
#include <sys/stat.h>
#include <sys/types.h>
#include <errno.h>
#include "mpg123.h"
#include <ao/ao.h>

char* get_list_at(char **list, int idx) { return list[idx]; }

*/
import "C"

import (
        "fmt"
	"unsafe"
)

func DecodeTrack(file string) {

        var v1 *C.mpg123_id3v1
        var v2 *C.mpg123_id3v2

        C.mpg123_init()
        defer C.mpg123_exit()

        m := C.mpg123_new(nil, nil)
        defer C.mpg123_delete(m)

        f := C.CString(file)

        if err := C.mpg123_open(m, f); err != C.MPG123_OK {
                panic("Error reading file")
        }
        defer C.mpg123_close(m)

        C.mpg123_scan(m)
        meta := C.mpg123_meta_check(m)

        if meta == C.MPG123_ID3 && C.mpg123_id3(m, &v1, &v2) == C.MPG123_OK {
                var title, artist, album, genre string
                switch false {
                case v2 == nil:
                        fmt.Println("ID3V2 tag found")
                        title = C.GoString(v2.title.p)
                        artist = C.GoString(v2.artist.p)
                        album = C.GoString(v2.album.p)
                        genre = C.GoString(v2.genre.p)

                case v1 == nil:
                        fmt.Println("ID3V2 tag found")
                        title = C.GoString(&v1.title[0])
                        artist = C.GoString(&v1.artist[0])
                        album = C.GoString(&v1.album[0])
                        genre = "Unknown" // FIXME convert int to string
                }

                fmt.Println(title)
                fmt.Println(artist)
                fmt.Println(album)
                fmt.Println(genre)
        }


	C.ao_initialize()
	defer C.ao_shutdown()

	var ret C.int
	var fill C.size_t
	buf := make([]C.uchar,1024*16)

        for {
                ret = C.mpg123_read(m, (*C.uchar)(unsafe.Pointer(&buf)), 16*1024, &fill)
		if ret == C.MPG123_DONE {
			break;
		}
        }
}

func ListDecoders() {
        C.mpg123_init()
        defer C.mpg123_exit()

        m := C.mpg123_new(nil, nil)
        defer C.mpg123_delete(m)

        decoders := C.mpg123_decoders()
        val_ptr := C.get_list_at(decoders, C.int(0))

        if val_ptr == nil {
                fmt.Println("nil")
        }
        str := C.GoString(val_ptr)
        fmt.Println(str)
}
