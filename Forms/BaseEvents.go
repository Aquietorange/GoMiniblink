package Forms

type BaseEvents struct {
	EvLoad   map[string]func(target interface{})
	EvResize map[string]func(target interface{}, width, height int)
	EvMove   map[string]func(target interface{}, x, y int)
}

func (_this *BaseEvents) init() *BaseEvents {
	_this.EvLoad = make(map[string]func(target interface{}))
	_this.EvResize = make(map[string]func(target interface{}, width, height int))
	_this.EvMove = make(map[string]func(target interface{}, x, y int))
	return _this
}
