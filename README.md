# somnogo - control your Somneo from go

1. Install and import the somnego dependency:

```go
import "github.com/IPnator/somnego"
```

2. Create a Somneo instance:

```go
device := somnego.Somneo{Host: "<IP-or-local-domain>"}
```

3. Do whatever you want - load information, control light

```go
// Load device details
details, err := device.GetDeviceDetails()

// Turn main light on
err = device.TurnMainLightOn()

// Set main light level (in percent)
err = device.SetMainLightLevel(85)

// Get sensor data such as humidity
sensorData, err := device.GetSensorData()
fmt.Printf("Humidity: %v\n", sensorData.AverageHumidity)

// and more…
```
