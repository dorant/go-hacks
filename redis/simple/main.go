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

	if _, err = conn.Do("SET", "key", "Hello World"); err != nil {
		log.Fatal(err)
	}

	str, err := redis.String(conn.Do("GET", "key"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(str)
}
