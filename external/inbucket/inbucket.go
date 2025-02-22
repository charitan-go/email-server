package inbucket

type InbucketService interface {
	SendEmail() error
}

type inbucketServiceImpl struct {
}

func NewInbucketService() InbucketService {
	return &inbucketServiceImpl{}
}

func (s *inbucketServiceImpl) SendEmail() error {

	// TODO: Implement send email
	return nil
}
