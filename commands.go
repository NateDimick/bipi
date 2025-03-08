package main

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

// note about pacmd - it was replaced in ubuntu 24.04. (pipewire vs pulseaudio)

func BigPictureMode() error {
	commands := []string{
		"xrandr --output DisplayPort-1 --mode 1920x1080",
		"xrandr --output DisplayPort-1 --left-of DisplayPort-2",
		"xrandr --output DisplayPort-1 --primary",
		"wpctl set-default 53",
		//		"pacmd set-default-sink alsa_output.pci-0000_03_00.1.hdmi-stereo", // this needs to happen last, after the display is turned back on, because the sink name changes when the display is off and the index increments each time too.  alsa_output.pci-0000_03_00.1.hdmi-stereo-extra3
	}
	return errors.Join(RunCommands(commands), TurnOn())
}

func NormalMode() error {
	commands := []string{
		//		"pacmd set-default-sink alsa_output.pci-0000_00_1f.3.analog-stereo",
		"wpctl set-default 33",
		"xrandr --output DisplayPort-2 --primary",
		"xrandr --output DisplayPort-1 --off",
	}
	return errors.Join(RunCommands(commands), TurnOff())
}

func RunCommands(cmds []string) error {
	var allErr error
	for _, c := range cmds {
		fmt.Println(c)
		out, err := exec.Command("bash", "-c", c).Output()
		fmt.Println(string(out))
		errors.Join(allErr, err)
	}
	return allErr
}

func StartSteamBigPicture() {
	// first check if steam is running
	steamPID := getSteamPID()
	if steamPID != "" {
		err := exec.Command("kill", steamPID)
		if err != nil {
			// TODO
		}
		// wait for steam to be dead
		for steamPID != "ok" {
			steamPID = getSteamPID()
			time.Sleep(time.Millisecond * 10)
		}
	}
	// then, open steam in big picture mode
	RunCommands([]string{
		"steam steam://open/bigpicture",
	})
}

func getSteamPID() string {
	stdout, err := exec.Command("bash", "-c", "pidof steam || echo 'ok'").Output()
	if err != nil {
		// TODO
		return ""
	}
	return strings.TrimSpace(string(stdout))
}
