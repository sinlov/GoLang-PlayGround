package go_base

import (
	"errors"
	"fmt"
)

//内置有一个error类型，专门用来处理错误信息,并且有专门的包 errors 来处理错误

func go_base_error() {
	err := errors.New("emit macho dwarf: elf header corrupted")
	if err != nil {
		fmt.Print(err)
	}
}
