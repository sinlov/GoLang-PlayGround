package goBuild

import "fmt"

func init() {
	fmt.Println("just goBuild package init!")
}

func Copy(from string, to string) {
	fmt.Printf("build copy from %v, to %v\n",from, to)
}
