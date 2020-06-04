package db

import (
	"github.com/pkg/errors"
)

// Error404 is 404 err for store, when key is not found
var Error404 = errors.New("key not found")
