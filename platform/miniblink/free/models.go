package free

type wkeRequestType int

const (
	wkeRequestType_Unknow = 1
	wkeRequestType_Get    = 2
	wkeRequestType_Post   = 3
	wkeRequestType_Put    = 4
)

type wkeKeyFlags int

const (
	wkeKeyFlags_Extend wkeKeyFlags = 0x0100
	wkeKeyFlags_Repeat wkeKeyFlags = 0x4000
)

type wkeRect struct {
	x, y, w, h int32
}

type wkeNetJob uintptr

type wkeMouseFlags int

const (
	wkeMouseFlags_None    wkeMouseFlags = 0
	wkeMouseFlags_LBUTTON wkeMouseFlags = 0x01
	wkeMouseFlags_RBUTTON wkeMouseFlags = 0x02
	wkeMouseFlags_SHIFT   wkeMouseFlags = 0x04
	wkeMouseFlags_CONTROL wkeMouseFlags = 0x08
	wkeMouseFlags_MBUTTON wkeMouseFlags = 0x10
)

type wkeCursorType int

const (
	wkeCursorType_Pointer wkeCursorType = iota
	wkeCursorType_Cross
	wkeCursorType_Hand
	wkeCursorType_IBeam
	wkeCursorType_Wait
	wkeCursorType_Help
	wkeCursorType_EastResize
	wkeCursorType_NorthResize
	wkeCursorType_NorthEastResize
	wkeCursorType_NorthWestResize
	wkeCursorType_SouthResize
	wkeCursorType_SouthEastResize
	wkeCursorType_SouthWestResize
	wkeCursorType_WestResize
	wkeCursorType_NorthSouthResize
	wkeCursorType_EastWestResize
	wkeCursorType_NorthEastSouthWestResize
	wkeCursorType_NorthWestSouthEastResize
	wkeCursorType_ColumnResize
	wkeCursorType_RowResize
	wkeCursorType_MiddlePanning
	wkeCursorType_EastPanning
	wkeCursorType_NorthPanning
	wkeCursorType_NorthEastPanning
	wkeCursorType_NorthWestPanning
	wkeCursorType_SouthPanning
	wkeCursorType_SouthEastPanning
	wkeCursorType_SouthWestPanning
	wkeCursorType_WestPanning
	wkeCursorType_Move
	wkeCursorType_VerticalText
	wkeCursorType_Cell
	wkeCursorType_ContextMenu
	wkeCursorType_Alias
	wkeCursorType_Progress
	wkeCursorType_NoDrop
	wkeCursorType_Copy
	wkeCursorType_None
	wkeCursorType_NotAllowed
	wkeCursorType_ZoomIn
	wkeCursorType_ZoomOut
	wkeCursorType_Grab
	wkeCursorType_Grabbing
	wkeCursorType_Custom
)
