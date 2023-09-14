package dao

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGetSid(t *testing.T) {
	var (
		oid = "1024"
	)

	convey.Convey("GetSid", t, func(ctx convey.C) {
		result, err := dao.GetSid(oid)
		ctx.Convey("The err should be nil and user `should not be` nil", func(c convey.C) {
			ctx.So(err, convey.ShouldBeNil)
		})
		fmt.Printf("result: '%s'\n", result)
	})
}
