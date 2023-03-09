/*
如何使用 goland debug goroutine
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
)

type Tester interface {
	test1()
	test2()
	test3()
}

type T struct {
	a int
	s *sync.WaitGroup
}

func main() {
	fmt.Println("main start!")
	s := &sync.WaitGroup{}
	s.Add(3)
	t := T{a: 1, s: s} // 断点
	// bebug 到这里时,需要在帧栏目中选择 main.test1 的协程帧,不然不会 debug到 test1 函数内部的.或者光标放入 test1 函数内部,点击运行到光标处也行
	go t.test1() // 断点
	go t.test2()
	go t.test3() // 断点
	s.Wait()
	fmt.Println("main end!")
}

// 如何进入 goroutine
func (t *T) test1() {
	defer func() {
		t.s.Done()
	}()
	t.a++ // 断点
	fmt.Println("test1", t.a)
}

// 异常断点
func (t *T) test2() {
	defer func() {
		// 发生宕机时，获取panic传递的上下文并打印
		err := recover()
		switch err.(type) {
		case runtime.Error: // 运行时错误
			fmt.Println("runtime error:", err)
		default: // 非运行时错误
			fmt.Println("error:", err)
		}
		t.s.Done()
	}()

	t.a++
	panic("制造一个错误")
	fmt.Println("test2", t.a)
}

// 如何跳过循环
func (t *T) test3() {
	fmt.Println("test start!") // 断点
	for i := 0; i < 30; i++ {  // debug 到一些循环时,不想跟踪每个循环运行,可以点击 步出(F8)跳到下一个 debug 点
		fmt.Println("test3", i)
	}
	fmt.Println("test end!") // 断点
	t.s.Done()
}

/*
调试心得:
1.要是协程帧打不开,那可能是帧太多了,可以在 goland 的设置中设置协程数限制
2.有时候上面的方式切换不到想看的协程,可以在协程运行附近(go XXX()处),试试使用运行到光标处(光标放到协程函数中)
3.有时候单步运行到了循环中,同样可以使用运行到光标处(光标点击到循环外)或者点击回复程序按钮(最好在循环外打一个端点,点击恢复程序按钮会运行到此处)
3.https://www.jetbrains.com/help/go/debugging-code.html 官方教程,但感觉没说很细
*/
