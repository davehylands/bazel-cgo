//
// main.go
// Copyright (C) 2019 Tim Hughes
//
// Distributed under terms of the MIT license.
//
package main

/*
#cgo LDFLAGS: -L../person-rs/target/debug -lperson
#include "person.h"

// Forward declaration of the gateway function.
void MyCallback_cgo(int num);
*/
import "C"
import (
	"fmt"
	"unsafe"
)

type (
	Person C.struct_APerson
)

func (p *Person) Name() string {
	return C.GoString(p.name)
}

func (p *Person) LongName() string {
	return C.GoString(p.long_name)
}

func GetPerson(name string, long_name string) *Person {
	return (*Person)(C.get_person(C.CString(name), C.CString(long_name)))
}

func FreePerson(person *Person) {
	C.free_person((*C.struct_APerson)(person))
}

//export MyCallback
func MyCallback(num C.int) {

	fmt.Printf("my_callback: num = %d\n", num)
}

func DoSomeWork() {
	C.do_some_work((C.CallBackFuncPtr)(unsafe.Pointer(C.MyCallback_cgo)))
}

func main() {
	var f *Person
	f = GetPerson("tim", "tim hughes")
	fmt.Printf("Hello Go rust world: My name is %s, %s.\n", C.GoString(f.name), C.GoString(f.long_name))
	fmt.Printf("Hello Go ruat world: My name is %s, %s.\n", f.Name(), f.LongName())
	FreePerson(f)

	DoSomeWork()
}
