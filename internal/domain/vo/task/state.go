package task

type State struct {
	result Result
}

func (s State) Result() Result {
	return s.result
}

func (s State) SetResult(res Result) State {
	s.result = res
	return s
}
