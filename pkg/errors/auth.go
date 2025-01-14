package errors

import "errors"

var InvalidPassword = errors.New("invalid password")
var UserAlreadyExists = errors.New("user already exists")
