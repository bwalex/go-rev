package main

import (
	"os"
	"fmt"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s <fqdn>\n", os.Args[0]);
		os.Exit(1)
	}

	parts := strings.Split(os.Args[1], ".")
	for i, j := 0, len(parts)-1; i < j; i, j = i+1, j-1 {
		parts[i], parts[j] = parts[j], parts[i]
	}

	str := strings.Join(parts, "/")
	fmt.Printf(str)
}
