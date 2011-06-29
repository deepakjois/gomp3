package mp3

// #cgo LDFLAGS: -lao -lm
/*
#include <stdio.h>
#include <ao/ao.h>
#include <math.h>

*/
import "C"

func TestLibAo() {
	C.ao_initialize()

	/* -- Setup for default driver -- */

	//default_driver := C.ao_default_driver_id()

	C.ao_shutdown()

}
