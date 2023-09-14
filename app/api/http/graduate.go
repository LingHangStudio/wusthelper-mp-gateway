package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"wusthelper-mp-gateway/library/ecode"
	"wusthelper-mp-gateway/library/log"
)

type graduateLoginReq struct {
	UserAccount  string `json:"userAccount"`
	UserPassword string `json:"userPassword"`
}

func graduateLogin(c *gin.Context) {
	platform := getPlatform(c)
	req := new(graduateLoginReq)
	err := c.BindJSON(req)
	if err != nil {
		responseEcode(c, ecode.ParamWrong, nil)
		return
	}

	oid, err := getOid(c)
	if err != nil {
		responseEcode(c, ecode.ParamWrong, nil)
		return
	}

	ctx := c.Request.Context()
	_, student, err := serv.GraduateLogin(&ctx, req.UserAccount, req.UserPassword, oid, true, platform)
	if err != nil {
		return
	}

	resp := GraduateStudentInfoResp{
		StuId:       student.Sid,
		StuName:     student.Name,
		StuGrade:    student.Sid[:4],
		MentorName:  "🥰",
		StuCategory: student.Clazz,
		Department:  student.College,
		Major:       student.Major,
	}

	response(c, ecode.GraduateLoginOk, "ok", resp)
	c.Next()
}

func graduateGetStudentInfo(c *gin.Context) {
	oid, err := getOid(c)
	if err != nil {
		responseEcode(c, ecode.ParamWrong, nil)
		return
	}
	platform := getPlatform(c)
	ctx := c.Request.Context()
	student, err := serv.GraduateGetStudentInfo(&ctx, oid, platform)
	if err != nil {
		responseEcode(c, ecode.ServerErr, nil)
		return
	}

	resp := StudentInfoResp{
		StuId:     student.Sid,
		Name:      student.Name,
		ClassName: student.Clazz,
		College:   student.College,
		Major:     student.Major,
		Year:      student.Sid[:4],
	}

	response(c, ecode.GraduateRequestOk, "ok", resp)
}

func graduateGetCourseTable(c *gin.Context) {
	oid, err := getOid(c)
	if err != nil {
		responseEcode(c, ecode.ParamWrong, nil)
		return
	}

	platform := getPlatform(c)
	ctx := c.Request.Context()
	courses, err := serv.GraduateGetCourseTable(&ctx, oid, platform)
	if err != nil {
		responseEcode(c, ecode.ServerErr, nil)
		return
	}

	// 数据格式兼容转换
	courseList := make([]CourseRespItem, len(*courses))
	for i, course := range *courses {
		courseList[i] = CourseRespItem{
			Name:      course.ClassName,
			Room:      course.Classroom,
			Day:       course.WeekDay,
			Length:    2,
			Teacher:   course.Teacher,
			StartWeek: course.StartWeek,
			EndWeek:   course.EndWeek,
			StartTime: course.Section*2 - 1,
		}
	}
	resp := CourseResp{
		// todo 记得一定要想办法改一下，最好可以是加一个接口来修改这些以前的旧配置
		TermStartDate:     "2023-09-04",
		LessonData:        courseList,
		WeekLessonNumList: *_getWeekCourseCount(&courseList),
	}

	response(c, ecode.GraduateRequestOk, "ok", resp)
}

func graduateGetScore(c *gin.Context) {
	oid, err := getOid(c)
	if err != nil {
		responseEcode(c, ecode.ParamWrong, nil)
		return
	}

	platform := getPlatform(c)
	ctx := c.Request.Context()
	scores, err := serv.GraduateGetScore(&ctx, oid, platform)
	if err != nil {
		responseEcode(c, ecode.ServerErr, nil)
		return
	}

	sid, err := serv.GetSid(oid)
	if err != nil {
		log.Error("从oid获取学号时出现异常", zap.String("err", err.Error()))
		sid = ""
	}

	// 数据格式兼容转换，部分信息会有丢失和不准确
	scoreList := make([]GraduateScoreRespItem, len(*scores))
	for i, score := range *scores {
		scoreList[i] = GraduateScoreRespItem{
			AchievementId: i,
			StuNum:        sid,
			CourseName:    score.Name,
			Credit:        score.Credit,
			Term:          fmt.Sprintf("%d", score.Term),
			Score:         score.Point,
		}
	}

	response(c, ecode.GraduateRequestOk, "ok", scoreList)
}
