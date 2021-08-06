package ui

import (
	"math"

	"github.com/bus710/matrixd/src/matrixd/app/common"
	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

func (w *GtkWindow) matrix() (*gtk.Widget, error) {

	// Create EventBox
	eb, err := gtk.EventBoxNew()
	if err != nil {
		return nil, err
	}
	eb.SetSizeRequest(common.WIDTH, common.HEIGHT)
	eb.SetBorderWidth(3)
	// Create Drawing Area
	da, err := gtk.DrawingAreaNew()
	if err != nil {
		return nil, err
	}
	eb.SetSizeRequest(common.WIDTH, common.HEIGHT)

	// Layout the event box
	// https://stackoverflow.com/questions/14002549/how-to-get-mouse-position-at-mouse-click-python-gtk
	eb.Add(da)
	eb.SetHExpand(false)
	eb.SetHAlign(gtk.ALIGN_CENTER)

	eb.Connect("button-press-event", w.catcher)

	// Drawing Area
	// https://github.com/gotk3/gotk3-examples/blob/master/gtk-examples/drawingarea/game.go
	// https://www.cairographics.org/tutorial/
	da.Connect("draw", w.drawer)

	return &eb.Container.Widget, nil
}

func (w *GtkWindow) catcher(win *gtk.Window, ev *gdk.Event) {
	mouseEvent := gdk.EventButtonNewFromEvent(ev)
	mex := math.Round(mouseEvent.X()*100) / 100
	mey := math.Round(mouseEvent.Y()*100) / 100
	// log.Println("event box is clicked at: ", mex, "/", mey)

	// Check if a box is clicked
	for i, c := range w.points {
		if mex > c.X && mex < c.X+c.W {
			if mey > c.Y && mey < c.Y+c.H {
				w.points[i].Clicked = !w.points[i].Clicked
			}
		}
	}
	win.QueueDraw()
}

func (w *GtkWindow) drawer(da *gtk.DrawingArea, cr *cairo.Context) {
	// BG
	cr.Rectangle(0, 0, common.WIDTH, common.HEIGHT) // x, y, w, h
	cr.SetSourceRGB(0.9, 0.9, 0.9)
	cr.Fill()

	for _, c := range w.points {
		// Rect
		cr.SetSourceRGB(c.R, c.G, c.B)
		cr.Rectangle(c.X, c.Y, c.W, c.H) // x, y, w, h
		cr.Fill()
		// Stroke
		if c.Clicked {
			cr.SetSourceRGB(1, 0, 0)
			cr.Rectangle(c.X, c.Y, c.W, c.H) // x, y, w, h
			cr.SetLineWidth(2)
			cr.Stroke()
		} else {
			cr.SetSourceRGB(0.1, 0.3, 0.3)
			cr.Rectangle(c.X, c.Y, c.W, c.H) // x, y, w, h
			cr.SetLineWidth(2)
			cr.Stroke()
		}
	}
}
