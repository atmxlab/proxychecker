package queue

type result struct {
	err  error
	task Task
}

func newResult(err error, task Task) result {
	return result{err: err, task: task}
}
