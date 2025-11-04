package app

type Config struct {
	queue Queue
}

type Queue struct {
	QueueWorkerCount int16
	QueueBufferSize  int16
}
