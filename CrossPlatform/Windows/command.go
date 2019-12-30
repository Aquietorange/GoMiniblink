package Windows

const (
	cmd_invoke      = 100
	cmd_mouse_click = 200
)

type InvokeContext struct {
	fn    func(state interface{})
	state interface{}
	key   string
}
