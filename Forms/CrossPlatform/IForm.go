package CrossPlatform

type IForm interface {
	IEmptyControl
	SetTitle(title string)
	SetBorderStyle(style IFormBorder)
	ShowDialog()
}
