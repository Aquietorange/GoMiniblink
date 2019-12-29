package Windows

const (
	cmd_invoke = 100
)

type InvokeContext struct {
	fn    func(state interface{})
	state interface{}
	key   string
}
