package bridge

import fm "github.com/hujun528/GoMiniblink/forms"

type MsgBox interface {
	Show(param fm.MsgBoxParam) fm.MsgBoxResult
}
