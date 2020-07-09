package handlers

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	error500 = errors.New("shit happens")
)

func TestGetAllProjects(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		//setup

		//assertions
		assert.Equal(t, http.StatusOK, rec.Code)
	})
}
