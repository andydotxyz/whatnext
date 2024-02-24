package main

import (
	"fmt"
	"time"

	"fyne.io/cloud"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type ui struct {
	app  fyne.App
	date time.Time
	tabs *container.AppTabs

	title      *widget.Label
	prev, next *widget.Button

	goal, top1, top2, top3        *widget.Entry
	rem1, rem2, rem3              *widget.Label
	doneGoal, done1, done2, done3 *widget.Check

	area1, area2, area3          *widget.Label
	aim1, aim2, aim3             *widget.Entry
	aimDone1, aimDone2, aimDone3 *widget.Check

	mile1, mile2, mile3                   *widget.Entry
	mileDetail1, mileDetail2, mileDetail3 *widget.Entry
}

func (u *ui) makeDayUI() fyne.CanvasObject {
	u.goal = widget.NewEntry()
	u.doneGoal = widget.NewCheck("", func(bool) {})
	u.rem1 = widget.NewLabel("")
	u.top1 = widget.NewEntry()
	u.done1 = widget.NewCheck("", func(bool) {})
	u.rem2 = widget.NewLabel("")
	u.top2 = widget.NewEntry()
	u.done2 = widget.NewCheck("", func(bool) {})
	u.rem3 = widget.NewLabel("")
	u.top3 = widget.NewEntry()
	u.done3 = widget.NewCheck("", func(bool) {})

	return container.NewVBox(
		container.NewBorder(widget.NewLabel("Highlight"), nil, u.doneGoal, nil, u.goal),
		widget.NewLabel("Targets"),
		container.NewBorder(nil, nil,
			container.NewVBox(widget.NewLabel("1:"), u.done1), nil,
			container.NewVBox(u.rem1, u.top1)),
		container.NewBorder(nil, nil,
			container.NewVBox(widget.NewLabel("2:"), u.done2), nil,
			container.NewVBox(u.rem2, u.top2)),
		container.NewBorder(nil, nil,
			container.NewVBox(widget.NewLabel("3:"), u.done3), nil,
			container.NewVBox(u.rem3, u.top3)))
}

func (u *ui) makeQuarterUI() fyne.CanvasObject {
	u.mile1 = widget.NewEntry()
	u.mileDetail1 = widget.NewMultiLineEntry()
	u.mile2 = widget.NewEntry()
	u.mileDetail2 = widget.NewMultiLineEntry()
	u.mile3 = widget.NewEntry()
	u.mileDetail3 = widget.NewMultiLineEntry()

	return container.NewGridWithRows(3,
		container.NewBorder(nil, nil, widget.NewLabel("1:"), nil,
			container.NewBorder(u.mile1, nil, nil, nil, u.mileDetail1)),
		container.NewBorder(nil, nil, widget.NewLabel("2:"), nil,
			container.NewBorder(u.mile2, nil, nil, nil, u.mileDetail2)),
		container.NewBorder(nil, nil, widget.NewLabel("3:"), nil,
			container.NewBorder(u.mile3, nil, nil, nil, u.mileDetail3)))
}

func (u *ui) makeWeekUI() fyne.CanvasObject {
	u.area1 = widget.NewLabel("")
	u.aim1 = widget.NewEntry()
	u.aimDone1 = widget.NewCheck("", func(bool) {})
	u.area2 = widget.NewLabel("")
	u.aim2 = widget.NewEntry()
	u.aimDone2 = widget.NewCheck("", func(bool) {})
	u.area3 = widget.NewLabel("")
	u.aim3 = widget.NewEntry()
	u.aimDone3 = widget.NewCheck("", func(bool) {})

	return container.NewVBox(
		widget.NewLabel("Aims"),
		container.NewBorder(nil, nil,
			container.NewVBox(widget.NewLabel("1:"), u.aimDone1), nil,
			container.NewVBox(u.area1, u.aim1)),
		container.NewBorder(nil, nil,
			container.NewVBox(widget.NewLabel("2:"), u.aimDone2), nil,
			container.NewVBox(u.area2, u.aim2)),
		container.NewBorder(nil, nil,
			container.NewVBox(widget.NewLabel("3:"), u.aimDone3), nil,
			container.NewVBox(u.area3, u.aim3)))
}

func (u *ui) refreshDate() (w, q, y int) {
	y, w = u.date.Add(time.Hour * 48).ISOWeek() // start the week 2 days earlier
	q = ((w - 1) / 13) + 1

	switch u.tabs.SelectedIndex() {
	case 0:
		u.title.SetText(fmt.Sprintf("Quarter %d, %d", q, y))

		u.prev.OnTapped = func() {
			u.setDate(u.date.Add(time.Hour * -24 * 7 * 13))
		}
		u.next.OnTapped = func() {
			u.setDate(u.date.Add(time.Hour * 24 * 7 * 13))
		}
	case 1:
		u.title.SetText(fmt.Sprintf("Week %d, Q%d %d", ((w-1)%13)+1, q, y))

		u.prev.OnTapped = func() {
			u.setDate(u.date.Add(time.Hour * -24 * 7))
		}
		u.next.OnTapped = func() {
			u.setDate(u.date.Add(time.Hour * 24 * 7))
		}
	case 2:
		u.title.SetText(u.date.Format("Mon, 02 Jan 2006"))

		u.prev.OnTapped = func() {
			u.setDate(u.date.Add(time.Hour * -24))
		}
		u.next.OnTapped = func() {
			u.setDate(u.date.Add(time.Hour * 24))
		}
	}

	return
}

func (u *ui) setDate(t time.Time) {
	pref := u.app.Preferences()
	u.date = t
	w, q, y := u.refreshDate()
	dateKey := t.Format("20060102")
	weekKey := fmt.Sprintf("%dw%02d", y, w)
	quarterKey := fmt.Sprintf("%dq%d", y, q)

	bindPrefString(u.mile1, quarterKey+".mile1", pref)
	bindPrefString(u.mileDetail1, quarterKey+".info1", pref)
	bindPrefString(u.mile2, quarterKey+".mile2", pref)
	bindPrefString(u.mileDetail2, quarterKey+".info2", pref)
	bindPrefString(u.mile3, quarterKey+".mile3", pref)
	bindPrefString(u.mileDetail3, quarterKey+".info3", pref)

	bindPrefString(u.area1, quarterKey+".mile1", pref)
	bindPrefString(u.aim1, weekKey+".aim1", pref)
	bindPrefBool(u.aimDone1, weekKey+".aim1.done", pref)
	bindPrefString(u.area2, quarterKey+".mile2", pref)
	bindPrefString(u.aim2, weekKey+".aim2", pref)
	bindPrefBool(u.aimDone2, weekKey+".aim2.done", pref)
	bindPrefString(u.area3, quarterKey+".mile3", pref)
	bindPrefString(u.aim3, weekKey+".aim3", pref)
	bindPrefBool(u.aimDone3, weekKey+".aim3.done", pref)

	bindPrefString(u.goal, dateKey+".goal", pref)
	bindPrefBool(u.doneGoal, dateKey+".goal.done", pref)
	bindPrefString(u.rem1, weekKey+".aim1", pref)
	bindPrefString(u.top1, dateKey+".top1", pref)
	bindPrefBool(u.done1, dateKey+".top1.done", pref)
	bindPrefString(u.rem2, weekKey+".aim2", pref)
	bindPrefString(u.top2, dateKey+".top2", pref)
	bindPrefBool(u.done2, dateKey+".top2.done", pref)
	bindPrefString(u.rem3, weekKey+".aim3", pref)
	bindPrefString(u.top3, dateKey+".top3", pref)
	bindPrefBool(u.done3, dateKey+".top3.done", pref)

}

func (u *ui) makeUI() fyne.CanvasObject {
	u.date = time.Now()
	u.title = widget.NewLabel("Date placeholder")
	u.prev = widget.NewButtonWithIcon("", theme.NavigateBackIcon(), nil)
	u.next = widget.NewButtonWithIcon("", theme.NavigateNextIcon(), nil)
	date := container.NewBorder(nil, nil, u.prev, u.next, u.title)

	u.tabs = container.NewAppTabs(
		container.NewTabItemWithIcon("Quarter", theme.NewThemedResource(resourceQuarterSvg), u.makeQuarterUI()),
		container.NewTabItemWithIcon("Week", theme.NewThemedResource(resourceWeekSvg), u.makeWeekUI()),
		container.NewTabItemWithIcon("Day", theme.NewThemedResource(resourceDaySvg), u.makeDayUI()))
	if fyne.CurrentDevice().IsMobile() {
		u.tabs.SetTabLocation(container.TabLocationTrailing)
	} else {
		u.tabs.SetTabLocation(container.TabLocationLeading)
	}
	u.tabs.OnSelected = func(_ *container.TabItem) {
		_, _, _ = u.refreshDate()
	}
	u.tabs.SelectIndex(2)

	now := widget.NewButtonWithIcon("", theme.NewThemedResource(resourceTodaySvg), func() {
		u.setDate(time.Now())
	})
	sync := widget.NewButtonWithIcon("", theme.ViewRefreshIcon(), func() {
		cloud.ShowSettings(u.app, u.app.Driver().AllWindows()[0])
	})
	tools := container.NewBorder(nil, nil,
		now, sync, container.NewCenter(date))
	return container.NewBorder(tools, nil, nil, nil, u.tabs)
}
