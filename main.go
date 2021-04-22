//go:generate fyne bundle -o bundled.go Icon.png

package main

import (
	"time"

	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.NewWithID("xyz.andy.whatnext")
	a.SetIcon(resourceIconPng)
	w := a.NewWindow("What Next")

	u := &ui{pref: a.Preferences()}
	w.SetContent(u.makeUI())
	u.setDate(time.Now())
	w.ShowAndRun()
}
