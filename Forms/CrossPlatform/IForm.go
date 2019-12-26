package CrossPlatform

type IForm interface {
	IEmptyControl
	ShowDialog()
	SetTitle(title string)
	SetBorderStyle(style IFormBorder)
	ShowInTaskbar(isShow bool)
}
