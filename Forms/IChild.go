package Forms

import "GoMiniblink/CrossPlatform"

type IChild interface {
	IBaseUI

	toChild() CrossPlatform.IControl
}
