package network_test

import (
	"bytes"
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/chromedp/chromedp"
	"github.com/stretchr/testify/assert"

	"github.com/hawari17/page-profiler/network"
)

func TestNetworkMonitoring(t *testing.T) {
	responsePage, err := ioutil.ReadFile("sample_page.html")
	if err != nil {
		t.Fatal(err)
	}

	handlerFunc := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(responsePage)
	})

	server := httptest.NewServer(handlerFunc)
	defer server.Close()

	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	buf := new(bytes.Buffer)
	network.MonitorPageNetwork(ctx, buf, server.URL)

	assert.NotEmpty(t, buf, "Output should not be empty")
}
