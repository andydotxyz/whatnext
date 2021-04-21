package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.NewWithID("xyz.andy.whatnext")
	w := a.NewWindow("What Next")

	s1 := binding.BindPreferenceString("20210421.1", a.Preferences())
	top1 := widget.NewEntryWithData(s1)
	top1.Validator = nil
	s2 := binding.BindPreferenceString("20210421.2", a.Preferences())
	top2 := widget.NewEntryWithData(s2)
	top2.Validator = nil
	s3 := binding.BindPreferenceString("20210421.3", a.Preferences())
	top3 := widget.NewEntryWithData(s3)
	top3.Validator = nil

	w.SetContent(container.NewVBox(widget.NewLabel("Top Priorities 21 April"),
		container.NewBorder(nil, nil, widget.NewLabel("1:"), nil, top1),
		container.NewBorder(nil, nil, widget.NewLabel("2:"), nil, top2),
		container.NewBorder(nil, nil, widget.NewLabel("3:"), nil, top3)))
	w.ShowAndRun()
}
