package main

import (
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Println("# Set key/value to DB...")
	if _, err = conn.Do("SET", "key", "Hello World"); err != nil {
		log.Fatal(err)
	}

	fmt.Println("# Get key/value from DB:")
	str, err := redis.String(conn.Do("GET", "key"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(str)
}
