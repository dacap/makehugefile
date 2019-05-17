// Copyright (C) 2019 David Capello
//
// This file is released under the terms of the MIT license.
// Read LICENSE.txt for more information.

package main

import (
	"flag"
	"log"
	"math/rand"
	"os"
)

var gigas int
var fast bool

func makeHugeFile(fn string) {
	println("Making file", fn, "with", gigas, "GB")

	file, err := os.OpenFile(fn, os.O_WRONLY | os.O_CREATE | os.O_EXCL, 0444)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	data := make([]byte, 1024*1024)
	rand.Read(data)

	for i := 0; i < gigas*1024; i++ {
		file.Write(data)
		if !fast {
			rand.Read(data)
		}
	}
}

func main() {
	flag.IntVar(&gigas, "g", 1, "GB to write on each file")
	flag.BoolVar(&fast, "f", false, "Fast mode: generates 1MB of random data and repeat")
	flag.Parse()

	for _, file := range flag.Args() {
		makeHugeFile(file)
	}
}
