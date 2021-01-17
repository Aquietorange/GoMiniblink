package windows

import (
	fm "gitee.com/aochulai/GoMiniblink/forms"
	w "gitee.com/aochulai/GoMiniblink/forms/windows/win32"
)

type winMsgBox struct {
}

func (_this *winMsgBox) init() *winMsgBox {
	return _this
}

func (_this *winMsgBox) Show(param fm.MsgBoxParam) fm.MsgBoxResult {
	flag := w.MB_TASKMODAL
	switch param.Icon {
	case fm.MsgBoxIcon_Info:
		flag |= w.MB_ICONINFORMATION
	case fm.MsgBoxIcon_Warn:
		flag |= w.MB_ICONWARNING
	case fm.MsgBoxIcon_Error:
		flag |= w.MB_ICONERROR
	case fm.MsgBoxIcon_Question:
		flag |= w.MB_ICONQUESTION
	}
	switch param.Button {
	case fm.MsgBoxButton_Ok:
		flag |= w.MB_OK
	case fm.MsgBoxButton_YesNo:
		flag |= w.MB_YESNO
	case fm.MsgBoxButton_YesNoCancel:
		flag |= w.MB_YESNOCANCEL
	case fm.MsgBoxButton_AbortRetryIgnore:
		flag |= w.MB_ABORTRETRYIGNORE
	case fm.MsgBoxButton_OkCancel:
		flag |= w.MB_OKCANCEL
	case fm.MsgBoxButton_RetryCancel:
		flag |= w.MB_RETRYCANCEL
	}

	rs := w.MessageBox(w.HWND(0), sto16(param.Text), sto16(param.Title), uint32(flag))
	return fm.MsgBoxResult(rs)
}
