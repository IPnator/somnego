package somnego

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

type Somneo struct {
	Host string
}

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
	fmt.Printf("Data of type %T: %v\n", result, result)
	return result, nil
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

func (somneo *Somneo) GetWakeupLightStatus() (WakeUpLightStatus, error) {
	return getData[WakeUpLightStatus](somneo, WakeUpLightStatusEndpoint)
}

func (somneo *Somneo) GetSensorData() (SensorData, error) {
	return getData[SensorData](somneo, SensorDataEnpoint)
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
