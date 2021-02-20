package routes

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHome404(t *testing.T) {
	// t.Parallel()
	t.Name()
	resp, err := http.Get("http://localhost:4000/missing")
	require.NoError(t, err, "GET '/missing'")
	defer resp.Body.Close()
	require.Equal(t, "404 Not Found", resp.Status, "Home - 404 Not Found")

}

func TestHome200(t *testing.T) {
	// t.Parallel()
	t.Name()
	resp, err := http.Get("http://localhost:4000/")
	require.NoError(t, err, "GET '/'")

	defer resp.Body.Close()
	require.NoError(t, err, "Read body")
	require.Equal(t, "2001 OK", resp.Status, "Home - 200 OK")

}
