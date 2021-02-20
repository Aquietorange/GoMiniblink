package MsgBox

import (
	fm "github.com/hujun528/GoMiniblink/forms"
	cs "github.com/hujun528/GoMiniblink/forms/controls"
)

func Show(param fm.MsgBoxParam) fm.MsgBoxResult {
	return cs.App.NewMsgBox().Show(param)
}

func ShowInfo(title, text string) {
	Show(fm.MsgBoxParam{
		Title:  title,
		Text:   text,
		Icon:   fm.MsgBoxIcon_Info,
		Button: fm.MsgBoxButton_Ok,
	})
}

func ShowQuestion(title, text string) fm.MsgBoxResult {
	return Show(fm.MsgBoxParam{
		Title:  title,
		Text:   text,
		Icon:   fm.MsgBoxIcon_Question,
		Button: fm.MsgBoxButton_YesNo,
	})
}

func ShowWarn(title, text string) {
	Show(fm.MsgBoxParam{
		Title:  title,
		Text:   text,
		Icon:   fm.MsgBoxIcon_Warn,
		Button: fm.MsgBoxButton_Ok,
	})
}

func ShowError(title, text string) {
	Show(fm.MsgBoxParam{
		Title:  title,
		Text:   text,
		Icon:   fm.MsgBoxIcon_Error,
		Button: fm.MsgBoxButton_Ok,
	})
}
