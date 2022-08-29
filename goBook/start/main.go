package main

import (
"fmt"
"runtime"
)

func main() {
fmt.Println(runtime.GOMAXPROCS(1))
fmt.Println(runtime.GOMAXPROCS(0))

go func() {
fmt.Println("Hello world")
}()

for {

}

}