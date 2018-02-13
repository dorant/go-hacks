/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package main

import (
	"fmt"
	"log"
	"os"
	"time"

	pb "github.com/dorant/go-grpc-tryout/helloworld"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func waitForGrpcReady(cc grpc.ClientConn, timeout int) error {
	fmt.Printf("Wait %d sec for ready state...: (%v)", timeout, cc.GetState())
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	for {
		s := cc.GetState()
		if s == connectivity.Ready {
			break
		}
		if !cc.WaitForStateChange(ctx, s) {
			// ctx got timeout or canceled.
			return ctx.Err()
		}
	}
	return nil
}

func main() {
	// Set up a connection to the server.:
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Printf("did not connect: %v\n", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	err = waitForGrpcReady(*conn, 10)
	if err != nil {
		log.Printf("did not connect: %v\n", err)
	}
	fmt.Printf("Done (%v)\n", conn.GetState())

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	for i := 0; i < 40; i++ {
		log.Printf("State: %v\n", conn.GetState())
		fmt.Printf("Wait no. %d...: ", i)
		time.Sleep(2 * time.Second)

		r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
		if err != nil {
			log.Printf("could not greet: %v\n", err)
		} else {
			fmt.Printf("Done: %s\n", r)
		}
	}
}
