package go_base

import "os"

// 类似java的异常机制,但不建议使用这个
// 把它作为最后的手段来使用，也就是说，你的代码中应当没有，或者很少有panic的东西

//  -- Panic --
// 内建函数，可以中断原有的控制流程，进入一个令人恐慌的流程中。
// 当函数F调用panic，函数F的执行被中断，但是F中的延迟函数会正常执行，然后F返回到调用它的地方。
// 在调用的地方，F的行为就像调用了panic。
// 这一过程继续向上，直到发生panic的goroutine中所有调用的函数返回，此时程序退出。
// 恐慌可以直接调用panic产生。也可以由运行时错误产生，例如访问越界的数组。


// -- Recover --
// 内建的函数，可以让进入令人恐慌的流程中的goroutine恢复过来。
// recover仅在延迟函数中有效。
// 在正常的执行过程中，调用recover会返回nil，并且没有其它任何效果。
// 如果当前的goroutine陷入恐慌，调用recover可以捕获到panic的输入值，并且恢复正常的执行。

var user = os.Getenv("User")

func init() {
	if user == "" {
		panic("no value of $User")
	}
}

// 这个函数检查作为其参数的函数在执行时是否会产生panic
func throwPanic(f func()) (b bool) {
	defer func() {
		if x := recover(); x != nil {
			b = true;
		}
	}()
	f()
	return
}