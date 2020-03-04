package forms

import (
	mb "qq2564874169/goMiniblink"
	p "qq2564874169/goMiniblink/platform"
)

type IChildContainer interface {
	IBaseUI

	toControls() p.IControls
}

type ChildContainer struct {
	Childs map[uintptr]IChild

	container IChildContainer
	logAnchor map[uintptr]mb.Bound2
}

func (_this *ChildContainer) init(container IChildContainer) *ChildContainer {
	_this.Childs = make(map[uintptr]IChild)
	_this.logAnchor = make(map[uintptr]mb.Bound2)
	_this.container = container
	var bakResize p.WindowResizeProc
	bakResize = container.toControls().SetOnResize(func(e mb.Rect) bool {
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

func (_this *ChildContainer) onAnchor(rect mb.Rect) {
	def := mb.AnchorStyle_Left | mb.AnchorStyle_Top
	for _, n := range _this.Childs {
		anc := n.getAnchor()
		if anc == def {
			continue
		}
		b := _this.logAnchor[n.GetHandle()]
		p := n.GetLocation()
		s := n.GetSize()
		if anc&mb.AnchorStyle_Left != 0 && anc&mb.AnchorStyle_Right != 0 {
			s.Width = rect.Width - b.Left - b.Right
			p.X = b.Left
		} else if anc&mb.AnchorStyle_Right != 0 {
			p.X = rect.Width - b.Right - s.Width
		}
		if anc&mb.AnchorStyle_Top != 0 && anc&mb.AnchorStyle_Bottom != 0 {
			s.Height = rect.Height - b.Top - b.Bottom
			p.Y = b.Top
		} else if anc&mb.AnchorStyle_Bottom != 0 {
			p.Y = rect.Height - b.Bottom - s.Height
		}
		n.SetSize(s.Width, s.Height)
		n.SetLocation(p.X, p.Y)
	}
}

func (_this *ChildContainer) AddChild(child IChild) {
	if _, ok := _this.Childs[child.GetHandle()]; ok == false {
		_this.container.toControls().AddControl(child.toControl())
		_this.Childs[child.GetHandle()] = child
		ps := _this.container.GetSize()
		cp := child.GetLocation()
		cs := child.GetSize()
		rect := mb.Bound2{
			Left:   cp.X,
			Top:    cp.Y,
			Right:  ps.Width - cs.Width - cp.X,
			Bottom: ps.Height - cs.Height - cp.Y,
		}
		_this.logAnchor[child.GetHandle()] = rect
	}
}

func (_this *ChildContainer) RemoveChild(child IChild) {
	if _, ok := _this.Childs[child.GetHandle()]; ok {
		_this.container.toControls().RemoveControl(child.toControl())
		delete(_this.Childs, child.GetHandle())
		delete(_this.logAnchor, child.GetHandle())
	}
}
