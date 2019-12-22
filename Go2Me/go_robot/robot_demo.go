package go_robot

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

func eachSecKeyBroadInput(word string) {
	robotgo.Sleep(1)
	robotgo.TypeString(word)
	robotgo.Sleep(1)
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
	// robotgo.KeyTap("command", "space")
	// robotgo.Sleep(1)
	// time.Sleep(200 * time.Millisecond)
	//robotgo.KeyTap("space")
	//time.Sleep(200 * time.Millisecond)
	//robotgo.KeyToggle("command", "up")
	//time.Sleep(200 * time.Millisecond)
	// eachSecKeyBroadInput("pwd")
	// robotgo.KeyTap("enter")
	// eachSecKeyBroadInput("brew info golang")
	// robotgo.KeyTap("enter")
	// robotgo.Sleep(5)
	// robotgo.KeyTap("a", "ctrl")
	// robotgo.KeyTap("k", "ctrl")

	// eachSecKeyBroadInput("chrome")
	// robotgo.Sleep(2)
	// time.Sleep(200 * time.Millisecond)
	// robotgo.KeyTap("enter")
	// x, y := robotgo.GetMousePos()
	// fmt.Println("pos: ", x, y)
	// color := robotgo.GetPixelColor(100, 100)
	// fmt.Println("color---- ", color)

	// bitmap := robotgo.CaptureScreen(20, 20, 180, 700)
	// // use `defer robotgo.FreeBitmap(bit)` to free the bitmap
	// defer robotgo.FreeBitmap(bitmap)

	// fmt.Println("...", bitmap)

	// fx, fy := robotgo.FindBitmap(bitmap)
	// fmt.Println("FindBitmap------ ", fx, fy)

	// robotgo.SaveBitmap(bitmap, "test.png")

	// keve := robotgo.AddEvent("k")
	// if keve == 0 {
	// 	fmt.Println("you press... ", "k")
	// }
	count := 10
	for i := 0; i < count; i++ {
		mLeft := robotgo.AddEvent("mLeft")
		if mLeft {
			x, y := robotgo.GetMousePos()
			fmt.Printf("you press %v => x,y [%d, %d]\n", "mouse left button", x, y)
		}
	}

	// fpid, err := robotgo.FindIds("Electron")
	// if err == nil {
	// 	fmt.Println("pids... ", fpid)
	// 	if len(fpid) > 0 {
	// 		firstID := fpid[0]
	// 		pName, err := robotgo.FindName(firstID)
	// 		if err == nil {
	// 			fmt.Printf("fpid %d => pName %v\n", fpid, pName)
	// 			if pName == "Electron" {
	// 				robotgo.ActivePID(firstID)
	// 				robotgo.KeyTap("n", "command")
	// 				robotgo.Sleep(1)
	// 				eachSecKeyBroadInput("github.com/sinlov")
	// 				robotgo.KeyTap("enter")
	// 				robotgo.Sleep(5)
	// 			}
	// 		}
	// 	}
	// }
	// robotgo.ActiveName("chrome")
}
