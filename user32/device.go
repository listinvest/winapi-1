package user32

import "github.com/Magic-Library/winapi/tools"

type DeviceSettings struct {
	Name                 [tools.DeviceNameLength]uint16
	SpecificationVersion uint16
	DriverVersion        uint16
	Size                 uint16
	DriverExtra          uint16
	Fields               uint32
	Orientation          int16
	PaperSize            int16
	PaperLength          int16
	PaperWidth           int16
	Scale                int16
	Copies               int16
	DefaultSource        int16
	PrintQuality         int16
	Color                int16
	Duplex               int16
	YResolution          int16
	TrueTypeFont         int16
	Collate              int16
	FormName             [tools.DeviceNameLength]uint16
	LogPixels            uint16
	BitsPerPx            uint32
	PxWidth              uint32
	PxHeight             uint32
	DisplayFlags         uint32
	DisplayFrequency     uint32
	ICMMethod            uint32
	ICMIntent            uint32
	MediaType            uint32
	DitherType           uint32
	Reserved1            uint32
	Reserved2            uint32
	PanningWidth         uint32
	PanningHeight        uint32
}
