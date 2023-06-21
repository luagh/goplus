// Package sms 发送短信
package sms

import (
	"Goplus/pkg/config"
	"sync"
)

// Message 是短信的结构体
type Message struct {
	Template string
	Data     map[string]string

	Content string
}

// SMS 是我们发送短信的操作类
type SMS struct {
	Driver Driver
}

// once 单例模式
var once sync.Once
var internalSMS *SMS

func NewSMS() *SMS {
	once.Do(func() {
		internalSMS = &SMS{
			Driver: &Aliyun{},
		}
	})
	return internalSMS
}
func (sms *SMS) Send(phone string, message Message) bool {
	return sms.Driver.Send(phone, message, config.GetStringMapString("sms.aliyun"))
}
