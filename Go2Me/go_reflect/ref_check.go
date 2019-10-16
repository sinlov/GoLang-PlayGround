package go_reflect

import "fmt"

const anonymousTypeCheckCount int = 5

func anonymousTypeCheck() {
	for i := 0; i < anonymousTypeCheckCount; i++ {
		p := func(i int) {
			fmt.Printf("now - count: %d", i)
		}
		p(i)
		fmt.Printf("- p typeOf %T and value %p\n", p, p)
	}
}
