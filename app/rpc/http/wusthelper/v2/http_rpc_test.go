package v2

import (
	"os"
	"testing"
	"wusthelper-mp-gateway/app/conf"
)

var (
	rpc *WusthelperHttpRpc
)

func TestMain(m *testing.M) {
	if err := conf.Init(); err != nil {
		panic(err)
	}

	rpc = NewRpcClient(&conf.Conf.Wusthelper)
	m.Run()
	os.Exit(0)
}
