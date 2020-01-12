package CrossPlatform

type IMiniblink interface {
	IControl

	LoadUri(uri string)
	Resize(width, height int)
}
