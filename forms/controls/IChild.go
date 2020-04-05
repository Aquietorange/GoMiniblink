package controls

import (
	f "qq2564874169/goMiniblink/forms"
	p "qq2564874169/goMiniblink/forms/platform"
)

type IChild interface {
	IBaseUI

	toControl() p.IControl
	getAnchor() f.AnchorStyle
}
