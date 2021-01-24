package test

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func httpResponse(test *testing.T, router *gin.Engine, request *http.Request, f func(w *httptest.ResponseRecorder) bool) {
	writer := httptest.NewRecorder()

	router.ServeHTTP(writer, request)

	if !f(writer) {
		test.Fail()
	}
}
