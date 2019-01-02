package main

import (
   "encoding/json"
   "net/http"
   "log"
   "net/http/httptest"
   "testing"
   "github.com/gin-gonic/gin"
   	"github.com/stretchr/testify/assert"
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
   req, _ := http.NewRequest(method, path, nil)
   req.Header.Set("X-Forwarded-For", "127.0.0.1")
   req.Header.Set("Host", "google.com")
   w := httptest.NewRecorder()
   r.ServeHTTP(w, req)
   return w
}

type result struct {
    IP []string `json:"IP"`
   	RemoteIP string  `json:"RemoteIP"`
}
func TestRequest(t *testing.T) {
   
   body := gin.H{
      "IP": "127.0.0.1",
      "RemoteIP" : "127.0.0.1",
   }
  
   router := routerEngine()
   // Perform a GET request with that handler.
   w := performRequest(router, "GET", "/ips")

   assert.Equal(t, http.StatusOK, w.Code)
   var response result
 //   var response []interface{}
   err := json.Unmarshal([]byte(w.Body.String()), &response)

  //  response["RemoteIP"]
   log.Print(w.Body.String())
   log.Print(err)
   log.Print(response)
   assert.Nil(t, err)
   assert.True(t, response.RemoteIP != "")
  // assert.True(t, remoteIpValExist)
   assert.Equal(t, body["IP"], response.IP[0])
   
}