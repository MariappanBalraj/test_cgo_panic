//hello.go
package main

/*
#include <stdio.h>

extern void Test4(void);

void test1(void) {
   Test4();
}

void test2(void) {
    int val = 2;
    test1();
}

void test3(void) {
    int val = 3;
    test2();
}

void test4(void) {
    printf("Test4()");
}
*/
import "C"

func Test5() {
    C.test4()
}

func main() {
    C.test3()
}
