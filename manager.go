package sms

import "github.com/cocotyty/summer"

func init() {
	summer.Put(&SMSManager{})
}

func NewManager(appKey, appSecret string) (manager *SMSManager) {
	return &SMSManager{AppKey: appKey, AppSecret: appSecret}
}

type SMSManager struct {
	AppKey    string `sm:"#.ali.appKey"`
	AppSecret string `sm:"#.ali.appSecret"`
}

func (m *SMSManager) Handler(signName, templateCode string) *SMSClient {
	return NewAliSMSClient(signName, templateCode, m.AppKey, m.AppSecret)
}
