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

func TestDuskGetLightThemes(t *testing.T) {
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
	device.GetSensorData()
}
