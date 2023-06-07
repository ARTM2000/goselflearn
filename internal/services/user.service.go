package services

import (
	"fmt"
	"strconv"
	"time"

	"github.com/ARTM2000/goselflearn/internal/common"
	"github.com/ARTM2000/goselflearn/internal/controllers/dto"
	"github.com/ARTM2000/goselflearn/internal/initializers"
	"github.com/ARTM2000/goselflearn/internal/models"
	"github.com/ARTM2000/goselflearn/internal/repositories"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func NewUserService() userService {
	userRepo := repositories.NewUserRepository()
	return userService{
		userRepository: userRepo,
	}
}

type userService struct {
	userRepository repositories.UserRepository
}

func (us *userService) RegisterUser(user *dto.UserRegister) (*models.User, error) {
	existingUser, err := us.userRepository.FindByEmail(user.Email)
	if err != nil {
		return nil, fmt.Errorf(common.MessageInternalServerError)
	}

	if existingUser != nil {
		return nil, fmt.Errorf(common.MessageUserWithThisEmailExists)
	}

	passwordBytes := []byte(user.Password)
	hashPassword, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("hashing password error:", err.Error())
		return nil, fmt.Errorf(common.MessageInternalServerError)
	}

	newUser := models.User{
		Name:         user.Name,
		Email:        user.Email,
		HashPassword: string(hashPassword),
	}

	if err := us.userRepository.CreateUser(&newUser); err != nil {
		return nil, fmt.Errorf(common.MessageInternalServerError)
	}
	return &newUser, nil
}

func (us *userService) LoginUser(user *dto.UserLogin) (accessToken *string, err error) {
	existingUser, err := us.userRepository.FindByEmail(user.Email)
	if err != nil {
		return nil, fmt.Errorf(common.MessageInternalServerError)
	}

	if existingUser == nil {
		return nil, fmt.Errorf(common.MessageUserEmailOrPasswordIsInvalid)
	}

	err = bcrypt.CompareHashAndPassword([]byte(existingUser.HashPassword), []byte(user.Password))
	if err != nil {
		return nil, fmt.Errorf(common.MessageUserEmailOrPasswordIsInvalid)
	}

	now := time.Now().UTC()
	claims := &jwt.MapClaims{
		"exp": now.Add(initializers.Config.JWTExpiresInMin).Unix(),
		"iat": now.Unix(),
		"nbf": now.Unix(),
		"ext": map[string]string{
			"id":    fmt.Sprint(existingUser.ID),
			"email": existingUser.Email,
		},
	}

	accessTokenByte := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := accessTokenByte.SignedString([]byte(initializers.Config.JWTSecret))

	if err != nil {
		fmt.Println("error in creating token", err)
		return nil, fmt.Errorf(common.MessageInternalServerError)
	}

	return &tokenString, nil
}

func (us *userService) VerifyUserAccessToken(token string) (user *models.User, err error) {
	tokenByte, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("2>>>>>>>>>>>>>>>>", t.Header["alg"])
			return nil, fmt.Errorf("unexpected signing method: %s", t.Header["alg"])
		}
		return []byte(initializers.Config.JWTSecret), nil
	})

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	claims, ok := tokenByte.Claims.(jwt.MapClaims)
	if !ok || !tokenByte.Valid {
		return nil, fmt.Errorf("claims not reachable or token is not valid")
	}

	extraData, _ := claims["ext"].(map[string]interface{})
	userId, _ := extraData["id"].(string)
	userIdUInt, _ := strconv.ParseUint(userId, 10, 32)
	user, err = us.userRepository.FindUserById(userIdUInt)

	if err != nil {
		return nil, fmt.Errorf("no user for this token exists")
	}

	return user, nil
}
