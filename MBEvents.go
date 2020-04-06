package GoMiniblink

type RequestEvArgs interface {
	GetUrl() string
	GetMethod() string
	SetData([]byte)
	GetData() []byte
	SetCancel(b bool)
	IsCancel() bool
}
