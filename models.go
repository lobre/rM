package rm

import (
	"os"
)

// Width and Height of the device
const (
	Width  int = 1404
	Height int = 1872
)

type BrushColor int
type BrushType int
type BrushSize float32

// Brush colors
const (
	Black BrushColor = 0
	Grey  BrushColor = 1
	White BrushColor = 2
)

// Brush types
const (
	Brush       BrushType = 0
	PencilTilt  BrushType = 1
	Pen         BrushType = 2
	Marker      BrushType = 3
	Fineliner   BrushType = 4
	Highlighter BrushType = 5
	Eraser      BrushType = 6
	PencilSharp BrushType = 7
	EraseArea   BrushType = 8
)

// Brush sizes
const (
	Small  BrushSize = 1.875
	Medium BrushSize = 2.0
	Large  BrushSize = 2.125
)

type Point struct {
	X           float32
	Y           float32
	PenPressure float32
	XRotation   float32
	YRotation   float32
}

type Line struct {
	BrushType  BrushType
	BrushColor BrushColor
	BrushSize  BrushSize
	Points     []Point
}

type Layer struct {
	Lines []Line
}

type Page struct {
	Layers []Layer
	//template  string
	//thumbnail image.Image
}

type ContentTransform struct {
	M11 int `json:"m11"`
	M12 int `json:"m12"`
	M13 int `json:"m13"`
	M21 int `json:"m21"`
	M22 int `json:"m22"`
	M23 int `json:"m23"`
	M31 int `json:"m31"`
	M32 int `json:"m32"`
	M33 int `json:"m33"`
}

type ContentExtraMetadata struct {
	LastColor      string `json:"LastColor"`
	LastTool       string `json:"LastTool"`
	ThicknessScale string `json:"ThicknessScale"`
}

type Content struct {
	ExtraMetadata  ContentExtraMetadata `json:"extraMetadata"`
	FileType       string               `json:"fileType"`
	FontName       string               `json:"fontName"`
	LastOpenedPage int                  `json:"lastOpenedPage"`
	LineHeight     int                  `json:"lineHeight"`
	Margins        int                  `json:"margins"`
	PageCount      int                  `json:"pageCount"`
	TextScale      int                  `json:"textScale"`
	Transform      ContentTransform     `json:"transform"`
}

// Notebook parsed from the reMarkable
type Notebook struct {
	Name string

	Id      string
	Pages   []Page
	Content Content
	Pdf     os.File
	Epub    os.File
	Hash    string
}

const header = "reMarkable lines with selections and layers"