package main

import (
	"fmt"
	"path"
) //var dir, file string
func main() {
	dir, file := path.Split("css/html/design.html")
	fmt.Println("Dir : ", dir)
	fmt.Println("File:", file)
}
