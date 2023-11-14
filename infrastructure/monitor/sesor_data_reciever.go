package monitor

func (j *SensorDataMonitorJob) RecieveSensorData(sensorId int) {
	for {
		event := <-j.Channles[sensorId]
		j.SensorApp.SaveData(event)
	}
}
