package goCommonRegex

import (
	"regexp"
	"strconv"
)

func IsPositiveInt(s string) bool {
	re := regexp.MustCompile(`^\d+$`)
	return re.MatchString(s)
}

func IsInt(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func IsDecimalNum(s string) bool {
	re := regexp.MustCompile(`^\d*(\.\d+)?$`)
	return re.MatchString(s)
}

func IsNum(s string) bool {
	re := regexp.MustCompile(`^-?\d*(\.\d+)?$`)
	return re.MatchString(s)
}

func IsAlphaNumeric(s string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9]*$`)
	return re.MatchString(s)
}

func IsAlphaNumericWithSpace(s string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9 ]*$`)
	return re.MatchString(s)
}

func IsEmail(email string) bool {
	// Compile the regular expression
	re := regexp.MustCompile(`^([a-zA-Z0-9._-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,6})$`)

	// Check if the email matches the regular expression
	if re.MatchString(email) {
		return true
	}
	return false
}

// At least 1 lowercase letter
// At least 1 uppercase letter
// At least 1 number
// At least 1 special character
// At least 8 characters long
func IsGoodPassword(password string) bool {

	return len(password) >= 8 &&
		regexp.MustCompile(`[A-Z]`).MatchString(password) && regexp.MustCompile(`[a-z]`).MatchString(password) &&
		regexp.MustCompile(`[0-9]`).MatchString(password) && (regexp.MustCompile(`[^\w\s]`).MatchString(password))

}

func IsUsername(username string) bool {
	// Use a regular expression to check if the string is alphanumeric and may include _ and -
	// and has a length of 3 to 16 characters
	matched, err := regexp.MatchString("^(?![-_]*$)[A-Za-z0-9][A-Za-z0-9-_]{3,20}[A-Za-z0-9]$", username)
	if err != nil {
		// There was an error executing the regular expression, so return false
		return false
	}
	return matched
}

func IsURL(url string) bool {
	// Use a regular expression to check if the string is a URL with the http or https protocol
	matched, err := regexp.MatchString(`(([\w]+:)?//)?(([\d\w]|%[a-fA-F\d]{2,2})+(:([\d\w]|%[a-fA-f\d]{2,2})+)?@)?([\d\w][-\d\w]{0,253}[\d\w]\.)+[\w]{2,4}(:[\d]+)?(/([-+_~.\d\w]|%[a-fA-f\d]{2,2})*)*(\?(&?([-+_~.\d\w]|%[a-fA-f\d]{2,2})=?)*)?(#([-+_~.\d\w]|%[a-fA-f\d]{2,2})*)?`, url)
	if err != nil {
		// There was an error executing the regular expression, so return false
		return false
	}
	return matched
}

func IsIPv4(IPv4 string) bool {
	// Use a regular expression to check if the string is a URL with the http or https protocol
	matched, err := regexp.MatchString(`^((25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])$`, IPv4)
	if err != nil {
		// There was an error executing the regular expression, so return false
		return false
	}
	return matched
}

func IsIPv6(IPv6 string) bool {
	// Use a regular expression to check if the string is a URL with the http or https protocol
	matched, err := regexp.MatchString(`^([0-9A-Fa-f]{0,4}:){2,7}([0-9A-Fa-f]{1,4}$|((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.|$)){4})$`, IPv6)
	if err != nil {
		// There was an error executing the regular expression, so return false
		return false
	}
	return matched
}

func IsDate_yyyy_mm_dd(date string) bool {
	// Use a regular expression to check if the string is a URL with the http or https protocol
	matched, err := regexp.MatchString(`^([12]\d{3}-(0[1-9]|1[0-2])-(0[1-9]|[12]\d|3[01]))$`, date)
	if err != nil {
		// There was an error executing the regular expression, so return false
		return false
	}
	return matched
}

func IsDate_dd_mm_yyyy(date string) bool {
	// Use a regular expression to check if the string is a URL with the http or https protocol
	matched, err := regexp.MatchString(`^((0[1-9]|[12]\d|3[01])-(0[1-9]|1[0-2])-[12]\d{3})$`, date)
	if err != nil {
		// There was an error executing the regular expression, so return false
		return false
	}
	return matched
}

func IsTime_hh_mm_12h(time string) bool {
	// Use a regular expression to check if the string is a URL with the http or https protocol
	matched, err := regexp.MatchString(`((1[0-2]|0?[1-9]):([0-5][0-9])( ?([AaPp][Mm]))?)`, time)
	if err != nil {
		// There was an error executing the regular expression, so return false
		return false
	}
	return matched
}

func IsTime_hh_mm_24h(time string) bool {
	// Use a regular expression to check if the string is a URL with the http or https protocol
	matched, err := regexp.MatchString(`^([0-9]|0[0-9]|1[0-9]|2[0-3]):[0-5][0-9]$`, time)
	if err != nil {
		// There was an error executing the regular expression, so return false
		return false
	}
	return matched
}

func IsHtmlTag(html string) bool {
	// Use a regular expression to check if the string is a URL with the http or https protocol
	matched, err := regexp.MatchString(`<\/?[a-z][^>]*>`, html)
	if err != nil {
		// There was an error executing the regular expression, so return false
		return false
	}
	return matched
}

func IsSlug(slug string) bool {
	// Use a regular expression to check if the string is a URL with the http or https protocol
	matched, err := regexp.MatchString(`^[a-z0-9]+(?:-[a-z0-9]+)*$`, slug)
	if err != nil {
		// There was an error executing the regular expression, so return false
		return false
	}
	return matched
}

func IsTel(tel string) bool {
	// Use a regular expression to check if the string is a URL with the http or https protocol
	matched, err := regexp.MatchString(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`, tel)
	if err != nil {
		// There was an error executing the regular expression, so return false
		return false
	}
	return matched
}
