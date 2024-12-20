package somnego

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

type Somneo struct {
	Host        string
	LightStatus *LightStatus
	SensorData  *SensorData
}

// getData fetches data from the specified endpoint, unmarshals it into the specified type, and returns it.
// It uses a generic type T to allow for different return types based on the endpoint.
func getData[T any](somneo *Somneo, endpoint string) (T, error) {
	var result T
	data, err := GET(fmt.Sprintf("https://%v%v", somneo.Host, endpoint))
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(data, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func setData[T any](somneo *Somneo, endpoint string, payload T) error {
	_, err := PUT[T](fmt.Sprintf("https://%v%v", somneo.Host, endpoint), payload)
	return err
}

func (somneo *Somneo) getLightStatusIfNil() {
	if somneo.LightStatus == nil {
		data, _ := somneo.GetLightStatus()
		somneo.LightStatus = &data
	}
}

func (somneo *Somneo) GetDeviceDetails() (DeviceDetails, error) {
	res, err := GET(fmt.Sprintf("https://%v%v", somneo.Host, DeviceDetailsEndpoint))
	var deviceDetails DeviceDetails
	if err != nil {
		return deviceDetails, err
	}
	err = xml.Unmarshal(res, &deviceDetails)
	if err != nil {
		return deviceDetails, err
	}
	return deviceDetails, nil
}

func (somneo *Somneo) GetWakeupTones() (WakeupTones, error) {
	return getData[WakeupTones](somneo, WakeupTonesEndpoint)
}

func (somneo *Somneo) GetWindDownDusk() (WindDownDusk, error) {
	return getData[WindDownDusk](somneo, WindDownDuskEndpoint)
}

func (somneo *Somneo) GetLightThemes() (LightThemes, error) {
	return getData[LightThemes](somneo, LightThemesEndpoint)
}

func (somneo *Somneo) GetDuskLightThemes() (LightThemes, error) {
	return getData[LightThemes](somneo, DuskLightThemesEndpoint)
}

func (somneo *Somneo) GetWakeupStatus() (WakeUpStatus, error) {
	return getData[WakeUpStatus](somneo, WakeUpStatusEndpoint)
}

func (somneo *Somneo) GetLightStatus() (LightStatus, error) {
	return getData[LightStatus](somneo, LightEndpoint)
}

func (somneo *Somneo) GetSensorData() (SensorData, error) {
	data, err := getData[SensorData](somneo, SensorDataEnpoint)
	if err != nil {
		return data, err
	}
	somneo.SensorData = &data
	return data, nil
}

func (somneo *Somneo) GetDuskSettings() (DuskSettings, error) {
	return getData[DuskSettings](somneo, DuskSettingsEndpoint)
}

func (somneo *Somneo) GetWakeupAlarmStatus() (WakeupAlarmStatus, error) {
	return getData[WakeupAlarmStatus](somneo, WakeUpAlarmStatusEndpoint)
}

func (somneo *Somneo) GetAlarmEnvironment() (AlarmEnvironment, error) {
	return getData[AlarmEnvironment](somneo, AlarmEnvironmentEndpoint)
}

func (somneo *Somneo) GetActiveAlarms() (ActiveAlarms, error) {
	return getData[ActiveAlarms](somneo, ActiveAlarmsEndpoint)
}

func (somneo *Somneo) GetWakeupPlay() (WakeupPlay, error) {
	return getData[WakeupPlay](somneo, WakeupPlayEndpoint)
}

// Control main light
func (somneo *Somneo) TurnMainLightOn() error {
	somneo.getLightStatusIfNil()
	somneo.LightStatus.IsOn = true
	return setData[LightStatus](somneo, LightEndpoint, *somneo.LightStatus)
}

func (somneo *Somneo) TurnMainLightOff() error {
	somneo.getLightStatusIfNil()
	somneo.LightStatus.IsOn = false
	return setData[LightStatus](somneo, LightEndpoint, *somneo.LightStatus)
}

func (somneo *Somneo) SetMainLightLevel(percent float64) error {
	somneo.getLightStatusIfNil()

	// Light level is a value between 0 and 25
	lightLevel := percent * 0.25
	if percent < 0 {
		lightLevel = 0
	}
	if percent > 100 {
		lightLevel = 25
	}

	var payload LightStatus = *somneo.LightStatus
	payload.IsOn = true
	if percent == 0 {
		payload.IsOn = false
	}
	payload.LightLevel = int(lightLevel)
	return setData[LightStatus](somneo, LightEndpoint, payload)
}

// Control night light
func (somneo *Somneo) TurnNightlightOn() error {
	somneo.getLightStatusIfNil()
	if somneo.LightStatus.IsOn == true {
		somneo.LightStatus.IsOn = false
	}
	somneo.LightStatus.NightLight = true
	return setData[LightStatus](somneo, LightEndpoint, *somneo.LightStatus)
}

func (somneo *Somneo) TurnNightlightOff() error {
	somneo.getLightStatusIfNil()
	if somneo.LightStatus.IsOn {
		somneo.LightStatus.IsOn = false
	}
	somneo.LightStatus.NightLight = false
	return setData[LightStatus](somneo, LightEndpoint, *somneo.LightStatus)
}

func (somneo *Somneo) TurnOffAllLights() error {
	somneo.getLightStatusIfNil()
	somneo.LightStatus.IsOn = false
	somneo.LightStatus.NightLight = false
	return setData[LightStatus](somneo, LightEndpoint, *somneo.LightStatus)
}
