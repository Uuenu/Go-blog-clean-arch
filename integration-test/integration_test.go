package integrationtest

import (
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	. "github.com/Eun/go-hit"
)

const (
	// Attempts connection
	host       = "localhost:3000"
	healthPath = host + "/healthz"
	attempts   = 20

	// HTTP REST
	basePath = host + "/v1"

	// RabbitMQ RPC
	// rmqURL            = "amqp://guest:guest@rabbitmq:5672/"
	// rpcServerExchange = "rpc_server"
	// rpcClientExchange = "rpc_client"
	// requests          = 10
)

func TestMain(m *testing.M) {
	err := healthCheck(attempts)
	if err != nil {
		log.Fatalf("Integration tests: host %s is not available: %s", host, err)
	}

	log.Printf("Integration tests: host %s is available", host)

	code := m.Run()
	os.Exit(code)
}

func healthCheck(attempts int) error {
	var err error

	for attempts > 0 {
		err = Do(Get(healthPath), Expect().Status().Equal(http.StatusOK))
		if err == nil {
			return nil
		}

		log.Printf("Integration tests: url %s is not available, attempts left: %d", healthPath, attempts)

		time.Sleep(time.Second)

		attempts--
	}

	return err
}

// HTTP POST: /translation/do-translate.
func TestHTTPDoSignup(t *testing.T) {
	body := `{
		"username": "name",
		"email": "mame@gmail.com",
		"password": "qwerty"
	}`
	Test(t,
		Description("DoSignup Success"),
		Post(basePath+"/author/signup"),
		Send().Headers("Content-Type").Add("application/json"),
		Send().Body().String(body),
		Expect().Status().Equal(http.StatusOK),
		Expect().Body().JSON().JQ(".translation").Equal("text for translation"),
	)

	body = `{
		"username": "rony",
		"email": "rony99@gmail.com",
		"password": "qwerty"
	}`
	Test(t,
		Description("DoSignup Fail"),
		Post(basePath+"/author/signup"),
		Send().Headers("Content-Type").Add("application/json"),
		Send().Body().String(body),
		Expect().Status().Equal(http.StatusBadRequest),
		Expect().Body().JSON().JQ(".error").Equal("invalid request body"),
	)
}
