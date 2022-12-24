### goCommonRegex
A Go package for common regular expressions.

## Installation
To install goCommonRegex, use `go get`:
```go
go get -u github.com/fadedreams/goCommonRegex 
```

### Usage
Import the package in your Go code:
```go
import "github.com/fadedreams/goCommonRegex"
```


#### Functions

- IsPositiveInt(s string) bool: Returns true if s is a positive integer.
- IsInt(s string) bool: Returns true if s is an integer.
- IsDecimalNum(s string) bool: Returns true if s is a decimal number.
- IsNum(s string) bool: Returns true if s is a number (integer or decimal).
- IsAlphaNumeric(s string) bool: Returns true if s is alphanumeric (contains only letters and digits).
- IsAlphaNumericWithSpace(s string) bool: Returns true if s is alphanumeric with spaces (contains only letters, digits, and spaces).
- IsEmail(email string) bool: Returns true if email is a valid email address.
- IsGoodPassword(password string) bool: Returns true if password is a valid password.
- IsUsername(username string) bool: Returns true if username is a valid username.
- IsURL(url string) bool: Returns true if url is a valid URL with the http or https protocol.
- IsIPv4(IPv4 string) bool: Returns true if IPv4 is a valid IPv4 address.

#### Example

```go
package main

import (
	"fmt"
	"github.com/fadedreams/goCommonRegex"
)

func main() {
	fmt.Println(goCommonRegex.IsEmail("user1@example.com"))  // Output: true
	fmt.Println(goCommonRegex.IsEmail("user1.example.com"))  // Output: false
	fmt.Println(goCommonRegex.IsGoodPassword("P@ssw0rd"))      // Output: true
	fmt.Println(goCommonRegex.IsGoodPassword("passw0rd"))      // Output: false
	fmt.Println(goCommonRegex.IsUsername("user123"))       // Output: true
	fmt.Println(goCommonRegex.IsUsername("user@123"))      // Output: false
}
```
