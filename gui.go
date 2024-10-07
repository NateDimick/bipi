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
	var steamButton widget.Clickable
	for {
		switch e := window.Event().(type) {
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)
			if steamButton.Clicked(gtx) {
				StartSteamBigPicture()
			}
			if systemSwitch.Update(gtx) {
				fmt.Println("switch toggle", systemSwitch.Value)
				var toggleErr error
				if systemSwitch.Value {
					toggleErr = BigPictureMode()
				} else {
					toggleErr = NormalMode()
				}
				if toggleErr != nil {
					fmt.Println(toggleErr.Error())
				}
			}
			ss := material.Switch(theme, &systemSwitch, "Switch System Settings")
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
					return layout.Inset{Left: unit.Dp(10), Right: unit.Dp(10)}.Layout(gtx, sb.Layout)
				}),
			)
			e.Frame(gtx.Ops)
		case app.DestroyEvent:
			return e.Err
		}
	}
}
