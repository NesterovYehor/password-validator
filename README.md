# Password Validator

Password Validator is a Go-based package designed to evaluate the strength and validity of passwords. It includes various checks such as length, entropy, and other customizable criteria, ensuring the security of passwords for use in applications.

## Features

- **Entropy Calculation**: Calculates the entropy of a password to gauge its randomness and strength.
- **Length Validation**: Ensures that passwords meet the required minimum and maximum length criteria.
- **Modular Design**: Easily extensible to add more validation checks.
- **Unit Tests**: Includes a set of tests to verify the correct functionality of the validation rules.

## Requirements

- Go 1.x or higher

## Installation

1. Clone the repository:
   `git clone https://github.com/NesterovYehor/password-validator.git` and navigate to the project directory.

2. Install dependencies using Go modules: `go mod tidy`.

## Usage

The package provides functions to validate passwords based on customizable rules such as length and entropy. You can integrate it into your own projects to enforce strong password policies.

Example of validating a password:
```go
package main

import (
    "fmt"
    "password-validator"
)

func main() {
    valid, err := password_validator.Validate("examplePassword123!")
    if err != nil {
        fmt.Println("Error:", err)
    }
    if valid {
        fmt.Println("Password is valid")
    } else {
        fmt.Println("Password is invalid")
    }
}
