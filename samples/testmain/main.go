package main

import (
	"fmt" 
	"path/filepath"
 )

func main() {
 
	path := "../xxx/a/b/c/d/e/f/g"
	i:= 0
	for len(path)>0 && i<15 {
		fmt.Println(path)
		i++
		path = filepath.Dir(path)

	}

	fmt.Println("last:",path) 
}
