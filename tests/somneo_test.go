package tests

import (
	"fmt"
	"testing"

	"github.com/IPnator/somnego"
)

var device = somnego.Somneo{Host: "somneo.local"}

func TestGetDeviceDetails(t *testing.T) {
	_, err := device.GetDeviceDetails()
	if err != nil {
		t.Error("Loading device details failed: ", err)
	}
}

func TestGetWakeupTones(t *testing.T) {
	data, err := device.GetWakeupTones()
	if err != nil {
		t.Error("Loading wakeup tones failed: ", err)
	}

	expected := [8]string{"Forest Birds", "Summer Birds", "Buddha Wakeup", "Morning Alps", "Yoga Harmony", "Nepal Bowls", "Summer Lake", "Ocean Waves"}

	for i := range 8 {
		if data[fmt.Sprintf("%d", i+1)].Name != expected[i] {
			t.Errorf("Wakeup tone %v is empty", i+1)
		}
	}
}

func TestGetWindDownDusk(t *testing.T) {
	data, err := device.GetWindDownDusk()
	if err != nil {
		t.Error("Loading wind down dusk failed: ", err)
	}

	expected := [4]string{"Soft Rain", "Ocean Waves", "Under Water", "Summer Lake"}

	for i := range 4 {
		if data[fmt.Sprintf("%d", i+1)].Name != expected[i] {
			t.Errorf("Wind down dusk %v is empty", i+1)
		}
	}
}

func TestGetLightThemes(t *testing.T) {
	data, err := device.GetLightThemes()
	if err != nil {
		t.Error("Loading light themes failed: ", err)
	}

	expected := [4]string{"Sunny day", "Island red", "Nordic white", "Caribbean red"}

	for i := range 4 {
		if data[fmt.Sprintf("%d", i+1)].Name != expected[i] {
			t.Errorf("Light theme %v is empty", i+1)
		}
	}
}

func TestGetDuskLightThemes(t *testing.T) {
	data, err := device.GetDuskLightThemes()
	if err != nil {
		t.Error("Loading light themes failed: ", err)
	}

	expected := [4]string{"Sunny day", "Island red", "Nordic white", "Caribbean red"}

	for i := range 4 {
		if data[fmt.Sprintf("%d", i+1)].Name != expected[i] {
			t.Errorf("Light theme %v is empty", i+1)
		}
	}
}

func TestGetSensorData(t *testing.T) {
	_, err := device.GetSensorData()
	if err != nil {
		t.Error("Loading sensor data failed: ", err)
	}
}

func TestTurnMainLightOn(t *testing.T) {
	err := device.TurnMainLightOn()
	if err != nil {
		t.Error("Turning main light on failed: ", err)
	}
}

func TestTurnMainLightOff(t *testing.T) {
	err := device.TurnMainLightOff()
	if err != nil {
		t.Error("Turning main light on failed: ", err)
	}
}

func TestSetMainLightLevel(t *testing.T) {
	err := device.SetMainLightLevel(100)
	err = device.SetMainLightLevel(0)
	err = device.SetMainLightLevel(100)
	err = device.SetMainLightLevel(1078)
	err = device.SetMainLightLevel(-10)

	if err != nil {
		t.Error("Setting main light failed: ", err)
	}
}

func TestTurnNightLightOn(t *testing.T) {
	err := device.TurnNightlightOn()
	if err != nil {
		t.Error("Turning night light on failed: ", err)
	}
}

func TestTurnNightLightOff(t *testing.T) {
	err := device.TurnNightlightOff()
	if err != nil {
		t.Error("Turning night light off failed: ", err)
	}
}

func TestTurnOffAllLights(t *testing.T) {
	err := device.TurnOffAllLights()
	if err != nil {
		t.Error("Turning all lights off failed: ", err)
	}
}
