package main

import (
	"fmt"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func RunUI(window *app.Window) error {
	theme := material.NewTheme()
	var ops op.Ops
	var systemSwitch widget.Bool
	var autoLaunch widget.Bool
	autoLaunch.Value = true
	var steamButton widget.Clickable
	var resetButton widget.Clickable
	for {
		switch e := window.Event().(type) {
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)
			if resetButton.Clicked(gtx) {
				if !systemSwitch.Value {
					go NormalMode()
				}
			}
			if steamButton.Clicked(gtx) {
				go StartSteamBigPicture()
			}
			if systemSwitch.Update(gtx) {
				go func() {
					fmt.Println("switch toggle", systemSwitch.Value)
					var toggleErr error
					if systemSwitch.Value {
						toggleErr = BigPictureMode()
						if autoLaunch.Value {
							go StartSteamBigPicture()
						}
					} else {
						toggleErr = NormalMode()
					}
					if toggleErr != nil {
						fmt.Println(toggleErr.Error())
					}
				}()
			}
			ss := material.Switch(theme, &systemSwitch, "Switch System Settings")
			al := material.CheckBox(theme, &autoLaunch, "Launch Steam with Big Picture Mode Toggle")
			rb := material.Button(theme, &resetButton, "Normal Mode Manual Reset")
			sb := material.Button(theme, &steamButton, "Start Steam Big Picture")
			layout.Flex{
				Axis:      layout.Vertical,
				Spacing:   layout.SpaceEvenly,
				Alignment: layout.Middle,
			}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{
						Axis:      layout.Horizontal,
						Spacing:   layout.SpaceEvenly,
						Alignment: layout.Middle,
					}.Layout(gtx,
						layout.Rigid(material.Body2(theme, "Normal Mode").Layout),
						layout.Rigid(ss.Layout),
						layout.Rigid(material.Body2(theme, "Big Picture Mode").Layout),
					)
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Center.Layout(gtx, al.Layout)
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Inset{Left: unit.Dp(10), Right: unit.Dp(10)}.Layout(gtx, rb.Layout)
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Inset{Left: unit.Dp(10), Right: unit.Dp(10)}.Layout(gtx, sb.Layout)
				}),
			)
			e.Frame(gtx.Ops)
		case app.DestroyEvent:
			return e.Err
		}
	}
}
