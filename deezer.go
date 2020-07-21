package main

import (
	"fmt"

	deezer "github.com/erebid/go-deezer/deezer"
)

func main() {
	c, err := deezer.NewClient("no")
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
