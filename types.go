package somnego

// Device details (XML representation)
type DeviceDetails struct {
	SpecVersion DeviceDetailsSpecVersion `xml:"specVersion"`
	Device      DeviceDetailsDevice      `xml:"device"`
}

type DeviceDetailsSpecVersion struct {
	Major int `xml:"major"`
	Minor int `xml:"minor"`
}

type DeviceDetailsDevice struct {
	DeviceType   string `xml:"deviceType"`
	FriendlyName string `xml:"friendlyName"`
	Manufacturer string `xml:"manufacturer"`
	ModelName    string `xml:"modelName"`
	ModelNumber  string `xml:"modelNumber"`
	UDN          string `xml:"UDN"`
	CppId        string `xml:"cppId"`
}

// Wakeup tones
type WakeupTones map[string]Tone

// Wind down dust
type WindDownDusk map[string]Tone

// Light themes
type LightThemes map[string]Tone

// General tone struct
type Tone struct {
	Name string `json:"name"`
}

type WakeUpStatus struct {
	WakeUpStatus int  `json:"wusts"`
	ReadyToPair  bool `json:"rpair"`
	PreviewMode  bool `json:"prvmd"`
	StoreDemo    bool `json:"sdemo"`
	PowerSave    bool `json:"pwrsz"`
	CurrentValue int  `json:"nrcur"`
	WizardStatus int  `json:"wizrd"`
	Brightness   int  `json:"brght"`
	DisplayOn    bool `json:"dspon"`
	WakeUpTime   int  `json:"wutim"`
	DurationTime int  `json:"dutim"`
	CanUpdate    bool `json:"canup"`
	SunsetTime   int  `json:"sntim"`
}

type LightStatus struct {
	LightLevel   int   `json:"ltlvl"`
	IsOn         bool  `json:"onoff"`
	Temperature  bool  `json:"tempy"`
	ControlType  int   `json:"ctype"`
	NightLight   bool  `json:"ngtlt"`
	WakeUpCurve  []int `json:"wucrv"`
	PWMOn        bool  `json:"pwmon"`
	PWMValues    []int `json:"pwmvs"`
	DimmerManual int   `json:"diman"`
}

type SensorData struct {
	// Current values
	MeasuredLux      float64 `json:"mslux"` // Brightness in Lux
	MeasuredTemp     float64 `json:"mstmp"` // Temperature in °C
	MeasuredHumidity float64 `json:"msrhu"` // Humidity in %
	MeasuredSound    int     `json:"mssnd"` // Sound level

	// Average values
	AverageLux      float64 `json:"avlux"` // Average brightness
	AverageTemp     float64 `json:"avtmp"` // Average temperature
	AverageHumidity float64 `json:"avhum"` // Average humidity
	AverageSound    int     `json:"avsnd"` // Average sound level

	EnergyScore int `json:"enscr"` // Energy score
}

type DuskSettings struct {
	IsOn         bool   `json:"onoff"`
	CurveType    int    `json:"curve"`
	Duration     int    `json:"durat"`
	ControlType  int    `json:"ctype"`
	SoundType    int    `json:"sndtp"`
	SoundDevice  string `json:"snddv"`
	SoundChannel string `json:"sndch"`
	SoundLevel   int    `json:"sndlv"`
	SoundStatus  int    `json:"sndss"`
}

type WakeupAlarmStatus struct {
	SnoozeTime       int              `json:"snztm"`
	AlarmEnvironment AlarmEnvironment `json:"aenvs"`
	ActiveAlarms     ActiveAlarms     `json:"aalms"`
	ProfileShare     int              `json:"prfsh"`
}

// Alarm environment settings
type AlarmEnvironment struct {
	ProfilesEnabled [16]bool `json:"prfen"`
	ProfileValues   [16]bool `json:"prfvs"`
	PowerSave       [48]int  `json:"pwrsv"`
}

// Active alarms
type ActiveAlarms struct {
	Year       [16]int `json:"ayear"` // Year of the alarm
	Month      [16]int `json:"amnth"` // Month (0-11)
	Day        [16]int `json:"alday"` // Day of the month
	DayBitMask [16]int `json:"daynm"` // Weekdays as bitmask (e.g., 254 = Mon-Sun)
	Hour       [16]int `json:"almhr"` // Hour (0-23)
	Minute     [16]int `json:"almmn"` // Minute (0-59)
}

type WakeupPlay struct {
	IsOn         bool   `json:"onoff"`
	Temperature  bool   `json:"tempy"`
	SoundVolume  int    `json:"sdvol"`
	SoundStatus  int    `json:"sndss"`
	SoundDevice  string `json:"snddv"`
	SoundChannel string `json:"sndch"`
}
