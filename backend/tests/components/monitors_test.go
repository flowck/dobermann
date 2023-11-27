package components_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/flowck/dobermann/backend/tests/client"
)

func TestMonitors(t *testing.T) {
	user := createAccount(t)
	token := login(t, user.Email, user.Password)
	cli := getClient(token)

	resp01, err := cli.CreateMonitor(ctx, client.CreateMonitorRequest{
		EndpointUrl: "http://localhost:8090",
	})
	require.NoError(t, err)

	assert.Equal(t, http.StatusCreated, resp01.StatusCode)
}
