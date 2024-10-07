package main

import "fmt"
import (
	mod1 "github.com/techfunda4all/workspace/module-1"
	mod2 "github.com/techfunda4all/workspace/module-2"
)

func main() {
	fmt.Println(mod1.Name())
	fmt.Println(mod2.Name())
}
