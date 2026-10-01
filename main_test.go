package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestSetupRelationServesLoadedTable(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := setupRelation(map[string]string{"gu": "choki", "choki": "pa", "pa": "gu"})

	for hand, want := range map[string]string{"gu": "choki", "choki": "pa", "pa": "gu"} {
		w := httptest.NewRecorder()
		r.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/janken/"+hand, nil))
		if w.Code != http.StatusOK || w.Body.String() != want {
			t.Errorf("/janken/%s = %d %q, want 200 %q", hand, w.Code, w.Body.String(), want)
		}
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/janken/rock", nil))
	if w.Code != http.StatusNotFound {
		t.Errorf("/janken/rock = %d, want 404 for a hand not in the table", w.Code)
	}
}
