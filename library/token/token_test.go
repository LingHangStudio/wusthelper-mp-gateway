package token

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
	"time"
)

var (
	_jwt *Token
)

func TestMain(m *testing.M) {
	_jwt = &Token{
		SecretKey: "",
		Timeout:   time.Second * 3600,
	}

	m.Run()
	os.Exit(0)
}

func TestToken(t *testing.T) {
	convey.Convey("SignToken", t, func(ctx convey.C) {
		token := _jwt.Sign("oid")
		ctx.So(token, convey.ShouldNotBeEmpty)
		fmt.Printf("token: %s\n", token)

		valid := _jwt.Verify(token)
		ctx.So(valid, convey.ShouldBeTrue)
		fmt.Println("valid token check pass")

		c := _jwt.GetClaimWithoutVerify(token)
		ctx.So(c, convey.ShouldNotBeNil)
		fmt.Println("token claim pass")
		fmt.Println(c)

		token = ""
		valid = _jwt.Verify(token)
		ctx.So(valid, convey.ShouldBeFalse)
		fmt.Println("invalid token check pass")
	})
}
