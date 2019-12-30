package Forms

import "GoMiniblink"

type BaseEvents struct {
	EvLoad       map[string]func(target interface{})
	EvResize     map[string]func(target interface{}, width, height int)
	EvMove       map[string]func(target interface{}, x, y int)
	EvMouseMove  map[string]func(target interface{}, e GoMiniblink.MouseEvArgs)
	EvMouseDown  map[string]func(target interface{}, e GoMiniblink.MouseEvArgs)
	EvMouseUp    map[string]func(target interface{}, e GoMiniblink.MouseEvArgs)
	EvMouseWheel map[string]func(target interface{}, e GoMiniblink.MouseEvArgs)
	EvMouseClick map[string]func(target interface{}, e GoMiniblink.MouseEvArgs)
}

func (_this *BaseEvents) init() *BaseEvents {
	_this.EvLoad = make(map[string]func(target interface{}))
	_this.EvResize = make(map[string]func(target interface{}, width, height int))
	_this.EvMove = make(map[string]func(target interface{}, x, y int))
	_this.EvMouseMove = make(map[string]func(target interface{}, args GoMiniblink.MouseEvArgs))
	_this.EvMouseDown = make(map[string]func(target interface{}, args GoMiniblink.MouseEvArgs))
	_this.EvMouseUp = make(map[string]func(target interface{}, args GoMiniblink.MouseEvArgs))
	_this.EvMouseWheel = make(map[string]func(target interface{}, args GoMiniblink.MouseEvArgs))
	_this.EvMouseClick = make(map[string]func(target interface{}, args GoMiniblink.MouseEvArgs))
	return _this
}
