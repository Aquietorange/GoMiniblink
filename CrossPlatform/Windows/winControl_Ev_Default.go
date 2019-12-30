package Windows

import MB "GoMiniblink"

func (_this *winControl) defOnMouseMove(args MB.MouseEvArgs) {
	for _, v := range _this.evMouseMove {
		v(_this, args)
	}
}
