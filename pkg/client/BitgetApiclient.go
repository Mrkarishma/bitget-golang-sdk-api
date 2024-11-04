package client

import (
	"github.com/Mrkarishma/bitget-golang-sdk-api/internal"
	"github.com/Mrkarishma/bitget-golang-sdk-api/internal/common"
)

type BitgetApiClient struct {
	BitgetRestClient *common.BitgetRestClient
}

func NewBitgetApiClient(ApiKey, SecretKey, PASSPHRASE, BaseUrl string, TimeoutSecond int) *BitgetApiClient{
	p := new(BitgetApiClient)
    p.BitgetRestClient = common.NewBitgetRestClient(ApiKey, SecretKey, PASSPHRASE, BaseUrl, TimeoutSecond)
    return p
}

func (p *BitgetApiClient) Init() *BitgetApiClient {
	p.BitgetRestClient = new(common.BitgetRestClient).Init()
	return p
}

func (p *BitgetApiClient) Post(url string, params map[string]string) (string, error) {
	postBody, jsonErr := internal.ToJson(params)
	if jsonErr != nil {
		return "", jsonErr
	}
	resp, err := p.BitgetRestClient.DoPost(url, postBody)
	return resp, err
}

func (p *BitgetApiClient) Get(url string, params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet(url, params)
	return resp, err
}

func (p *BitgetApiClient) Delete(url string, params map[string]string) (string, error) {