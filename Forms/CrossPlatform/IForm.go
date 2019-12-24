package CrossPlatform

type IForm interface {
	IEmptyControl
	SetTitle(title string)
	ShowDialog()
}