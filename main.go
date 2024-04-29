//go:generate fyne bundle -o bundled.go Icon.png
//go:generate fyne bundle -o bundled.go -append img/day.svg
//go:generate fyne bundle -o bundled.go -append img/week.svg
//go:generate fyne bundle -o bundled.go -append img/quarter.svg
//go:generate fyne bundle -o bundled.go -append img/today.svg
//go:generate fyne bundle -o bundled.go -append img/habit.svg

package main

import (
	"time"

	"fyne.io/cloud"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.NewWithID("xyz.andy.whatnext")
	cloud.Enable(a)
	a.SetIcon(resourceIconPng)
	w := a.NewWindow("What Next")

	u := &ui{app: a}
	w.SetContent(u.makeUI())
	w.Resize(fyne.NewSize(400, 300))
	u.setDate(time.Now())
	w.ShowAndRun()
}
