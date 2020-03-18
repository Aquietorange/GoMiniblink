package free

type wkePaintBitUpdatedCallback func(wke wkeHandle, param, buf uintptr, rect *wkeRect, width, height int32) uintptr

type wkeNetResponseCallback func(wke wkeHandle, param, utf8Url uintptr, job wkeNetJob) uintptr

type wkeLoadUrlBeginCallback func(wke wkeHandle, param, utf8Url uintptr, job wkeNetJob) uintptr

type wkeJsNativeFunction func(es jsExecState, param uintptr) uintptr
