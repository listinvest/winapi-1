package user32

//Window class styles
const (
	Class_VerticalRedraw      = 0x00000001
	Class_HorizontalRedraw    = 0x00000002
	Class_DoubleClick         = 0x00000008
	Class_UniqueDeviceContext = 0x00000020
	Class_ClassDeviceContext  = 0x00000040
	Class_ParentDeviceContext = 0x00000080
	Class_NoCloseButton       = 0x00000200
	Class_SaveBitmap          = 0x00000800
	Class_AlignClientAreaByte = 0x00001000
	Class_AlignWindowByte     = 0x00002000
	Class_GlobalClass         = 0x00004000
	Class_IME                 = 0x00010000
	Class_DropDownShadow      = 0x00020000
)

//Icon constants
const (
	Icon_Application = 32512
	Icon_Hand        = 32513
	Icon_Question    = 32514
	Icon_Exclamation = 32515
	Icon_Asterisk    = 32516
	Icon_WinLogo     = 32517
	Icon_Warning     = Icon_Exclamation
	Icon_Error       = Icon_Hand
	Icon_Information = Icon_Asterisk
)

//Cursor constants
const (
	Cursor_Arrow                  = 32512
	Cursor_IBeam                  = 32513
	Cursor_Wait                   = 32514
	Cursor_Cross                  = 32515
	Cursor_ArrowUp                = 32516
	Cursor_SizeNorthEastSouthWest = 32642
	Cursor_SizeNorthWestSouthEast = 32643
	Cursor_SizeWestEast           = 32644
	Cursor_SizeNorthSouth         = 32645
	Cursor_SizeAll                = 32646
	Cursor_SlashedCircle          = 32648
	Cursor_Hand                   = 32649
	Cursor_AppStarting            = 32650
	Cursor_Help                   = 32651
	Cursor_Icon                   = 32641
	Cursor_Size                   = 32640
)

//Brashes
const (
	Brush_3D_DarkShadow           = 21
	Brush_3D_Face                 = 15
	Brush_3D_HighLight            = 20
	Brush_3D_Light                = 22
	Brush_ButtonHighLight         = 20
	Brush_3D_Shadow               = 16
	Brush_ActiveBorder            = 10
	Brush_ActiveCaption           = 2
	Brush_AppWorkSpace            = 12
	Brush_Background              = 1
	Brush_Desktop                 = 1
	Brush_ButtonFace              = 15
	Brush_ButtonShadow            = 16
	Brush_ButtonText              = 18
	Brush_CaptionText             = 9
	Brush_GrayText                = 17
	Brush_Highlight               = 13
	Brush_HighlightText           = 14
	Brush_InactiveBorder          = 11
	Brush_InactiveCaption         = 3
	Brush_InactiveCaptionText     = 19
	Brush_InfoBackground          = 24
	Brush_InfoText                = 23
	Brush_Menu                    = 4
	Brush_MenuText                = 7
	Brush_Scrollbar               = 0
	Brush_Window                  = 5
	Brush_WindowFrame             = 6
	Brush_WwindowText             = 8
	Brush_HotLight                = 26
	Brush_GradientActiveCaption   = 27
	Brush_GradientInactiveCaction = 28
)

//Window styles
const (
	Style_Tiled            = 0x00000000
	Style_Popup            = 0x80000000
	Style_Child            = 0x40000000
	Style_Minimized        = 0x20000000
	Style_Visible          = 0x10000000
	Style_Disabled         = 0x08000000
	Style_ClipSiblings     = 0x04000000
	Style_ClipChildren     = 0x02000000
	Style_Maximize         = 0x01000000
	Style_Caption          = 0x00C00000
	Style_Border           = 0x00800000
	Style_DialogFrame      = 0x00400000
	Style_VerticalScroll   = 0x00200000
	Style_HorizontalScroll = 0x00100000
	Style_SystemMenu       = 0x00080000
	Style_Resizeable       = 0x00040000
	Style_Group            = 0x00020000
	Style_TabStop          = 0x00010000
	Style_Minimizeable     = 0x00020000
	Style_Maximizeable     = 0x00010000
	Style_TiledWindow      = Style_Tiled | Style_Caption | Style_SystemMenu | Style_Resizeable | Style_Maximizeable | Style_Minimizeable
	Style_PopupWindow      = Style_Popup | Style_Border | Style_SystemMenu
	Style_ChildWindow      = 0x40000000
)

//Extended window styles
const (
	Style_DialogModal               = 0x00000001
	Style_NoParentNotify            = 0x00000004
	Style_TopMost                   = 0x00000008
	Style_FileDragNDrop             = 0x00000010
	Style_Transparent               = 0x00000020
	Style_MDIChild                  = 0x00000040
	Style_ToolWindow                = 0x00000080
	Style_RaisedEdge                = 0x00000100
	Style_SunkenEdge                = 0x00000200
	Style_ContextHelp               = 0x00000400
	Style_RTL_Alignment             = 0x00001000
	Style_LTR_Alignment             = 0x00000000
	Style_RTL_Reading               = 0x00002000
	Style_LTR_Reading               = 0x00000000
	Style_LeftScrollbar             = 0x00004000
	Style_RightScrollbar            = 0x00000000
	Style_ControlParent             = 0x00010000
	Style_StaticEdge                = 0x00020000
	Style_ApplicationTopLevelWindow = 0x00040000
	Style_OverlappedWindow          = Style_RaisedEdge | Style_SunkenEdge
	Style_PaletteWindow             = Style_RaisedEdge | Style_ToolWindow | Style_TopMost
	Style_Layered                   = 0x00080000
	Style_NoInheritLayout           = 0x00100000
	Style_RTL_Layout                = 0x00400000
	Style_NoActive                  = 0x08000000
)

const (
	Window_Use_Default = ^0x7fffffff
)

const (
	Action_Hide                    = 0
	Action_Show_Normal             = 1
	Action_Show_Minimized          = 2
	Action_Maximize                = 3
	Action_Show_Maximized          = 3
	Action_Show_Normal_Inactive    = 4
	Action_Show                    = 5
	Action_Minimize                = 6
	Action_Shoe_Minimized_Inactive = 7
	Action_Show_Inactive           = 8
	Action_Restore                 = 9
	Action_Show_Default            = 10
	Action_Force_Minimize          = 11
)
