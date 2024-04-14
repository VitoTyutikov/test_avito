package tests

import (
	"avito_test_task/internal/db"
	"avito_test_task/internal/routers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserBanner(t *testing.T) {
	if err := db.InitDatabase(); err != nil {
		t.Fatalf("init/connect to database error: %v", err)
	}
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	routers.InitRoutes(router)

	ts := httptest.NewServer(router)
	defer ts.Close()

	testCases := []struct {
		description    string
		token          string
		featureId      string
		tagId          string
		expectedStatus int
	}{
		{"User_token", "user_token", "11", "10", http.StatusNotFound},
		{"Admin_token", "admin_token", "11", "10", http.StatusOK},
		{"Invalid token", "invalid_token", "1", "2", http.StatusForbidden},
		{"No token", "", "1", "2", http.StatusUnauthorized},
		{"Not fount", "user_token", "10000", "20000", http.StatusNotFound},
	}
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			req, _ := http.NewRequest("GET", ts.URL+"/user_banner?feature_id="+tc.featureId+"&tag_id="+tc.tagId, nil)
			req.Header.Set("token", tc.token)

			res, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Fatalf("Failed to make request: %v", err)
			}
			assert.Equal(t, tc.expectedStatus, res.StatusCode)
			res.Body.Close()
		})
	}
}
