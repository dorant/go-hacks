//go:generate stringer -type=pill
package main

import "fmt"

type pill int

const (
	placebo pill = iota
	aspirin
	ibuprofen
	paracetamol
	acetaminophen = paracetamol
)

func main() {
	pill := aspirin
	fmt.Println(pill)
}
