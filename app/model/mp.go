package model

type VersionLog struct {
	Version string `json:"version"`
	Content []struct {
		Tag    string `json:"tag"`
		Date   string `json:"date"`
		Detail string `json:"detail"`
	} `json:"content"`
}

type AdminConfig struct {
	TermList    []string `json:"termList"`
	Openadvance bool     `json:"openadvance"`
	Schedule    struct {
		ScheduleVersion string `json:"scheduleVersion"`
		RefreshSchedule bool   `json:"refreshSchedule"`
	} `json:"schedule"`
	MenuList struct {
		News      bool `json:"news"`
		Volunteer bool `json:"volunteer"`
	} `json:"menuList"`
	JumpUnion  int    `json:"jumpUnion"`
	Banner     bool   `json:"banner"`
	Term       string `json:"term"`
	ShowNotice bool   `json:"showNotice"`
	Union      int    `json:"union"`
}
