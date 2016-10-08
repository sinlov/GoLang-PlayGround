package go_robot

import (
	"github.com/go-vgo/robotgo"
	"time"
)

func eachSecKeyBroadInput(word string) {
	time.Sleep(200 * time.Millisecond)
	robotgo.TypeString(word)
	time.Sleep(200 * time.Millisecond)
}

func OSX_Open_Chrome_To_GitHub() {
	//time.Sleep(200 * time.Millisecond)
	//robotgo.DragMouse(360, 799)
	//time.Sleep(200 * time.Millisecond)
	//time.Sleep(200 * time.Millisecond)
	//robotgo.DragMouse(360, 751)
	//time.Sleep(200 * time.Millisecond)
	//robotgo.MouseClick()
	//time.Sleep(200 * time.Millisecond)
	//robotgo.DragMouse(270, 80)
	//time.Sleep(200 * time.Millisecond)
	//robotgo.MouseClick()
	robotgo.KeyToggle("command", "down")
	time.Sleep(200 * time.Millisecond)
	robotgo.KeyTap("space")
	time.Sleep(200 * time.Millisecond)
	robotgo.KeyToggle("command", "up")
	time.Sleep(200 * time.Millisecond)
	//eachSecKeyBroadInput("https://
}
