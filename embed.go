package main

import _ "embed"

//go:embed data/post.txt
var fileString string

func main() {
	println(fileString)
}
