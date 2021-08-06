package ui

import (
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func (w *GtkWindow) windowWidget() (*gtk.Widget, error) {
	// Create the biggest widget to contain everything (label + scale + horizontal box)
	vb, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 10)
	if err != nil {
		return nil, err
	}

	// Create a matrix wrapped by an event box
	matrixBox, err := w.matrix()
	if err != nil {
		return nil, err
	}
	// Create a set of inputs like sliders and buttons
	vBoxSliders, vBoxButtons, err := w.inputs()
	if err != nil {
		return nil, err
	}

	// Layout the v box
	vb.SetMarginStart(10)
	vb.SetMarginEnd(10)
	vb.SetMarginTop(10)
	vb.SetMarginBottom(10)
	vb.PackStart(matrixBox, false, false, 10)
	vb.PackStart(vBoxSliders, false, false, 10)
	vb.PackStart(vBoxButtons, false, false, 10)

	// This go routine exposes the gtk loop
	// so we can access to the loop via go channel
	go func() {
		for req := range w.chanRequest {
			// WHY IdleAdd? to safely put a request to Gtk main loop
			// https://github.com/gotk3/gotk3/issues/492
			glib.IdleAdd(func() {
				// To change the points variable,
				// those SetAll/SetNone functions should be called in this scope.
				switch req.CMD {
				case "All":
					w.setAll()
				case "None":
					w.setNone()
				case "R", "G", "B":
					w.setColor(req)
				case "Random":
					w.setRandom()
				case "Submit":
					w.setSubmit()
				}

				matrixBox.QueueDraw()
			})
		}
	}()

	// Glib timer for demo
	w.tos = glib.TimeoutAdd(uint(1000), func() bool {
		// fmt.Println("timed out")
		return true
	})

	return &vb.Container.Widget, nil
}
