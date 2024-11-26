package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

type boolWid interface {
	Unbind()
	Bind(binding.Bool)
}

func bindPrefBool(w boolWid, key string, p fyne.Preferences) {
	w.Unbind()
	g := binding.BindPreferenceBool(key, p)
	w.Bind(g)
}

type stringWid interface {
	Unbind()
	Bind(binding.String)
}

func bindPrefString(w stringWid, key string, p fyne.Preferences) {
	w.Unbind()
	g := binding.BindPreferenceString(key, p)
	w.Bind(g)

	if e, ok := w.(*widget.Entry); ok {
		e.Validator = nil
	}
	if e, ok := w.(*tapScrollLineEntry); ok {
		e.Validator = nil
	}
}
