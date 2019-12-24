package main

import (
	"GoMiniblink/Forms"
	"GoMiniblink/Forms/CrossPlatform/Windows"
)

func main() {
	Forms.Provider = new(Windows.Provider).Init()
	var frm = new(Forms.Form).Init()
	frm.SetSize(300, 500)
	frm.EvMove = append(frm.EvMove, func(f *Forms.Form, x, y int) {
		println("x", x, "y", y)
	})
	Forms.Run(frm)
}
