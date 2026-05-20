package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func main() {
	var input []byte
	var err error

	if len(os.Args) > 1 {
		input, err = os.ReadFile(os.Args[1])
	} else {
		input, err = io.ReadAll(os.Stdin)
	}
	if err != nil { fmt.Fprintln(os.Stderr, "Error:", err); os.Exit(1) }

	var buf bytes.Buffer
	if err = json.Indent(&buf, bytes.TrimSpace(input), "", "  "); err != nil {
		fmt.Fprintln(os.Stderr, "❌ Invalid JSON:", err)
		os.Exit(1)
	}
	fmt.Println(buf.String())
}
