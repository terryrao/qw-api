package qwapi

import (
	"context"

	"github.com/dghubble/sling"
)

// 先生成各个服务类

type API struct {
	CorpId string // 企微企业 id
	logger Logger
	client *sling.Sling
}

func NewAPI(corpId string, options ...APIOptions) *API {
	o := &Option{}
	for _, v := range options {
		v(o)
	}

	if o.httpClient == nil {
		o.httpClient = defaultHttpClient
	}
	if o.logger == nil {
		o.logger = &emptyLogger{}
	}
	return &API{
		CorpId: corpId,
		logger: o.logger,
		client: sling.New().Client(o.httpClient).New().Base(baseURL),
	}
}

func (a API) CustomerService(secret string) {

}

var AccessTokenGetter func(ctx context.Context, secret string) (AccessToken, error)
