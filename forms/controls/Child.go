package controls

import (
	f "qq2564874169/goMiniblink/forms"
	p "qq2564874169/goMiniblink/forms/platform"
)

type Child interface {
	BaseUI

	toControl() p.Control
	getAnchor() f.AnchorStyle
}

type ChildContainer interface {
	BaseUI

	toControls() p.Controls
}

type DefChildContainer struct {
	Childs map[uintptr]Child

	container ChildContainer
	logAnchor map[uintptr]f.Bound2
}

func (_this *DefChildContainer) init(container ChildContainer) *DefChildContainer {
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

func (_this *DefChildContainer) onAnchor(rect f.Rect) {
	def := f.AnchorStyle_Left | f.AnchorStyle_Top
	for _, n := range _this.Childs {
		anc := n.getAnchor()
		if anc == def {
			continue
		}
		b := _this.logAnchor[n.GetHandle()]
		p := n.GetLocation()
		s := n.GetSize()
		if anc&f.AnchorStyle_Left != 0 && anc&f.AnchorStyle_Right != 0 {
			s.Width = rect.Width - b.Left - b.Right
			p.X = b.Left
		} else if anc&f.AnchorStyle_Right != 0 {
			p.X = rect.Width - b.Right - s.Width
		}
		if anc&f.AnchorStyle_Top != 0 && anc&f.AnchorStyle_Bottom != 0 {
			s.Height = rect.Height - b.Top - b.Bottom
			p.Y = b.Top
		} else if anc&f.AnchorStyle_Bottom != 0 {
			p.Y = rect.Height - b.Bottom - s.Height
		}
		n.SetSize(s.Width, s.Height)
		n.SetLocation(p.X, p.Y)
	}
}

func (_this *DefChildContainer) AddChild(child Child) {
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

func (_this *DefChildContainer) RemoveChild(child Child) {
	if _, ok := _this.Childs[child.GetHandle()]; ok {
		_this.container.toControls().RemoveControl(child.toControl())
		delete(_this.Childs, child.GetHandle())
		delete(_this.logAnchor, child.GetHandle())
	}
}
