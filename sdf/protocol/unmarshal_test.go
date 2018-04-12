package protocol_test

import (
	"net/http"
	"strings"
	"testing"

	"github.com/polefishu/sdk-builder/sdf/protocol"
	"github.com/polefishu/sdk-builder/sdf/request"
	"github.com/stretchr/testify/assert"
)

type mockCloser struct {
	*strings.Reader
	Closed bool
}

func (m *mockCloser) Close() error {
	m.Closed = true
	return nil
}

func TestUnmarshalDrainBody(t *testing.T) {
	b := &mockCloser{Reader: strings.NewReader("example body")}
	r := &request.Request{HTTPResponse: &http.Response{
		Body: b,
	}}

	protocol.UnmarshalDiscardBody(r)
	assert.NoError(t, r.Error)
	assert.Equal(t, 0, b.Len())
	assert.True(t, b.Closed)
}

func TestUnmarshalDrainBodyNoBody(t *testing.T) {
	r := &request.Request{HTTPResponse: &http.Response{}}

	protocol.UnmarshalDiscardBody(r)
	assert.NoError(t, r.Error)
}
