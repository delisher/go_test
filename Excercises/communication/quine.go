// We just need code used to do the actual printing
// and data that represents the textual form of the code
// see also https://en.wikipedia.org/wiki/Quine_(computing)
package main

import "fmt"

func main() {
	fmt.Printf("%s%c%s%c\n", s, 0x60, s, 0x60) // 0x60 is '
}

var s = `// We just need code used to do the actual printing
// and data that represents the textual form of the code
// see also https://en.wikipedia.org/wiki/Quine_(computing)
package main

import "fmt"

func main() {
	fmt.Printf("%s%c%s%c\n", s, 0x60, s, 0x60) // 0x60 is '
}

var s = `
