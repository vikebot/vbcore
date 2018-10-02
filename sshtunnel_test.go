package vbcore

import "testing"

func TestNewEndpoint(t *testing.T) {
	endpoint := NewEndpoint("vikebot.com", 2400)
	if endpoint.Host != "vikebot.com" || endpoint.Port != 2400 {
		t.Fail()
	}
}

func TestNewEndpointAddr(t *testing.T) {
	endpoint := NewEndpointAddr("vikebot.com:2400")
	if endpoint.Host != "vikebot.com" || endpoint.Port != 2400 {
		t.Fail()
	}

	endpoint = NewEndpointAddr("")
	if endpoint != nil {
		t.Fail()
	}

	endpoint = NewEndpointAddr("vikebot.com")
	if endpoint != nil {
		t.Fail()
	}

	endpoint = NewEndpointAddr("vikebot.com:2400:2401")
	if endpoint != nil {
		t.Fail()
	}

	endpoint = NewEndpointAddr("vikebot.com:someillegalport")
	if endpoint != nil {
		t.Fail()
	}
}

func TestNewEndpointString(t *testing.T) {
	endpoint := NewEndpoint("vikebot.com", 2400)
	if endpoint.String() != "vikebot.com:2400" {
		t.Fail()
	}
}
