package Forms

import (
	"GoMiniblink/Forms/CrossPlatform"
)

type Control struct {
	Impl     CrossPlatform.IEmptyControl
	evLoad   []func(EventSet)
	evResize []func(EventSet, int, int)
}
func (_this *Control) EvLoad(f func(EventSet)) {
	_this.evLoad = append(_this.evLoad, f)
}
func (_this *Control) EvResize(f func(es EventSet, w int, h int)) {
	_this.evResize = append(_this.evResize, f)
}
