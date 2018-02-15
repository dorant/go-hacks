//go:generate echo Hello, Go Generate!
//go:generate stringer -type=pill
package main

import (
	"fmt"

	"github.com/dorant/go-hacks/generate-stringer/project"
)

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
	fmt.Println("Pill:", pill)

	fmt.Println("")
	fmt.Println("Contributors:")
	project.PrintContributors()
}
