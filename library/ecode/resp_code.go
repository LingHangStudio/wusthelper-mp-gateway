package ecode

const (
	TokenInvalid = 10000

	MpLoginOK              = 20001
	MpCountUserOk          = 20021
	MpDecodeTokenOk        = 10001
	MpDecodeTokenErr       = 10003
	MpDecodeTokenNoStudent = 10006
	MpGetUserInfoOk        = 10000
	MpGetUserInfoFail      = 10001
	MpUserProfileUploadOk  = 20011
	MpGetVersionLogOk      = 20041
	MpGetAdminConfigureOk  = 20051
	MpGetSecretOk          = 20061
	MpGetUnionStatusOk     = 20071

	UndergradLoginOk           = 30031
	UndergradGetCoursesOk      = 30001
	UndergradGetScoreOk        = 30011
	UndergradGetStudentInfoOk  = 30021
	UndergradGetTrainingPlanOk = 30031

	GraduateLoginOk   = 70021
	GraduateRequestOk = 70000
)
