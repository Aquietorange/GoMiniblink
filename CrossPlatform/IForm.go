package CrossPlatform

import MB "GoMiniblink"

type IForm interface {
	IWindow
	ShowDialog()
	SetTitle(title string)
	SetBorderStyle(style MB.FormBorder)
	ShowInTaskbar(isShow bool)

	SetSize(w int, h int)
	SetOnResize(func(w int, h int))

	SetLocation(x int, y int)
	SetOnMove(func(x int, y int))

	SetState(state MB.FormState)
	SetOnState(func(MB.FormState))

	SetMaximizeBox(isShow bool)
	SetMinimizeBox(isShow bool)
}
