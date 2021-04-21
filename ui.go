package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

type ui struct {
	pref fyne.Preferences

	goal, top1, top2, top3 *widget.Entry
	date                   *widget.Label
}

func (u *ui) makeUI() fyne.CanvasObject {
	u.goal = widget.NewEntry()
	u.top1 = widget.NewEntry()
	u.top2 = widget.NewEntry()
	u.top3 = widget.NewEntry()
	u.date = widget.NewLabelWithStyle("", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

	return container.NewVBox(u.date,
		widget.NewLabel("Goal"), u.goal,
		widget.NewLabel("Targets"),
		container.NewBorder(nil, nil, widget.NewLabel("1:"), nil, u.top1),
		container.NewBorder(nil, nil, widget.NewLabel("2:"), nil, u.top2),
		container.NewBorder(nil, nil, widget.NewLabel("3:"), nil, u.top3))
}

func (u *ui) setDate(t time.Time) {
	dateStr := t.Format("Mon, 02 Jan 2006")
	dateKey := t.Format("20060102")
	u.date.SetText(dateStr)

	u.goal.Unbind()
	g := binding.BindPreferenceString(dateKey+".goal", u.pref)
	u.goal.Bind(g)
	u.goal.Validator = nil

	u.top1.Unbind()
	s1 := binding.BindPreferenceString(dateKey+".top1", u.pref)
	u.top1.Bind(s1)
	u.top1.Validator = nil

	u.top2.Unbind()
	s2 := binding.BindPreferenceString(dateKey+".top2", u.pref)
	u.top2.Bind(s2)
	u.top2.Validator = nil

	u.top3.Unbind()
	s3 := binding.BindPreferenceString(dateKey+".top3", u.pref)
	u.top3.Bind(s3)
	u.top3.Validator = nil
}
