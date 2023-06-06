package common

type ValidationError struct {
	FailedField string `json:"field"`
	Message     string `json:"message"`
}

func (ve *ValidationError) String() string {
	return ve.Message
}

type ResponseData struct {
	Data    map[string]interface{}
	Message string
	TrackId string
	IsError bool
}

type finalResponse struct {
	TrackId string                 `json:"track_id"`
	Error   bool                   `json:"error"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}
