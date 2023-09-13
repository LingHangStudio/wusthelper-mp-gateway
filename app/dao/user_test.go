package dao

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestFindWechatUser(t *testing.T) {
	var (
		oid = "1024"
	)

	convey.Convey("FindWechatUser", t, func(ctx convey.C) {
		result, err := dao.GetUserBasic(oid)
		ctx.So(err, convey.ShouldBeNil)
		fmt.Println(result)
	})
}

func TestFindQQUser(t *testing.T) {
	var (
		oid = "1024"
	)

	convey.Convey("FindQQUser", t, func(ctx convey.C) {
		result, err := dao.GetUserBasic(oid)
		ctx.So(err, convey.ShouldBeNil)
		fmt.Println(result)
	})
}
