package errors

import "errors"

var ConversationAlreadyExists = errors.New("conversation already exists")
var ConversationNotFound = errors.New("conversation not found")
