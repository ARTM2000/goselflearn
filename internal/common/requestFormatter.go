package common

func FormatResponse(resData ResponseData) finalResponse {
	message := ""
	if resData.Message != "" {
		message = resData.Message
	}

	data := map[string]interface{}{}
	if resData.Data != nil {
		data = resData.Data
	}

	if resData.TrackId == "" {
		panic("track id should be defined in format response function")
	}

	return finalResponse{
		TrackId: resData.TrackId,
		Error:   resData.IsError,
		Message: message,
		Data:    data,
	}
}
