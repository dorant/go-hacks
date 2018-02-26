package main

import (
	"encoding/asn1"
	"fmt"
	"time"
)

func main() {
	// Int
	fmt.Println("------------- int")
	val1 := 1345678
	fmt.Println("Before:", val1)

	data, err := asn1.Marshal(val1)
	if err != nil {
		panic(err)
	}
	fmt.Println("Marshal data:", data)

	var n int64
	_, err = asn1.Unmarshal(data, &n)
	if err != nil {
		panic(err)
	}
	fmt.Println("Unmarshal:", n)

	// String
	fmt.Println("------------- string")
	val2 := "1345678"
	fmt.Println("Before:", val2)

	data, err = asn1.Marshal(val2)
	if err != nil {
		panic(err)
	}
	fmt.Println("Marshal data:", data)

	var s string
	_, err = asn1.Unmarshal(data, &s)
	if err != nil {
		panic(err)
	}
	fmt.Println("Unmarshal:", s)

	// UTC time
	fmt.Println("------------- time")
	val3 := time.Now()
	fmt.Println("Before:", val3)

	data, err = asn1.Marshal(val3)
	if err != nil {
		panic(err)
	}
	fmt.Println("Marshal data:", data)

	var t = new(time.Time)
	_, err = asn1.Unmarshal(data, t)
	if err != nil {
		panic(err)
	}
	fmt.Println("Unmarshal:", t)

	// Struct
	fmt.Println("------------- struct")

	type MyType1 struct {
		I int
		S string
	}

	type MyType2 struct {
		I  int
		S  string
		T1 int `asn1:"optional,default:55"`
	}

	val4 := MyType1{5, "test"}
	fmt.Println("Before MyType1:", val4)

	data, err = asn1.Marshal(val4)
	if err != nil {
		panic(err)
	}
	fmt.Println("Marshal data:", data)

	var typ1 MyType1
	_, err = asn1.Unmarshal(data, &typ1)
	if err != nil {
		panic(err)
	}
	fmt.Println("Unmarshal MyType1:", typ1)

	var typ2 MyType2
	_, err = asn1.Unmarshal(data, &typ2)
	if err != nil {
		panic(err)
	}
	fmt.Println("Unmarshal MyType2:", typ2)

}
