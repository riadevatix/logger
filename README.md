# Logger

Logger is a simple loging package implementation in go and a proof of concept. Sometimes we don't want to show the logs without but we don't want to remove the lines from source either. In that case we can use some inteface representing a an actual `Logger`. If we don't need the actual logger any more we can use a dummy logger.

## Example

```go
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
```
