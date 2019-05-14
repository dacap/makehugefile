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

func makeHugeFile(fn string, data []byte, gigas int) {
	println("Making file", fn, "with", gigas, "GB")

	file, err := os.OpenFile(fn, os.O_WRONLY | os.O_CREATE | os.O_EXCL, 0444)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	for i := 0; i < gigas*1024; i++ {
		file.Write(data)
	}
}

func main() {
	data := make([]byte, 1024*1024)
	rand.Read(data)

	var gigas int
	flag.IntVar(&gigas, "g", 1, "GB to write on each file")
	flag.Parse()

	for _, file := range flag.Args() {
		makeHugeFile(file, data, gigas)
	}
}
