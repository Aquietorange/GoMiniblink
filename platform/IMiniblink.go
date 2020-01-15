package platform

type IMiniblink interface {
	IControl

	LoadUri(uri string)
}
