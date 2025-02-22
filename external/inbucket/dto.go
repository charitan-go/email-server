package inbucket

type EmailTypeEnum string

const (
	EmailTypeRegisterAccount EmailTypeEnum = "REGISTER_ACCOUNT"
)

type SendEmailRequestDto struct {
	emailType  EmailTypeEnum
	recipients []string
}
