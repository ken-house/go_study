package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type myHandler struct {
}

func (h *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Hello string `json:"hello"`
	}{
		Hello: "hello world",
	}
	dataByte, _ := json.Marshal(data)
	w.Write(dataByte)
}

func main() {
	pool := x509.NewCertPool()
	caCertPath := "/Users/zonst/Documents/work/go_study/http/server/cert/ca.crt"

	caCrt, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return
	}
	pool.AppendCertsFromPEM(caCrt)

	s := &http.Server{
		Addr:    ":10443",
		Handler: &myHandler{},
		TLSConfig: &tls.Config{
			ClientAuth: tls.RequireAndVerifyClientCert,
			ClientCAs:  pool,
		},
	}

	err = s.ListenAndServeTLS("/Users/zonst/Documents/work/go_study/http/server/cert/server.crt", "/Users/zonst/Documents/work/go_study/http/server/cert/server.key")
	if err != nil {
		panic(err)
	}
}
