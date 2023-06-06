package response

type baseRes struct {
	Message string `json:"message" example:"declarative message"`
	TrackId string `json:"track_id" example:"408c49e3-ba4e-48bd-9ff9-d94614f87c30"`
	Error   bool   `json:"error" example:"false"`
}

type BaseError struct {
	baseRes
	Data  struct{} `json:"data"`
	Error bool     `json:"error" example:"true"`
}
