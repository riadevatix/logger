package main

import (
	"fmt"

	"github.com/riadevatix/logger/logger"
)

func main() {

	// If we use logger.NoOpLogger() no log will be printed

	l := logger.DefaultLogger()
	// l := logger.NoOpLogger()

	l.Logln("Declaring variable a")
	a := 2
	l.Logln("Declared variable a = ", a)

	l.Logln("Declaring variable b")
	b := 3
	l.Logln("Declared variable b = ", b)

	l.Logln("Summing up a and b")
	sum := a + b

	fmt.Printf("Sum of %d + %d = %d\n", a, b, sum)
}
