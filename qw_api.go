package qwapi

import (
	"context"
	"net/http"
	"time"

	"github.com/dghubble/sling"
)

// 先生成各个服务类

type API struct {
	CorpId string // 企微企业 id
	logger Logger
	client *sling.Sling
}

func (a *API) AccessTokenService(secret string) {

}

func NewAPI(corpId string) *API {
	return &API{CorpId: corpId}
}

func (a API) CustomerService(secret string) {

}

type AccessToken struct {
	Value   string    // accesstoken 值
	Expired time.Time // 过期时间点
}

var AccessTokenGetter func(ctx context.Context, secret string) (AccessToken, error)
