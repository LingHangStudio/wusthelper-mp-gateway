package v2

import (
	"bufio"
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

func TestCourses(t *testing.T) {
	fmt.Println("输入测试token：")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	token := scanner.Text()
	fmt.Printf("使用的请求token：'%s'\n", token)
	convey.Convey("GetCourse", t, func(ctx convey.C) {
		courses, err := rpc.UndergradCourses("2023-2024-1", token)
		ctx.Convey("The err should be nil and user `should not be` nil", func(c convey.C) {
			ctx.So(err, convey.ShouldBeNil)
		})

		fmt.Println(courses)
	})
}
