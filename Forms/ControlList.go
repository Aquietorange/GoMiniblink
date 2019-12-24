package Forms

import "GoMiniblink/Collections"

type controlList struct {
	data     *Collections.List
	onAdd    []func(s *controlList, e *Control)
	onRemove []func(s *controlList, e *Control)
}

func (_this *controlList) Init() *controlList {
	var list = controlList{
		data: new(Collections.List).Init(),
	}
	return &list
}

func (_this *controlList) getData() *Collections.List {
	if _this.data == nil {
		_this.data = new(Collections.List)
	}
	return _this.data
}

func (_this *controlList) Add(control *Control) int {
	var i = _this.getData().Add(control)
	for _, v := range _this.onAdd {
		v(_this, control)
	}
	return i
}

func (_this *controlList) Remove(control *Control) *Control {
	if _this.getData().Remove(control) {
		for _, v := range _this.onRemove {
			v(_this, control)
		}
		return control
	}
	return nil
}

func (_this *controlList) Count() int {
	return _this.getData().Count()
}

func (_this *controlList) Get(index int) *Control {
	if v, ok := _this.getData().GetItem(index).(*Control); ok {
		return v
	}
	return nil
}
