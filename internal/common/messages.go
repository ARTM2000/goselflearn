package common

import "fmt"

const (
	MessageUserCreated                  = "user successfully registered"
	MessageUserWithThisEmailExists      = "user with this email exists"
	MessageUserEmailOrPasswordIsInvalid = "email or password is invalid"
	MessageInternalServerError          = "there is a problem on our side"
	MessageSuccessfulLogin              = "welcome"
	MessageUnauthorized                 = "unauthorized"
	MessageNewPostCreated               = "new post created"
)

func GetValidatorErrorMessage(tag string, field string, param string) string {
	switch tag {
	case "required":
		return fmt.Sprintf("%s is required", field)
	case "omitempty":
		return ""
	case "len":
		return fmt.Sprintf("%s must be exactly %s characters long", field, param)
	case "min":
		return fmt.Sprintf("%s must be at least %s characters long", field, param)
	case "max":
		return fmt.Sprintf("%s must be at most %s characters long", field, param)
	case "eq":
		return fmt.Sprintf("%s must be equal to %s", field, param)
	case "ne":
		return fmt.Sprintf("%s must not be equal to %s", field, param)
	case "oneof":
		return fmt.Sprintf("%s must be one of %s", field, param)
	case "lt":
		return fmt.Sprintf("%s must be less than %s", field, param)
	case "lte":
		return fmt.Sprintf("%s must be less than or equal to %s", field, param)
	case "gt":
		return fmt.Sprintf("%s must be greater than %s", field, param)
	case "gte":
		return fmt.Sprintf("%s must be greater than or equal to %s", field, param)
	case "eqfield":
		return fmt.Sprintf("%s must be equal to %s", field, param)
	case "nefield":
		return fmt.Sprintf("%s must not be equal to %s", field, param)
	case "gtfield":
		return fmt.Sprintf("%s must be greater than %s", field, param)
	case "gtefield":
		return fmt.Sprintf("%s must be greater than or equal to %s", field, param)
	case "ltfield":
		return fmt.Sprintf("%s must be less than %s", field, param)
	case "ltefield":
		return fmt.Sprintf("%s must be less than or equal to %s", field, param)
	case "alpha":
		return fmt.Sprintf("%s must contain only letters", field)
	case "alphanum":
		return fmt.Sprintf("%s must contain only letters and numbers", field)
	case "alphaunicode":
		return fmt.Sprintf("%s must contain only Unicode letters", field)
	case "alphanumunicode":
		return fmt.Sprintf("%s must contain only Unicode letters and numbers", field)
	case "numeric":
		return fmt.Sprintf("%s must be a number", field)
	case "number":
		return fmt.Sprintf("%s must be a number", field)
	case "hexadecimal":
		return fmt.Sprintf("%s must be a hexadecimal number", field)
	case "hexcolor":
		return fmt.Sprintf("%s must be a valid hexadecimal color code", field)
	case "rgb":
		return fmt.Sprintf("%s must be a valid RGB color code", field)
	case "rgba":
		return fmt.Sprintf("%s must be a valid RGBA color code", field)
	case "hsl":
		return fmt.Sprintf("%s must be a valid HSL color code", field)
	case "hsla":
		return fmt.Sprintf("%s must be a valid HSLA color code", field)
	case "email":
		return fmt.Sprintf("%s must be a valid email address", field)
	case "url":
		return fmt.Sprintf("%s must be a valid URL", field)
	case "uri":
		return fmt.Sprintf("%s must be a valid URI", field)
	case "base64":
		return fmt.Sprintf("%s must be a valid base64-encoded string", field)
	case "contains":
		return fmt.Sprintf("%s must contain %s", field, param)
	case "containsany":
		return fmt.Sprintf("%s must contain at least one of the following characters: %s", field, param)
	case "excludes":
		return fmt.Sprintf("%s may not contain %s", field, param)
	case "excludesall":
		return fmt.Sprintf("%s may not contain any of the following characters: %s", field, param)
	case "excludesrune":
		return fmt.Sprintf("%s may not contain the following character: %s", field, param)
	case "isbn":
		return fmt.Sprintf("%s must be a valid ISBN number", field)
	case "isbn10":
		return fmt.Sprintf("%s must be a valid ISBN-10 number", field)
	case "isbn13":
		return fmt.Sprintf("%s must be a valid ISBN-13 number", field)
	case "uuid":
		return fmt.Sprintf("%s must be a valid UUID", field)
	case "uuid3":
		return fmt.Sprintf("%s must be avalid UUIDv3", field)
	case "uuid4":
		return fmt.Sprintf("%s must be a valid UUIDv4", field)
	case "uuid5":
		return fmt.Sprintf("%s must be a valid UUIDv5", field)
	case "ascii":
		return fmt.Sprintf("%s must contain only ASCII characters", field)
	case "printascii":
		return fmt.Sprintf("%s must contain only printable ASCII characters", field)
	case "multibyte":
		return fmt.Sprintf("%s must contain one or more multibyte characters", field)
	case "datauri":
		return fmt.Sprintf("%s must be a valid data URI", field)
	case "latitude":
		return fmt.Sprintf("%s must be a valid latitude value", field)
	case "longitude":
		return fmt.Sprintf("%s must be a valid longitude value", field)
	case "ssn":
		return fmt.Sprintf("%s must be a valid Social Security number", field)
	case "ipv4":
		return fmt.Sprintf("%s must be a valid IPv4 address", field)
	case "ipv6":
		return fmt.Sprintf("%s must be a valid IPv6 address", field)
	case "ip":
		return fmt.Sprintf("%s must be a valid IP address", field)
	case "cidr":
		return fmt.Sprintf("%s must be a valid CIDR notation", field)
	case "cidrv4":
		return fmt.Sprintf("%s must be a valid IPv4 CIDR notation", field)
	case "cidrv6":
		return fmt.Sprintf("%s must be a valid IPv6 CIDR notation", field)
	case "tcp4_addr":
		return fmt.Sprintf("%s must be a valid TCPv4 address", field)
	case "tcp6_addr":
		return fmt.Sprintf("%s must be a valid TCPv6 address", field)
	case "tcp_addr":
		return fmt.Sprintf("%s must be a valid TCP address", field)
	case "udp4_addr":
		return fmt.Sprintf("%s must be a valid UDPv4 address", field)
	case "udp6_addr":
		return fmt.Sprintf("%s must be a valid UDPv6 address", field)
	case "udp_addr":
		return fmt.Sprintf("%s must be a valid UDP address", field)
	case "unix_addr":
		return fmt.Sprintf("%s must be a valid Unix domain socket address", field)
	case "boolean":
		return fmt.Sprintf("%s must ba a valid Boolean", field)

	// custom tags
	case "password":
		return fmt.Sprintf("%s is not strong enough", field)

	default:
		fmt.Printf("not defined tag for validation in messages: %s\n", tag)
		return fmt.Sprintf("Validation error for %s", field)
	}
}
