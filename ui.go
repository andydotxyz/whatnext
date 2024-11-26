package main

import (
	"fmt"
	"time"

	"fyne.io/cloud"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type ui struct {
	app  fyne.App
	date time.Time
	tabs *container.AppTabs

	title      *widget.Label
	prev, next *widget.Button

	goal, top1, top2              *widget.Entry
	top3                          *tapScrollLineEntry
	rem1, rem2, rem3              *widget.Label
	doneGoal, done1, done2, done3 *widget.Check

	area1, area2, area3          *widget.Label
	aim1, aim2                   *widget.Entry
	aim3                         *tapScrollLineEntry
	aimDone1, aimDone2, aimDone3 *widget.Check

	mile1, mile2, mile3      *widget.Entry
	mileDetail1, mileDetail2 *widget.Entry
	mileDetail3              *tapScrollLineEntry

	habit1, habit2, habit3                                        *widget.Entry
	hDone11, hDone12, hDone13, hDone14, hDone15, hDone16, hDone17 *widget.Check
	hDone21, hDone22, hDone23, hDone24, hDone25, hDone26, hDone27 *widget.Check
	hDone31, hDone32, hDone33, hDone34, hDone35, hDone36, hDone37 *widget.Check
}

func (u *ui) makeDayUI(s *container.Scroll) fyne.CanvasObject {
	u.goal = widget.NewEntry()
	u.doneGoal = widget.NewCheck("", func(bool) {})
	u.rem1 = widget.NewLabel("")
	u.top1 = widget.NewEntry()
	u.done1 = widget.NewCheck("", func(bool) {})
	u.rem2 = widget.NewLabel("")
	u.top2 = widget.NewEntry()
	u.done2 = widget.NewCheck("", func(bool) {})
	u.rem3 = widget.NewLabel("")
	u.top3 = &tapScrollLineEntry{}
	u.top3.ExtendBaseWidget(u.top3)
	u.top3.parent = s
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

func (u *ui) makeQuarterUI(s *container.Scroll) fyne.CanvasObject {
	u.mile1 = widget.NewEntry()
	u.mileDetail1 = widget.NewMultiLineEntry()
	u.mile2 = widget.NewEntry()
	u.mileDetail2 = widget.NewMultiLineEntry()
	u.mile3 = widget.NewEntry()
	u.mileDetail3 = &tapScrollLineEntry{}
	u.mileDetail3.ExtendBaseWidget(u.mileDetail3)
	u.mileDetail3.MultiLine = true
	u.mileDetail3.parent = s

	return container.NewGridWithRows(3,
		container.NewBorder(nil, nil, widget.NewLabel("1:"), nil,
			container.NewBorder(u.mile1, nil, nil, nil, u.mileDetail1)),
		container.NewBorder(nil, nil, widget.NewLabel("2:"), nil,
			container.NewBorder(u.mile2, nil, nil, nil, u.mileDetail2)),
		container.NewBorder(nil, nil, widget.NewLabel("3:"), nil,
			container.NewBorder(u.mile3, nil, nil, nil, u.mileDetail3)))
}

func (u *ui) makeWeekUI(s *container.Scroll) fyne.CanvasObject {
	u.area1 = widget.NewLabel("")
	u.aim1 = widget.NewEntry()
	u.aimDone1 = widget.NewCheck("", func(bool) {})
	u.area2 = widget.NewLabel("")
	u.aim2 = widget.NewEntry()
	u.aimDone2 = widget.NewCheck("", func(bool) {})
	u.area3 = widget.NewLabel("")
	u.aim3 = &tapScrollLineEntry{}
	u.aim3.ExtendBaseWidget(u.aim3)
	u.aim3.parent = s
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
	case 3:
		u.title.SetText(fmt.Sprintf("Week %d habits (Q%d %d)", ((w-1)%13)+1, q, y))

		u.prev.OnTapped = func() {
			u.setDate(u.date.Add(time.Hour * -24 * 7))
		}
		u.next.OnTapped = func() {
			u.setDate(u.date.Add(time.Hour * 24 * 7))
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

	bindPrefString(u.habit1, quarterKey+".habit1", pref)
	bindPrefBool(u.hDone11, weekKey+".habit1.done1", pref)
	bindPrefBool(u.hDone12, weekKey+".habit1.done2", pref)
	bindPrefBool(u.hDone13, weekKey+".habit1.done3", pref)
	bindPrefBool(u.hDone14, weekKey+".habit1.done4", pref)
	bindPrefBool(u.hDone15, weekKey+".habit1.done5", pref)
	bindPrefBool(u.hDone16, weekKey+".habit1.done6", pref)
	bindPrefBool(u.hDone17, weekKey+".habit1.done7", pref)
	bindPrefString(u.habit2, quarterKey+".habit2", pref)
	bindPrefBool(u.hDone21, weekKey+".habit2.done1", pref)
	bindPrefBool(u.hDone22, weekKey+".habit2.done2", pref)
	bindPrefBool(u.hDone23, weekKey+".habit2.done3", pref)
	bindPrefBool(u.hDone24, weekKey+".habit2.done4", pref)
	bindPrefBool(u.hDone25, weekKey+".habit2.done5", pref)
	bindPrefBool(u.hDone26, weekKey+".habit2.done6", pref)
	bindPrefBool(u.hDone27, weekKey+".habit2.done7", pref)
	bindPrefString(u.habit3, quarterKey+".habit3", pref)
	bindPrefBool(u.hDone31, weekKey+".habit3.done1", pref)
	bindPrefBool(u.hDone32, weekKey+".habit3.done2", pref)
	bindPrefBool(u.hDone33, weekKey+".habit3.done3", pref)
	bindPrefBool(u.hDone34, weekKey+".habit3.done4", pref)
	bindPrefBool(u.hDone35, weekKey+".habit3.done5", pref)
	bindPrefBool(u.hDone36, weekKey+".habit3.done6", pref)
	bindPrefBool(u.hDone37, weekKey+".habit3.done7", pref)
}

func (u *ui) makeHabitUI() fyne.CanvasObject {
	u.habit1 = widget.NewEntry()
	u.habit2 = widget.NewEntry()
	u.habit3 = widget.NewEntry()
	u.hDone11 = widget.NewCheck("", func(bool) {})
	u.hDone12 = widget.NewCheck("", func(bool) {})
	u.hDone13 = widget.NewCheck("", func(bool) {})
	u.hDone14 = widget.NewCheck("", func(bool) {})
	u.hDone15 = widget.NewCheck("", func(bool) {})
	u.hDone16 = widget.NewCheck("", func(bool) {})
	u.hDone17 = widget.NewCheck("", func(bool) {})
	u.hDone21 = widget.NewCheck("", func(bool) {})
	u.hDone22 = widget.NewCheck("", func(bool) {})
	u.hDone23 = widget.NewCheck("", func(bool) {})
	u.hDone24 = widget.NewCheck("", func(bool) {})
	u.hDone25 = widget.NewCheck("", func(bool) {})
	u.hDone26 = widget.NewCheck("", func(bool) {})
	u.hDone27 = widget.NewCheck("", func(bool) {})
	u.hDone31 = widget.NewCheck("", func(bool) {})
	u.hDone32 = widget.NewCheck("", func(bool) {})
	u.hDone33 = widget.NewCheck("", func(bool) {})
	u.hDone34 = widget.NewCheck("", func(bool) {})
	u.hDone35 = widget.NewCheck("", func(bool) {})
	u.hDone36 = widget.NewCheck("", func(bool) {})
	u.hDone37 = widget.NewCheck("", func(bool) {})

	return container.NewVBox(
		container.NewHBox(layout.NewSpacer(),
			container.NewGridWithColumns(7,
				widget.NewLabelWithStyle(" S ", fyne.TextAlignCenter, fyne.TextStyle{}),
				widget.NewLabelWithStyle(" S ", fyne.TextAlignCenter, fyne.TextStyle{}),
				widget.NewLabelWithStyle(" M ", fyne.TextAlignCenter, fyne.TextStyle{}),
				widget.NewLabelWithStyle(" T ", fyne.TextAlignCenter, fyne.TextStyle{}),
				widget.NewLabelWithStyle(" W ", fyne.TextAlignCenter, fyne.TextStyle{}),
				widget.NewLabelWithStyle(" T ", fyne.TextAlignCenter, fyne.TextStyle{}),
				widget.NewLabelWithStyle(" F ", fyne.TextAlignCenter, fyne.TextStyle{}))),
		container.NewBorder(nil, nil, widget.NewLabel("1:"), nil,
			u.habit1),
		container.NewHBox(layout.NewSpacer(),
			container.NewGridWithColumns(7,
				u.hDone11, u.hDone12, u.hDone13, u.hDone14, u.hDone15, u.hDone16, u.hDone17)),
		container.NewBorder(nil, nil, widget.NewLabel("2:"), nil,
			u.habit2),
		container.NewHBox(layout.NewSpacer(),
			container.NewGridWithColumns(7,
				u.hDone21, u.hDone22, u.hDone23, u.hDone24, u.hDone25, u.hDone26, u.hDone27)),
		container.NewBorder(nil, nil, widget.NewLabel("3:"), nil,
			u.habit3),
		container.NewHBox(layout.NewSpacer(),
			container.NewGridWithColumns(7,
				u.hDone31, u.hDone32, u.hDone33, u.hDone34, u.hDone35, u.hDone36, u.hDone37)),
	)
}

func (u *ui) makeUI() fyne.CanvasObject {
	u.date = time.Now()
	u.title = widget.NewLabel("Date placeholder")
	u.prev = widget.NewButtonWithIcon("", theme.NavigateBackIcon(), nil)
	u.next = widget.NewButtonWithIcon("", theme.NavigateNextIcon(), nil)
	date := container.NewBorder(nil, nil, u.prev, u.next, u.title)

	qScroll := container.NewVScroll(nil)
	qScroll.Content = u.makeQuarterUI(qScroll)
	wScroll := container.NewVScroll(nil)
	wScroll.Content = u.makeWeekUI(wScroll)
	dScroll := container.NewVScroll(nil)
	dScroll.Content = u.makeDayUI(dScroll)
	u.tabs = container.NewAppTabs(
		container.NewTabItemWithIcon("Quarter", theme.NewThemedResource(resourceQuarterSvg),
			qScroll),
		container.NewTabItemWithIcon("Week", theme.NewThemedResource(resourceWeekSvg),
			wScroll),
		container.NewTabItemWithIcon("Day", theme.NewThemedResource(resourceDaySvg),
			dScroll),
		container.NewTabItemWithIcon("Habits", theme.NewThemedResource(resourceHabitSvg),
			container.NewVScroll(u.makeHabitUI())))
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

type tapScrollLineEntry struct {
	widget.Entry

	parent *container.Scroll
}

func (t *tapScrollLineEntry) Tapped(ev *fyne.PointEvent) {
	t.Entry.Tapped(ev)

	// TODO the VScroll listen for size
	// if resized to smaller than min, check focused object
	// if it's position is > 150 (for example) then scroll down

	// force scrolling to the bottom in a trivial sort of way
	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(25 * time.Millisecond)
			t.parent.ScrollToBottom()
		}
	}()
}
