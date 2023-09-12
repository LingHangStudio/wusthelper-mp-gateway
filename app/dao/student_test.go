package dao

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGetSidForWechatUser(t *testing.T) {
	var (
		oid = "1024"
	)

	convey.Convey("FindQQUser", t, func(ctx convey.C) {
		_, err := dao.FindQQUserBasic(oid)
		ctx.Convey("The err should be nil and user `should not be` nil", func(c convey.C) {
			ctx.So(err, convey.ShouldBeNil)
		})
	})
}
