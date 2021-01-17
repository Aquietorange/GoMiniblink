package bridge

import fm "gitee.com/aochulai/GoMiniblink/forms"

type MsgBox interface {
	Show(param fm.MsgBoxParam) fm.MsgBoxResult
}
