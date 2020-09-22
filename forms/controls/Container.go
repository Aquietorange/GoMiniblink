package controls

import (
	f "qq2564874169/goMiniblink/forms"
	p "qq2564874169/goMiniblink/forms/platform"
)

type Container interface {
	GUI

	toControls() p.Controls
}

type Child interface {
	GUI

	toControl() p.Control
	getAnchor() f.AnchorStyle
}

type BaseContainer struct {
	Childs map[uintptr]Child

	container Container
	logAnchor map[uintptr]f.Bound2
}

func (_this *BaseContainer) Init(container Container) *BaseContainer {
	_this.Childs = make(map[uintptr]Child)
	_this.logAnchor = make(map[uintptr]f.Bound2)
	_this.container = container
	var bakResize p.WindowResizeProc
	bakResize = container.toControls().SetOnResize(func(e f.Rect) bool {
		b := false
		if bakResize != nil {
			b = bakResize(e)
		}
		if !b {
			_this.onAnchor(e)
		}
		return b
	})
	return _this
}

func (_this *BaseContainer) onAnchor(rect f.Rect) {
	def := f.AnchorStyle_Left | f.AnchorStyle_Top
	for _, n := range _this.Childs {
		anc := n.getAnchor()
		if anc == def {
			continue
		}
		b := _this.logAnchor[n.GetHandle()]
		pos := n.GetLocation()
		sz := n.GetSize()
		if anc&f.AnchorStyle_Left != 0 && anc&f.AnchorStyle_Right != 0 {
			sz.Width = rect.Width - b.Left - b.Right
			pos.X = b.Left
		} else if anc&f.AnchorStyle_Right != 0 {
			pos.X = rect.Width - b.Right - sz.Width
		}
		if anc&f.AnchorStyle_Top != 0 && anc&f.AnchorStyle_Bottom != 0 {
			sz.Height = rect.Height - b.Top - b.Bottom
			pos.Y = b.Top
		} else if anc&f.AnchorStyle_Bottom != 0 {
			pos.Y = rect.Height - b.Bottom - sz.Height
		}
		n.SetSize(sz.Width, sz.Height)
		n.SetLocation(pos.X, pos.Y)
	}
}

func (_this *BaseContainer) AddChild(child Child) {
	if _, ok := _this.Childs[child.GetHandle()]; ok == false {
		_this.container.toControls().AddControl(child.toControl())
		_this.Childs[child.GetHandle()] = child
		ps := _this.container.GetSize()
		cp := child.GetLocation()
		cs := child.GetSize()
		rect := f.Bound2{
			Left:   cp.X,
			Top:    cp.Y,
			Right:  ps.Width - cs.Width - cp.X,
			Bottom: ps.Height - cs.Height - cp.Y,
		}
		_this.logAnchor[child.GetHandle()] = rect
	}
}

func (_this *BaseContainer) RemoveChild(child Child) {
	if _, ok := _this.Childs[child.GetHandle()]; ok {
		_this.container.toControls().RemoveControl(child.toControl())
		delete(_this.Childs, child.GetHandle())
		delete(_this.logAnchor, child.GetHandle())
	}
}
