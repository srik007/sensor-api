package scheduler

func (j *SchedulerJob) RecieveSensorData(sensorId int) {
	for {
		event := <-j.Channles[sensorId]
		j.SensorRepository.SaveData(event)
	}
}
