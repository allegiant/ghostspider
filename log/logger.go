package main

import (
	"fmt"
	"runtime"
	"log"
)

type Slice struct {
	X string
	Y string
}

func (slice Slice) Ab(str string) {
	fmt.Println(slice.X + slice.Y)
	fmt.Println(str)
}

func info() {

}



func test() {
    pc,file,line,ok := runtime.Caller(0)
    log.Println(pc)
    log.Println(file)
    log.Println(line)
    log.Println(ok)
    f := runtime.FuncForPC(pc)
    log.Println(f.Name())

    pc,file,line,ok = runtime.Caller(1)
    log.Println(pc)
    log.Println(file)
    log.Println(line)
    log.Println(ok)
    f = runtime.FuncForPC(pc)
    log.Println(f.Name())
}

func main() {
	test()
}

