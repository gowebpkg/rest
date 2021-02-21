package rest

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestResponse(t *testing.T) {
	rw := httptest.NewRecorder()
	ts := struct {
		String string `json:"hello"`
	}{String: "world"}

	AsJSON(rw, ts)
	result, err := ioutil.ReadAll(rw.Result().Body)
	require.NoError(t, err)
	require.Equal(t, "application/json", rw.Header().Get("Content-Type"))
	require.Equal(t, "{\"hello\":\"world\"}", string(result))
}
