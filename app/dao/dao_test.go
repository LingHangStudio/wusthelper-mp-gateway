package dao

import (
	"os"
	"testing"
	"wusthelper-mp-gateway/app/conf"
)

var (
	dao *Dao
)

func TestMain(m *testing.M) {
	if err := conf.Init(); err != nil {
		panic(err)
	}

	dao = New(conf.Conf)
	m.Run()
	os.Exit(0)
}
