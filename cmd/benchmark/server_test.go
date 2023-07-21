package benchmark

import (
	"bytes"
	"flights/pkg/server"
	"fmt"
	"net/http"
	"testing"
	"time"
)

func init() {
	go serverStart()
	time.Sleep(time.Second)
}

func BenchmarkServer(b *testing.B) {

	payload := []byte(`[["IND","EWR"],["SFO","ATL"],["GSO","IND"],["ATL", "GSO"]]`)
	url := "http://localhost:8080/calculate"

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
		if err != nil {
			b.Logf("http.NewRequest().error: %v", err.Error())
			b.FailNow()
		}

		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			b.Logf("client.Do(req).error: %v", err.Error())
			b.FailNow()
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			b.Logf("status code error")
			b.FailNow()
		}
	}
}

func serverStart() {
	var err error

	mux := http.NewServeMux()

	calculateHandler := server.MiddlewarePost(http.HandlerFunc(server.GeneratesSubRoutesOfRoute))
	mux.Handle("/calculate", calculateHandler)

	fmt.Println("Server started at http://localhost:8080")
	if err = http.ListenAndServe(":8080", mux); err != nil {
		panic(fmt.Errorf("main.http.ListenAndServe().error: %v", err))
	}
}
