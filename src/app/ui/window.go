package ui

import (
	"sync"

	"github.com/bus710/matrixd/src/matrixd/app/common"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

var Window GtkWindow

// var lastSlide common.LastSlide

type GtkWindow struct {
	win *gtk.Window

	// Data for the GUI matirx
	points    []common.Point
	lastSlide common.LastSlide
	// Join
	wait *sync.WaitGroup
	// Channels
	chanRequest chan common.Request
	// GLib Timeout
	tos glib.SourceHandle
}

func (w *GtkWindow) Init(wait *sync.WaitGroup, windowCloseIndicator chan bool) error {
	var err error
	// assign waig group
	w.wait = wait
	// Init states
	w.points = make([]common.Point, 64)
	w.setPoints()
	// Init channel
	w.chanRequest = make(chan common.Request)

	// Initialize GTK without parsing any command line arguments.
	gtk.Init(nil)
	// Create a new toplevel window, set its title, and connect it to the
	// "destroy" signal to exit the GTK main loop when it is destroyed.
	w.win, err = gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		return nil
	}
	w.win.SetTitle("")
	w.win.Connect("destroy", func() {
		windowCloseIndicator <- true
		gtk.MainQuit()
		w.wait.Done()
	})
	// Layout body
	widget, err := w.windowWidget()
	if err != nil {
		return nil
	}
	// Add the label to the window.
	w.win.Add(widget)

	// Set the default window size.
	w.win.SetDefaultSize(400, 640)
	w.win.SetSizeRequest(400, 640)
	w.win.SetResizable(false)

	w.win.Connect("destroy", gtk.MainQuit)

	return nil
}

func (w *GtkWindow) Run() {
	// Recursively show all widgets contained in this window.
	w.win.ShowAll()

	// Begin executing the GTK main loop.
	// This blocks until gtk.MainQuit() is run.
	gtk.Main()
}

func (w *GtkWindow) Close() {
	w.win.Destroy()
}
