package forms

import "qq.2564874169/miniblink/platform"

type IChild interface {
	IBaseUI

	toChild() platform.IControl
}
