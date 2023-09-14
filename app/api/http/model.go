package http

type wxUserProfileUploadReq struct {
	Oid       string `json:"wechat_openid"`
	Nickname  string `json:"nickname"`
	Gender    int32  `json:"gender"`
	City      string `json:"city"`
	Province  string `json:"province"`
	Country   string `json:"country"`
	AvatarUrl string `json:"avatarUrl"`
}

type qqUserProfileUploadReq struct {
	Oid       string `json:"qq_openid"`
	Nickname  string `json:"nickname"`
	Gender    int32  `json:"gender"`
	City      string `json:"city"`
	Province  string `json:"province"`
	Country   string `json:"country"`
	AvatarUrl string `json:"avatarUrl"`
}

// responses

type StudentInfoResp struct {
	StuId     string `json:"stuId"`
	Name      string `json:"name"`
	Sex       string `json:"sex"`
	ClassName string `json:"className"`
	College   string `json:"college"`
	Major     string `json:"major"`
	Year      string `json:"year"`
}

type GraduateStudentInfoResp struct {
	StuId       interface{} `json:"stuId"`
	StuName     interface{} `json:"stu_name"`
	StuGrade    interface{} `json:"stuGrade"`
	MentorName  interface{} `json:"mentorName"`
	StuCategory interface{} `json:"stuCategory"`
	Department  interface{} `json:"department"`
	Major       interface{} `json:"major"`
}

type CourseRespItem struct {
	Name      string `json:"name"`
	Room      string `json:"room"`
	Day       int    `json:"day"`
	Length    int    `json:"length"`
	Teacher   string `json:"teacher"`
	StartWeek int    `json:"startWeek"`
	EndWeek   int    `json:"endWeek"`
	StartTime int    `json:"startTime"`
}

type CourseResp struct {
	TermStartDate     string           `json:"termStartDate"`
	LessonData        []CourseRespItem `json:"lessonData"`
	WeekLessonNumList []int            `json:"weekLessonNumList"`
}

type ScoreRespItem struct {
	Order         string `json:"order"`
	Term          string `json:"term"`
	LessonId      string `json:"lessonId"`
	LessonName    string `json:"lessonName"`
	LessonGroup   string `json:"lessonGroup"`
	ScoreNum      string `json:"scoreNum"`
	GradeMark     string `json:"gradeMark"`
	Credit        string `json:"credit"`
	ClassPeriod   string `json:"classPeriod"`
	Grade         string `json:"grade"`
	MakeupTerm    string `json:"makeupTerm"`
	ExamType      string `json:"examType"`
	ExamPoperty   string `json:"examPoperty"`
	LessonPoperty string `json:"lessonPoperty"`
	LessonType    string `json:"lessonType"`
}

type ScoreInfo struct {
	LessonNum    int    `json:"lessonNum"`
	CreditNum    int    `json:"creditNum"`
	AverageGrade string `json:"averageGrade"`
	AverageScore int    `json:"averageScore"`
}

type ScoreResp struct {
	TimeInfo      []int `json:"timeArray"`
	ScoreInfoResp struct {
		GradeInfo ScoreInfo       `json:"gradeInfo"`
		GradeList []ScoreRespItem `json:"gradeList"`
	} `json:"grade"`
}

type GraduateScoreRespItem struct {
	AchievementId int     `json:"achievementId"`
	StuNum        string  `json:"stuNum"`
	CourseName    string  `json:"courseName"`
	Credit        float64 `json:"credit"`
	Term          string  `json:"term"`
	Score         string  `json:"score"`
}
