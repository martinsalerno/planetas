package controllers

import (
  "testing"
  "net/http"
  "net/http/httptest"
  "github.com/stretchr/testify/assert"
  _ "github.com/jinzhu/gorm/dialects/sqlite"
)

func TestWeatherHandler(t *testing.T) {
  r, _ := http.NewRequest("GET", "/clima?dia=180", nil)
  w := httptest.NewRecorder()

  WeatherHandler(w, r)
  assert.Equal(t, http.StatusOK, w.Code)

  assert.Equal(t, "{\"dia\":180,\"clima\":\"sequia\"}\n", string(w.Body.Bytes()))
}