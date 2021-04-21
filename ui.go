package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type ui struct {
	pref fyne.Preferences

	top1, top2, top3 *widget.Entry
	date             *widget.Label
}

func (u *ui) makeUI() fyne.CanvasObject {
	u.top1 = widget.NewEntry()
	u.top2 = widget.NewEntry()
	u.top3 = widget.NewEntry()
	u.date = widget.NewLabelWithStyle("", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})

	title := container.NewHBox(layout.NewSpacer(),
		widget.NewLabel("Top Priorities"),
		u.date, layout.NewSpacer())
	return container.NewVBox(title,
		container.NewBorder(nil, nil, widget.NewLabel("1:"), nil, u.top1),
		container.NewBorder(nil, nil, widget.NewLabel("2:"), nil, u.top2),
		container.NewBorder(nil, nil, widget.NewLabel("3:"), nil, u.top3))
}

func (u *ui) setDate(t time.Time) {
	dateStr := t.Format("Mon, 02 Jan 2006")
	dateKey := t.Format("20060102")
	u.date.SetText(dateStr)

	u.top1.Unbind()
	s1 := binding.BindPreferenceString(dateKey+".1", u.pref)
	u.top1.Bind(s1)
	u.top1.Validator = nil

	u.top2.Unbind()
	s2 := binding.BindPreferenceString(dateKey+".2", u.pref)
	u.top2.Bind(s2)
	u.top2.Validator = nil

	u.top3.Unbind()
	s3 := binding.BindPreferenceString(dateKey+".3", u.pref)
	u.top3.Bind(s3)
	u.top3.Validator = nil
}
