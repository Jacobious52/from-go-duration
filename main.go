package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func main() {
	s, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	d, err := time.ParseDuration(strings.Trim(string(s), "\n "))
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	fmt.Print(d.Seconds(), "sec")
}
