package main
 
// #include <stdio.h>
// int puts(const char*);
import "C"
import "unsafe"
 
func main() {
        buf := []byte("Hello world\x00")
        ptr := unsafe.Pointer(&buf[0])
        C.puts((*C.char)(ptr))
}
