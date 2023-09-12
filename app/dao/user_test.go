package dao

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestFindWechatUser(t *testing.T) {
	var (
		oid = "1024"
	)

	convey.Convey("FindWechatUser", t, func(ctx convey.C) {
		_, err := dao.FindUserBasic(oid)
		ctx.Convey("The err should be nil and user `should not be` nil", func(ctx convey.C) {
			ctx.So(err, convey.ShouldBeNil)
		})
	})
}

func TestFindQQUser(t *testing.T) {
	var (
		oid = "1024"
	)

	convey.Convey("FindQQUser", t, func(ctx convey.C) {
		_, err := dao.FindUserBasic(oid)
		ctx.Convey("The err should be nil and user `should not be` nil", func(c convey.C) {
			ctx.So(err, convey.ShouldBeNil)
		})
	})
}
