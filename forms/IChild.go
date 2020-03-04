package forms

import (
	mb "qq2564874169/goMiniblink"
	p "qq2564874169/goMiniblink/platform"
)

type IChild interface {
	IBaseUI

	toControl() p.IControl
	getAnchor() mb.AnchorStyle
}
