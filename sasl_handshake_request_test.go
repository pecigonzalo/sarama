//go:build !functional

package sarama

import "testing"

var baseSaslRequest = []byte{
	0, 3, 'f', 'o', 'o', // Mechanism
}

func TestSaslHandshakeRequest(t *testing.T) {
	request := new(SaslHandshakeRequest)
	request.Mechanism = "foo"
	testRequest(t, "basic", request, baseSaslRequest)
}
