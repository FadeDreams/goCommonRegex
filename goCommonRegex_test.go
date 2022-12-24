package goCommonRegex

import "testing"

func TestIsPositiveInt(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"123", true},
		{"0", true},
		{"-123", false},
		{"123.456", false},
		{"abc", false},
	}
	for _, test := range tests {
		if got := IsPositiveInt(test.input); got != test.want {
			t.Errorf("IsPositiveInt(%q) = %v, want %v", test.input, got, test.want)
		}
	}
}

func TestIsInt(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"123", true},
		{"0", true},
		{"-123", true},
		{"123.456", false},
		{"abc", false},
	}
	for _, test := range tests {
		if got := IsInt(test.input); got != test.want {
			t.Errorf("IsInt(%q) = %v, want %v", test.input, got, test.want)
		}
	}
}

func TestIsDecimalNum(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"123", true},
		{"0", true},
		{"-123", false},
		{"123.456", true},
		{"123.", false},
		{"abc", false},
	}
	for _, test := range tests {
		if got := IsDecimalNum(test.input); got != test.want {
			t.Errorf("IsDecimalNum(%q) = %v, want %v", test.input, got, test.want)
		}
	}
}

func TestIsNum(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"123", true},
		{"0", true},
		{"-123", true},
		{"123.456", true},
		{"123.", false},
	}
	for _, test := range tests {
		if got := IsNum(test.input); got != test.want {
			t.Errorf("IsNum(%q) = %v, want %v", test.input, got, test.want)
		}
	}
}

func TestIsAlphaNumeric(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"abc123", true},
		{"ABC123", true},
		{"123ABC", true},
		{"abc", true},
		{"ABC", true},
		{"123", true},
		{"!@#", false},
		{"abc!@#", false},
		{"", true},
	}
	for _, test := range tests {
		if got := IsAlphaNumeric(test.input); got != test.want {
			t.Errorf("IsAlphaNumeric(%q) = %v, want %v", test.input, got, test.want)
		}
	}
}

func TestIsAlphaNumericWithSpace(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"abc 123", true},
		{"ABC 123", true},
		{"123 ABC", true},
		{"abc", true},
		{"ABC", true},
		{"123", true},
		{"!@#", false},
		{"abc!@#", false},
		{"", true},
	}
	for _, test := range tests {
		if got := IsAlphaNumericWithSpace(test.input); got != test.want {
			t.Errorf("IsAlphaNumericWithSpace(%q) = %v, want %v", test.input, got, test.want)
		}
	}
}

func TestIsEmail(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"user@example.com", true},
		{"user@subdomain.example.com", true},
		{"user.name@example.com", true},
		{"user_name@example.com", true},
		{"user@123.com", true},
		{"user@example.co.uk", true},
		{"user@example.info", true},
		{"user@example", false},
		{"@example.com", false},
		{"user@", false},
		{"user@.", false},
		{"user@.com", false},
		{"user@example.c", false},
		{"", false},
	}
	for _, test := range tests {
		if got := IsEmail(test.input); got != test.want {
			t.Errorf("IsEmail(%q) = %v, want %v", test.input, got, test.want)
		}
	}
}

func TestIsGoodPassword(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"P@ssw0rd!", true},
		{"password1", false},
		{"PASSWORD1", false},
		{"p@ssw0rd", false},
		{"p@ssw0rd!", false},
		{"", false},
	}
	for _, test := range tests {
		if got := IsGoodPassword(test.input); got != test.want {
			t.Errorf("IsGoodPassword(%q) = %v, want %v", test.input, got, test.want)
		}
	}
}

func TestIsUsername(t *testing.T) {
	tests := []struct {
		username string
		expected bool
	}{
		{"abc.123", false},
		{"abc123", false},
		{"abc._123", false},
		{"abc123._", false},
		{"abc123._._", false},
		{"123", false},
		{"abc123.", false},
		{"abc123_", false},
	}
	for _, test := range tests {
		result := IsUsername(test.username)
		if result != test.expected {
			t.Errorf("IsUsername(%q) returned %v, expected %v", test.username, result, test.expected)
		}
	}
}

func TestIsURL(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"https://www.google.com", true},
		{"http://www.google.com", true},
		{"www.google.com", true},
		{"ftp://example.com", true},
	}
	for _, test := range tests {
		if got := IsURL(test.input); got != test.want {
			t.Errorf("IsURL(%q) = %v, want %v", test.input, got, test.want)
		}
	}
}

func TestIsIPv4(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"192.168.1.1", true},
		{"0.0.0.0", true},
		{"255.255.255.255", true},
		{"192.168.1.256", false},
		{"192.168.1", false},
		{"192.168.1.1.1", false},
		{"", false},
	}
	for _, test := range tests {
		if got := IsIPv4(test.input); got != test.want {
			t.Errorf("IsIPv4(%q) = %v, want %v", test.input, got, test.want)
		}
	}
}

func TestIsIPv6(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"2001:0db8:85a3:0000:0000:8a2e:0370:7334", true},
		{"2001:db8:85a3:0:0:8a2e:370:7334", true},
		{"2001:db8:85a3::8a2e:370:7334", true},
		{"2001:0db8:85a3:0000:0000:8a2e:0370:73341", false},
		{"2001:0db8:85a3:0000:0000:8a2e:0370:733:4", false},
		{"02001:0db8:85a3:0000:0000:8a2e:0370:7334", false},
		{"2001:0db8:85a3:0000:0000:8a2e:0370:73347", false},
		{"", false},
	}
	for _, test := range tests {
		if got := IsIPv6(test.input); got != test.want {
			t.Errorf("IsIPv6(%q) = %v, want %v", test.input, got, test.want)
		}
	}
}

func TestIsDate_yyyy_mm_dd(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"2022-12-24", true},
		{"2022-02-29", true},
		{"2022-13-01", false},
		{"2022-01-32", false},
		{"2022-01-00", false},
		{"2022-00-01", false},
		{"0000-01-01", false},
		{"", false},
	}
	for _, test := range tests {
		if got := IsDate_yyyy_mm_dd(test.input); got != test.want {
			t.Errorf("IsDate_yyyy_mm_dd(%q) = %v, want %v", test.input, got, test.want)
		}
	}
}
