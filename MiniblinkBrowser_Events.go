package GoMiniblink

func (_this *MiniblinkBrowser) defOnRequestBefore(e RequestEvArgs) {
	for _, v := range _this.EvRequestBefore {
		v(_this, e)
	}
}
