package free

type wkePaintBitUpdatedCallback func(wke wkeHandle, param, buf uintptr, rect *wkeRect, width, height int32) uintptr
