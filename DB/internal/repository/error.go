package repository

import "errors"

var ErrNotFound = errors.New("not found")
var ErrListEmpty = errors.New("list Empty")
var ErrHttpIssue = errors.New("error on http request")
