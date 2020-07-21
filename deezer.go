package main

import (
	"fmt"

	deezer "github.com/erebid/go-deezer/deezer"
)

func main() {
	c, err := deezer.NewClient("7579cd89a4d2ab3d6dc2b418446e35c7bd11ba7e62b11d7a2034d888b73864f16a7bc3088a5087a00f53d079eefce6821b0a5e2f746bd9ca8161789a4da11ff7ece21cfbbf692eb7e749c256b1df5bfd4be1e0b1bbc8a441b769d51daea39212")
	if err != nil {
		panic(err)
	}
	res, err := c.Search("shotgun", "", "", 1, 10)
	if err != nil {
		panic(err)
	}
	for _, a := range res.Artists.Data {
		fmt.Println(a)
	}
}
