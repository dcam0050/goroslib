package apislave

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/aler9/goroslib/pkg/xmlrpc"
)

func TestClient(t *testing.T) {
	s, err := xmlrpc.NewServer("localhost:9999")
	require.NoError(t, err)
	defer s.Close()

	go s.Serve(func(raw *xmlrpc.RequestRaw) interface{} {
		switch raw.Method {
		case "getPid":
			return ResponseGetPid{Code: 1}

		case "shutdown":
			return ResponseShutdown{Code: 1}

		case "requestTopic":
			return ResponseRequestTopic{Code: 1}

		case "getBusInfo":
			return ResponseGetBusInfo{Code: 1}
		}
		return xmlrpc.ErrorRes{}
	})

	c := NewClient("localhost:9999", "test")

	func() {
		res, err := c.GetPid()
		require.NoError(t, err)
		require.Equal(t, &ResponseGetPid{Code: 1}, res)
	}()

	func() {
		err := c.Shutdown("myreason")
		require.NoError(t, err)
	}()

	func() {
		res, err := c.RequestTopic("mytopic", [][]interface{}{{"testing"}})
		require.NoError(t, err)
		require.Equal(t, &ResponseRequestTopic{Code: 1}, res)
	}()

	func() {
		res, err := c.GetBusInfo()
		require.NoError(t, err)
		require.Equal(t, [][]interface{}(nil), res)
	}()
}
