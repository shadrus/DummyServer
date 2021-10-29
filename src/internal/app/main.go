package app

import (
	"bytes"
	"fmt"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"strings"
)

func createKeyValuePairs(m map[string][]string) string {
	b := new(bytes.Buffer)
	for key, value := range m {
		fmt.Fprintf(b, "%s=%s ", key, value)
	}
	return b.String()
}

func StartServer() {
	router := httprouter.New()
	router.GET("/*path", requestFunc)
	router.POST("/*path", requestFunc)
	router.PUT("/*path", requestFunc)
	router.DELETE("/*path", requestFunc)
	router.HEAD("/*path", requestFunc)
	router.PATCH("/*path", requestFunc)
	router.OPTIONS("/*path", requestFunc)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func requestFunc(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	buf := new(strings.Builder)
	io.Copy(buf, request.Body)
	log.
		WithFields(log.Fields{"params": params.ByName("path")}).
		WithFields(log.Fields{"Headers": createKeyValuePairs(request.Header)}).
		WithFields(log.Fields{"body": buf}).
		WithFields(log.Fields{"method": request.Method}).
		Info("Got request")
	writer.WriteHeader(200)
}
