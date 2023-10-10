package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
	"wusthelper-mp-gateway/library/ecode"
	respCode "wusthelper-mp-gateway/library/ecode/resp"
)

type undergradLoginReq struct {
	UserAccount  string `json:"userAccount"`
	UserPassword string `json:"userPassword"`
}

func undergradLogin(c *gin.Context) {
	platform := getPlatform(c)
	req := new(undergradLoginReq)
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
	_, student, err := serv.UndergradLogin(&ctx, req.UserAccount, req.UserPassword, oid, true, platform)
	if err != nil {
		responseEcode(c, err.(ecode.Code), nil)
		return
	}

	resp := StudentInfoResp{
		StuId:     student.Sid,
		Name:      student.Name,
		Sex:       "Unknown",
		ClassName: student.Clazz,
		College:   student.College,
		Major:     student.Major,
		Year:      "student.Sid[:4]",
	}

	respData := map[string]any{
		"info": resp,
	}

	response(c, respCode.UndergradLoginOk, "ok", respData)
	c.Next()
}

func undergradGetStudentInfo(c *gin.Context) {
	oid, err := getOid(c)
	if err != nil {
		responseEcode(c, ecode.ParamWrong, nil)
		return
	}

	platform := getPlatform(c)
	ctx := c.Request.Context()
	student, err := serv.UndergradGetStudentInfo(&ctx, oid, platform)
	if err != nil {
		responseEcode(c, err.(ecode.Code), nil)
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

	response(c, respCode.UndergradGetStudentInfoOk, "ok", resp)
}

func undergradGetCourseTable(c *gin.Context) {
	oid, err := getOid(c)
	if err != nil {
		responseEcode(c, ecode.ParamWrong, nil)
		return
	}

	term, has := c.GetQuery("term")
	if !has {
		term = _generateCurrentTermText()
	}

	platform := getPlatform(c)
	ctx := c.Request.Context()
	courses, err := serv.UndergradGetCourseTable(&ctx, oid, term, platform)
	if err != nil {
		responseEcode(c, err.(ecode.Code), nil)
		return
	}

	// 数据格式兼容转换
	courseList := make([]CourseRespItem, len(*courses))
	for i, course := range *courses {
		courseList[i] = CourseRespItem{
			Name:      course.ClassName,
			RoomName:  course.Classroom,
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

	response(c, respCode.UndergradGetCoursesOk, "ok", resp)
}

func _generateCurrentTermText() string {
	now := time.Now()
	year := now.Year()
	month, err := strconv.ParseInt(now.Month().String(), 10, 32)
	if err != nil {
		month = 0
	}

	term := 1
	if month >= 7 || month <= 2 {
		term = 1
	} else {
		term = 2
	}

	return fmt.Sprintf("%d-%d-%d", year, year+1, term)
}

func _getWeekCourseCount(courses *[]CourseRespItem) *[]int {
	weekCounts := make([]int, 25)
	for _, course := range *courses {
		for week := course.StartWeek; week <= course.EndWeek; week++ {
			weekCounts[week]++
		}
	}

	return &weekCounts
}

func undergradGetScore(c *gin.Context) {
	oid, err := getOid(c)
	if err != nil {
		responseEcode(c, ecode.ParamWrong, nil)
		return
	}

	platform := getPlatform(c)
	ctx := c.Request.Context()
	scores, err := serv.UndergradGetScore(&ctx, oid, platform)
	if err != nil {
		responseEcode(c, err.(ecode.Code), nil)
		return
	}

	// 数据格式兼容转换，部分信息会有丢失和不准确
	scoreList := make([]ScoreRespItem, len(*scores))
	for i, score := range *scores {
		scoreList[i] = ScoreRespItem{
			Order:         fmt.Sprintf("%d", i),
			Term:          score.SchoolTerm,
			LessonId:      score.CourseNum,
			LessonName:    score.CourseName,
			LessonGroup:   score.CourseNature,
			ScoreNum:      score.Grade,
			GradeMark:     score.ScoreFlag,
			Credit:        fmt.Sprintf("%.2f", score.CourseCredit),
			ClassPeriod:   fmt.Sprintf("%.0f", score.CourseHours),
			Grade:         fmt.Sprintf("%.2f", score.GradePoint),
			ExamType:      score.ExamNature,
			ExamPoperty:   score.ExamNature,
			LessonPoperty: score.CourseNature,
		}
	}

	// 意义不明的一堆字段
	resp := ScoreResp{
		// 不是很明白这个TimeInfo的意义
		TimeInfo: *_getCurrentTimeInfo(),
		ScoreInfoResp: struct {
			GradeInfo ScoreInfo       `json:"gradeInfo"`
			GradeList []ScoreRespItem `json:"gradeList"`
		}{
			GradeInfo: ScoreInfo{
				LessonNum:    0,
				CreditNum:    0,
				AverageGrade: "请选择学期",
				AverageScore: 0,
			},
			GradeList: scoreList,
		},
	}

	response(c, respCode.UndergradGetScoreOk, "ok", resp)
}

func _getCurrentTimeInfo() *[]int {
	timeInfo := make([]int, 6)
	now := time.Now()
	year, month, day := now.Date()
	timeInfo[0] = year
	m, _ := strconv.ParseInt(month.String(), 10, 32)
	timeInfo[1] = int(m)
	timeInfo[2] = day
	timeInfo[3] = now.Hour()
	timeInfo[4] = now.Minute()
	timeInfo[5] = now.Second()

	return &timeInfo
}

func undergradGetTrainingPlan(c *gin.Context) {
	oid, err := getOid(c)
	if err != nil {
		responseEcode(c, ecode.ParamWrong, nil)
		return
	}

	platform := getPlatform(c)
	ctx := c.Request.Context()
	page, err := serv.UndergradGetTrainingPlan(&ctx, oid, platform)
	if err != nil {
		responseEcode(c, err.(ecode.Code), nil)
		return
	}

	response(c, respCode.UndergradGetTrainingPlanOk, "ok", page)
}
