package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/koron/go-ssdp"
)

func FindRoku() string {
	services, err := ssdp.Search("roku:ecp", 5, ":3080")
	if err != nil {
		panic(fmt.Errorf("SSDP error: %w", err))
	}

	if len(services) == 0 {
		return ""
	}

	return services[0].Location
}

func TurnOn() error {
	rokuIP := FindRoku()
	if rokuIP == "" {
		return errors.New("no roku tv found")
	}
	// RIP - apparently roku tv is not always on? no way to send power on command
	// CEC (requires adapter)
	// IR? would also require adapter to blast IR signal
	req, _ := http.NewRequest(http.MethodPost, rokuIP+"keypress/InputHDMI2", nil)
	_, err := http.DefaultClient.Do(req)
	return err
}

func TurnOff() error {
	rokuIP := FindRoku()
	if rokuIP == "" {
		return errors.New("no roku tv found")
	}
	req, _ := http.NewRequest(http.MethodPost, rokuIP+"keypress/PowerOff", nil)
	_, err := http.DefaultClient.Do(req)
	return err
}
