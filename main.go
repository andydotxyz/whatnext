package main

import (
	"time"

	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.NewWithID("xyz.andy.whatnext")
	w := a.NewWindow("What Next")

	u := &ui{pref: a.Preferences()}
	w.SetContent(u.makeUI())
	u.setDate(time.Now())
	w.ShowAndRun()
}
