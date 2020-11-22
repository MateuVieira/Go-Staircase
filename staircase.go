package main

import (
	"fmt"
	"strings"
)

// Complete the staircase function below.
func staircase(n int32) {
	var i int32
	for i = 1; i <= n; i++ {
		fmt.Printf("%s%s\n",
			strings.Repeat(" ", int(n-i)),
			strings.Repeat("#", int(i)),
		)
	}

}

func main() {
	var n int32 = 9

	staircase(n)
}
