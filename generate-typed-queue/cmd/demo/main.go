package main

import (
	"fmt"

	queue "github.com/dorant/go-hacks/generate-typed-queue"
)

func main() {
	sq := queue.NewString()
	sq.Enqueue("first")
	sq.Enqueue("second")
	fmt.Println(sq.Dequeue())
}
