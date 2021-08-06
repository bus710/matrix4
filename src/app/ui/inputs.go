package ui

import (
	"github.com/gotk3/gotk3/gtk"
)

func (w *GtkWindow) inputs() (*gtk.Widget, *gtk.Widget, error) {
	// Create a v box to contain the sliders
	vBoxWithSliders, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 3)
	if err != nil {
		return nil, nil, err
	}

	slider01, err := w.slider("R")
	if err != nil {
		return nil, nil, err
	}
	slider02, err := w.slider("G")
	if err != nil {
		return nil, nil, err
	}
	slider03, err := w.slider("B")
	if err != nil {
		return nil, nil, err
	}

	// Layout the v box
	vBoxWithSliders.SetVAlign(gtk.ALIGN_END)
	vBoxWithSliders.PackStart(slider01, true, true, 1)
	vBoxWithSliders.PackStart(slider02, true, true, 1)
	vBoxWithSliders.PackStart(slider03, true, true, 1)

	// Create a v box to contain the buttons
	vBoxWithButtons, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 3)
	if err != nil {
		return nil, nil, err
	}

	buttons01, err := w.buttons("Select", "All", "None")
	if err != nil {
		return nil, nil, err
	}

	buttons02, err := w.buttons("Effect", "Random", "Submit")
	if err != nil {
		return nil, nil, err
	}

	// Layout the h box
	vBoxWithButtons.SetVAlign(gtk.ALIGN_END)
	vBoxWithButtons.PackStart(buttons01, false, false, 3)
	vBoxWithButtons.PackStart(buttons02, false, false, 3)
	// vBoxWithButtons.SetMarginEnd(20)

	return &vBoxWithSliders.Container.Widget, &vBoxWithButtons.Container.Widget, nil
}

func (w *GtkWindow) slider(name string) (*gtk.Widget, error) {
	hBoxSlider, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 3)
	if err != nil {
		return nil, err
	}
	// Create labels for RGB
	label, err := gtk.LabelNew(name)
	if err != nil {
		return nil, err
	}
	// Create a new scale
	scale, err := gtk.ScaleNewWithRange(gtk.ORIENTATION_HORIZONTAL, 0, 63, 1)
	if err != nil {
		return nil, err
	}

	scale.SetSizeRequest(240, 0)
	scale.SetName(name)
	hBoxSlider.SetSizeRequest(300, 0)
	hBoxSlider.SetHExpand(true)
	hBoxSlider.SetHAlign(gtk.ALIGN_CENTER)
	hBoxSlider.SetVAlign(gtk.ALIGN_END)
	hBoxSlider.PackStart(label, false, false, 10)
	hBoxSlider.PackStart(scale, false, false, 10)
	hBoxSlider.SetMarginStart(10)

	sliderChanged := func(slider *gtk.Scale) {
		name, _ := slider.GetName()
		rng := slider.GetValue()
		switch name {
		case "R":
			w.lastSlide.R = rng / 64
			req, _ := newRequest(name, rng, w.lastSlide.G, w.lastSlide.B)
			w.chanRequest <- req
		case "G":
			w.lastSlide.G = rng / 64
			req, _ := newRequest(name, w.lastSlide.R, rng, w.lastSlide.B)
			w.chanRequest <- req
		case "B":
			w.lastSlide.B = rng / 64
			req, _ := newRequest(name, w.lastSlide.R, w.lastSlide.G, rng)
			w.chanRequest <- req
		}
	}
	scale.Connect("value-changed", sliderChanged)

	return &hBoxSlider.Container.Widget, nil
}

func (w *GtkWindow) buttons(name01, name02, name03 string) (*gtk.Widget, error) {
	hBoxSelect, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 3)
	if err != nil {
		return nil, err
	}
	// Create items for select row
	labelSelect, err := gtk.LabelNew(name01)
	if err != nil {
		return nil, err
	}
	button01, err := gtk.ButtonNewWithLabel(name02)
	if err != nil {
		return nil, err
	}
	button02, err := gtk.ButtonNewWithLabel(name03)
	if err != nil {
		return nil, err
	}

	labelSelect.SetMarginEnd(10)
	button01.SetSizeRequest(100, 0)
	button02.SetSizeRequest(100, 0)

	hBoxSelect.SetHAlign(gtk.ALIGN_CENTER)
	hBoxSelect.PackStart(labelSelect, true, false, 10)
	hBoxSelect.PackStart(button01, false, false, 3)
	hBoxSelect.PackStart(button02, false, false, 3)

	// Declare handler function
	// To be connected, the function should have the same widget pointer as arg
	buttonClicked := func(btn *gtk.Button) {
		// name, _ := btn.GetName()
		buttonLabel, _ := btn.GetLabel()
		// log.Println("clicked", name, buttonLabel)
		req, _ := newRequest(buttonLabel, 0.0, 0.0, 0.0)
		w.chanRequest <- req
	}
	// Connect signal and widgets
	button01.Connect("clicked", buttonClicked)
	button02.Connect("clicked", buttonClicked)

	return &hBoxSelect.Container.Widget, nil
}
