package main

import "C"

//export HelloTest
func HelloTest(s *C.char) *C.char {
	input := C.GoString(s)
	return C.CString("Hello " + input + " ! Nice to meet you !")
}

func main() {
}
