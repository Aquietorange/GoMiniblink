package forms

import (
	mb "qq.2564874169/miniblink"
	p "qq.2564874169/miniblink/platform"
)

type IChild interface {
	IBaseUI

	toControl() p.IControl
	getAnchor() mb.AnchorStyle
}
