package main

import (
	"context"
	"fmt"
	kong "github.com/gamussa/testcontainers-go-kong"
	"github.com/gavv/httpexpect/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"testing"
)

func TestRateLimiting_CheckHeaders(t *testing.T) {

	if testing.Short() {
		t.Skip("skipping integration test")
	}

	ctx := context.Background()

	env := map[string]string{
		"KONG_DATABASE":               "off",
		"KONG_LOG_LEVEL":              "info",
		"KONG_NGINX_WORKER_PROCESSES": "1",
		"KONG_PROXY_ACCESS_LOG":       "/dev/stdout",
		"KONG_ADMIN_ACCESS_LOG":       "/dev/stdout",
		"KONG_PROXY_ERROR_LOG":        "/dev/stderr",
		"KONG_ADMIN_ERROR_LOG":        "/dev/stderr",
		"KONG_ADMIN_LISTEN":           "0.0.0.0:8001",
		"KONG_DECLARATIVE_CONFIG":     "/usr/local/kong/kong.yaml",
		//------------ WasmX -----------------
		"KONG_WASM_FILTERS_PATH": "/usr/local/kong/wasm/",
		"KONG_WASM":              "on",
		"KONG_NGINX_WASM_SHM_KONG_WASM_RATE_LIMITING_COUNTERS": "12m",
	}

	files := []testcontainers.ContainerFile{
		{
			HostFilePath:      "./test/config/demo.yml",
			ContainerFilePath: "/usr/local/kong/kong.yaml",
			FileMode:          0644, // see https://github.com/supabase/cli/pull/132/files
		},
		{
			HostFilePath:      "./rate-limiting.wasm", // copy the already compiled binary to the plugins dir
			ContainerFilePath: "/usr/local/kong/wasm/rate-limiting.wasm",
			FileMode:          0755,
		},
	}
	kong, err := kong.SetupKong(ctx,
		"kong/kong-gateway-dev:3.4.0.0-rc.1",
		env,
		files)
	require.NoError(t, err)

	consumer := TestLogConsumer{
		Msgs: []string{},
		Ack:  make(chan bool),
	}
	err = kong.StartLogProducer(ctx)
	assert.Nil(t, err)

	defer kong.StopLogProducer()
	kong.FollowOutput(&consumer)

	// Clean up the container after the test is complete
	t.Cleanup(func() {
		if err := kong.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate container: %s", err)
		}
	})

	e := httpexpect.Default(t, kong.ProxyURI)
	response := e.GET("/echo").Expect()

	// check if kong proxies request
	response.Status(200).Header("Via").Contains("kong")

	// check if RL headers are there
	response.Header("X-Ratelimit-Limit-Minute").
		IsEqual("3")
}

type TestLogConsumer struct {
	Msgs []string
	Ack  chan bool
}

const lastMessage = "DONE"

func (g *TestLogConsumer) Accept(l testcontainers.Log) {
	if string(l.Content) == fmt.Sprintf("echo %s\n", lastMessage) {
		g.Ack <- true
		return
	}

	logLine := string(l.Content)
	g.Msgs = append(g.Msgs, logLine)
	fmt.Print(logLine)
}
