package handlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/likejehu/dearapril/handlers/mocks"
	"github.com/stretchr/testify/assert"
)

var (
	error500 = errors.New("shit happens")
)

func TestGetAllProjects(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		//setup
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()

		mockApp := mocks.AppController{}
		handler := &Handler{&mockApp}
		handler.GetAllProjects(rec, req)
		//assertions
		assert.Equal(t, http.StatusOK, rec.Code)
	})
}
