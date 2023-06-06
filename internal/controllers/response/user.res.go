package response

import "goselflearn/internal/models"

type registerUserData struct {
	User models.User `json:"user"`
}
type RegisterUser struct {
	baseRes
	Data registerUserData `json:"data"`
}

type loginUserData struct {
	AccessToken string `json:"access_token"`
}
type LoginUser struct {
	baseRes
	Data loginUserData `json:"data"`
}

type getUserInfoWrapper struct {
	User models.User `json:"user"`
}
type GetUserInfo struct {
	baseRes
	Data getUserInfoWrapper `json:"data"`
}
