package Forms

import "GoMiniblink/CrossPlatform"

type Control struct {
	BaseUI

	impl CrossPlatform.IWindow
}

func (_this *Control) Init() *Control {

}
