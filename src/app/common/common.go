package common

// Gtk drawing area size
const WIDTH = 295
const HEIGHT = 295

// MatrixData ...
type MatrixData struct {
	R [64]uint8 `json:"R"`
	G [64]uint8 `json:"G"`
	B [64]uint8 `json:"B"`
}

// Request is used between the gtk components (i.e. button (or slider) => drawing area)
type Request struct {
	CMD string
	R   float64
	G   float64
	B   float64
}

// LastSlide ...
type LastSlide struct {
	R float64
	G float64
	B float64
}

// Point ...
type Point struct {
	// Coordinate
	X float64
	Y float64
	W float64
	H float64
	// Color
	R float64
	G float64
	B float64
	// State
	Clicked bool
}
