package app

type Config struct {
	Queue Queue `atmc:"queue"`
}

type Queue struct {
	QueueWorkerCount uint `atmc:"queueWorkerCount"`
	QueueBufferSize  int  `atmc:"queueBufferSize"`
}
