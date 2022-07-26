package main

/*
#include <stdio.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"

func malloc() unsafe.Pointer {
	return unsafe.Pointer(C.CString("test"))
}

func free(ptr unsafe.Pointer) {
	C.free(ptr)
}

func main() {}
