package initializers

import (
	"fmt"
	"goselflearn/internal/common"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type config struct {
	/**
	 * App
	 */
	Port *int `mapstructure:"PORT" validate:"omitempty,number"`

	/**
	 * Database
	 */
	DBHost    string `mapstructure:"DB_HOST" validate:"required,hostname|ip"`
	DBPort    string `mapstructure:"DB_PORT" validate:"required,numeric"`
	DBUser    string `mapstructure:"DB_USER" validate:"required"`
	DBPass    string `mapstructure:"DB_PASS" validate:"required"`
	DBName    string `mapstructure:"DB_NAME" validate:"required"`
	DBSSLMode bool   `mapstructure:"DB_SSL_MODE" validate:"omitempty,boolean"`

	/**
	 * JWT
	 */
	JWTExpiresInMin time.Duration `mapstructure:"JWT_EXPIRES_IN_MIN" validate:"required"`
	JWTSecret       string        `mapstructure:"JWT_SECRET" validate:"required,min=10"`
}

func (c *config) Validate() (errors []*common.ValidationError) {
	err := common.Validate.Struct(c)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var e common.ValidationError
			e.FailedField = err.Field()
			e.Message = common.GetValidatorErrorMessage(err.Tag(), err.Field(), err.Param())
			errors = append(errors, &e)
		}
	}
	return
}

var Config config

func LoadConfigurationFromDotEnv(path string) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("something went wrong in reading configuration from app.env > error: %s", err.Error())
		return
	}

	err = viper.Unmarshal(&Config)
	if err != nil {
		fmt.Printf("something went wrong in reading configuration from app.env > error: %s", err.Error())
		return
	}

	validationErrors := Config.Validate()
	if len(validationErrors) > 0 {
		fmt.Printf("[Error] configuration validation errors: %s\n", validationErrors)
		return
	}

	fmt.Println("configuration loaded and validated successfully from app.env")
}
