package qwapi

import (
	"context"
	"time"

	"github.com/dghubble/sling"
	"github.com/pkg/errors"
)

type AccessTokenService interface {
	GetAccessToken(ctx context.Context, p Credential) (AccessToken, error)
}

type Credential struct {
	CorpId string `url:"corpid"`
	Secret string `url:"corpsecret"`
}

type AccessToken struct {
	Value   string    // accessToken 值
	Expired time.Time // 过期时间点
}

type BasicAccessTokenService struct {
	client *sling.Sling
}

func (a *BasicAccessTokenService) GetAccessToken(ctx context.Context, p Credential) (AccessToken, error) {
	result := AccessToken{}
	req, err := a.client.QueryStruct(p).Request()
	if err != nil {
		return result, errors.Wrap(err, "get request err")
	}
	if _, err := a.client.Do(req.WithContext(ctx), &result, nil); err != nil {
		return result, errors.Wrap(err, "get response err ")
	}
	return result, nil
}

func newAccessTokenGetter(s *sling.Sling) *BasicAccessTokenService {
	return &BasicAccessTokenService{
		client: s,
	}
}

// mp *accessTokenGetter
