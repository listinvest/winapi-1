package user32

type WindowClassEx struct {
	Size            uint32
	Style           uint32
	WindowProcedure uintptr
	ClassExtra      int32
	WindowExtra     int32
	Instance        uintptr
	Icon            uintptr
	Cursor          uintptr
	Background      uintptr
	MenuName        *uint16
	ClassName       *uint16
	IconSmall       uintptr
}

type Message struct {
	Window  uintptr
	Message uint32
	WParam  uintptr
	LParam  uintptr
	Time    uint32
	Point   Point
}

type Point struct {
	X, Y int32
}

type Rect struct {
	Left, Top, Right, Bottom int32
}
