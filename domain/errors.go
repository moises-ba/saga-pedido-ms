package domain

import "errors"

var ErrDuplicateKey = errors.New("registro ja existente")
var ErrRecordNofFound = errors.New("registro não encontrado")
