package response

type UnauthorizedError struct {
	BaseError
	Message string `json:"message" example:"unauthorized"`
}

type singleValidationErr struct {
	Field   string `json:"field" example:"password"`
	Message string `json:"message" example:"Password is not strong enough"`
}
type validationErrorWrapper struct {
	Error singleValidationErr `json:"error"`
}
type ValidationError struct {
	BaseError
	Data validationErrorWrapper `json:"data"`
}
