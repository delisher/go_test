package main

var a int

func main() {
    a = 5
    print(a)
    f()
}

func f() {
    a := 6
    print(a)
    g()
}

func g() {
    print(a)
}