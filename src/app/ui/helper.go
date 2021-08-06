package ui

import (
	"math"
	"math/rand"

	"github.com/bus710/matrixd/src/matrixd/app/common"
	"github.com/bus710/matrixd/src/matrixd/app/matrix"
)

func (w *GtkWindow) setPoints() {
	// Set the default values
	y_offset := 30.0
	for i := 0; i < 8; i++ {
		x_offset := 30.0
		for j := 0; j < 8; j++ {
			index := i*8 + j
			// Position
			w.points[index].X = x_offset
			w.points[index].Y = y_offset
			w.points[index].W = 20
			w.points[index].H = 20
			// Color
			w.points[index].R = 0.0
			w.points[index].G = 0.0
			w.points[index].B = 0.0
			// State
			w.points[index].Clicked = false
			// Next x
			x_offset += 30
		}
		// Next y
		y_offset += 30
	}

	w.lastSlide.R = 0.0
	w.lastSlide.G = 0.0
	w.lastSlide.B = 0.0
}

func (w *GtkWindow) setAll() {
	for i := range w.points {
		w.points[i].Clicked = true
		w.points[i].R = 0.0
		w.points[i].G = 0.0
		w.points[i].B = 0.0
	}
}

func (w *GtkWindow) setNone() {
	for i := range w.points {
		w.points[i].Clicked = false
	}
}

func (w *GtkWindow) setColor(req common.Request) {
	for i := range w.points {
		if w.points[i].Clicked {
			if req.CMD == "R" && w.points[i].Clicked {
				w.points[i].R = req.R / 64
				w.points[i].G = w.lastSlide.G
				w.points[i].B = w.lastSlide.B
			}
			if req.CMD == "G" && w.points[i].Clicked {
				w.points[i].R = w.lastSlide.R
				w.points[i].G = req.G / 64
				w.points[i].B = w.lastSlide.B
			}
			if req.CMD == "B" && w.points[i].Clicked {
				w.points[i].R = w.lastSlide.R
				w.points[i].G = w.lastSlide.G
				w.points[i].B = req.B / 64
			}
		}
	}
}

func (w *GtkWindow) setRandom() {
	for i := range w.points {
		w.points[i].Clicked = false
		w.points[i].R = math.Round(rand.Float64()*100) / 100
		w.points[i].G = math.Round(rand.Float64()*100) / 100
		w.points[i].B = math.Round(rand.Float64()*100) / 100
	}
}

func (w *GtkWindow) setSubmit() {
	d := common.MatrixData{}
	for i, p := range w.points {
		d.R[i] = uint8(p.R * 255)
		d.G[i] = uint8(p.G * 255)
		d.B[i] = uint8(p.B * 255)
	}
	matrix.Push(&d)
}

func newRequest(name string, r float64, g float64, b float64) (common.Request, error) {
	req := common.Request{
		CMD: name,
		R:   r,
		G:   g,
		B:   b,
	}
	return req, nil
}
