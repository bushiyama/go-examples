package main

/*
#cgo LDFLAGS: -L./lib -lcrustaceanize
#include <stdlib.h>
#include "./lib/crustaceanize.h"
*/
import "C"

import (
	"fmt"
	"unsafe"
)

func main() {
	s := "I'm a Gopher"

	input := C.CString(s)               // Goの管理化のポインタではなくなる。
	defer C.free(unsafe.Pointer(input)) // そのためメモリ解放を実装する必要がある。

	/*
	 * @memo:
	 * 	以下の場合はinputのメモリはGoの管理化である。
	 * 	このときGo側でGCが働くのでこのプログラムではランタイムエラーが発生する!
	 * 		data := (*reflect.StringHeader)(unsafe.Pointer(&s)).Data
	 * 		input := (*C.char)(unsafe.Pointer(data))
	 */

	o := C.crustaceanize(input)

	output := C.GoString(o)
	fmt.Printf("%s\n", output)
}
