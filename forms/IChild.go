package forms

import (
	mb "qq.2564874169/goMiniblink"
	p "qq.2564874169/goMiniblink/platform"
)

type IChild interface {
	IBaseUI

	toControl() p.IControl
	getAnchor() mb.AnchorStyle
}
