package messages

type Expires struct {
	AtHeight uint64 `json:"at_height,omitempty"`
	AtTime   string `json:"at_time,omitempty"`
	Never    Never  `json:"never,omitempty"`
}

type Never struct{}
