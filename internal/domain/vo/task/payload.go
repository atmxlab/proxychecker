package task

type Payload struct {
	TargetURL *TargetURL `json:"target_url"`
}

type TargetURL struct {
	URL string `json:"url"`
}
