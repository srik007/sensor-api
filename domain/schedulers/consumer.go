package scheduler

func (j *SchedulerJob) Listen(sensorId int) {
	for {
		event := <-j.Channles[sensorId]
		j.SensorRepository.SaveData(event)
	}
}
