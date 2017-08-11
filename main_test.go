package main

import (
  "bytes"
  "github.com/gin-gonic/gin"
  "net/http"
  "net/http/httptest"
  "testing"
)

func TestGetInstruction(t *testing.T) {
  gin.SetMode(gin.TestMode)
  testRouter := SetupRouter()

  req, err := http.NewRequest("GET", "/api/v1/instructions/1", nil)
  if err != nil {
    t.Errorf("Get hearteat failed with error %d.", err)
  }

  resp := httptest.NewRecorder()
  testRouter.ServeHTTP(resp, req)

  if resp.Code != 200 {
    t.Errorf("/api/v1/instructions failed with error code %d.", resp.Code)
  }
}

func TestGetInstructions(t *testing.T) {
  gin.SetMode(gin.TestMode)
  testRouter := SetupRouter()

  req, err := http.NewRequest("GET", "/api/v1/instructions", nil)
  if err != nil {
    t.Errorf("Get hearteat failed with error %d.", err)
  }

  resp := httptest.NewRecorder()
  testRouter.ServeHTTP(resp, req)

  if resp.Code != 200 {
    t.Errorf("/api/v1/instructions failed with error code %d.", resp.Code)
  }
}

func TestPostInstruction(t *testing.T) {
  gin.SetMode(gin.TestMode)
  testRouter := SetupRouter()

  body := bytes.NewBuffer([]byte("{\"event_status\": \"83\", \"event_name\": \"100\"}"))

  req, err := http.NewRequest("POST", "/api/v1/instructions", body)
  req.Header.Set("Content-Type", "application/json")
  if err != nil {
    t.Errorf("Post hearteat failed with error %d.", err)
  }

  resp := httptest.NewRecorder()
  testRouter.ServeHTTP(resp, req)

  if resp.Code != 201 {
    t.Errorf("/api/v1/instructions failed with error code %d.", resp.Code)
  }
}

func TestPutInstruction(t *testing.T) {
  gin.SetMode(gin.TestMode)
  testRouter := SetupRouter()

  body := bytes.NewBuffer([]byte("{\"event_status\": \"83\", \"event_name\": \"100\"}"))

  req, err := http.NewRequest("PUT", "/api/v1/instructions/1", body)
  req.Header.Set("Content-Type", "application/json")
  if err != nil {
    t.Errorf("Put hearteat failed with error %d.", err)
  }

  resp := httptest.NewRecorder()
  testRouter.ServeHTTP(resp, req)

  if resp.Code != 200 {
    t.Errorf("/api/v1/instructions failed with error code %d.", resp.Code)
  }
}

