package main

import "fmt"
/*
#include <stdlib.h>
#include <stdio.h>
void hello() {
    printf("Hello, Cgo! -- From C world.\n");
}
*/
import "C"

func Random() int {
    return int(C.random())
}

func Seed(i int) {
    C.srandom(C.uint(i))
}

func Hello() {
    C.hello()
}

func main() {
    Seed(100)
    fmt.Println("Random:", Random())
    Hello()
}

