package Windows

const (
	CMD_Invoke = 100
)

type InvokeContext struct {
	fn    func(state interface{})
	state interface{}
	key   string
}
