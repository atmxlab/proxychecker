package aggregate

import "github.com/atmxlab/proxychecker/internal/domain/entity"

type Task struct {
	task    *entity.Task
	proxy   *entity.Proxy
	checker *entity.Checker
}
