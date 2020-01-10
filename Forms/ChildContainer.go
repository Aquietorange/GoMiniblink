package Forms

import "GoMiniblink/CrossPlatform"

type ChildContainer struct {
	container CrossPlatform.IControls
	Childs    []IChild
}

func (_this *ChildContainer) init(controls CrossPlatform.IControls) *ChildContainer {
	_this.container = controls
	return _this
}

func (_this *ChildContainer) AddChild(child IChild) {
	_this.container.AddControl(child.toChild())
	_this.Childs = append(_this.Childs, child)
}

func (_this *ChildContainer) RemoveChild(child IChild) {
	_this.container.RemoveControl(child.toChild())
	for i, n := range _this.Childs {
		if n.GetHandle() == child.GetHandle() {
			_this.Childs = append(_this.Childs[:i], _this.Childs[i+1:]...)
		}
	}
}
