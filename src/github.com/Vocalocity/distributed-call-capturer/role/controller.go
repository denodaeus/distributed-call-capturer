package role

import (
	"github.com/Vocalocity/distributed-call-capturer/adapter"
	"log"
)

type Controller struct{}

func (c Controller) Start(callId string) {
	log.Println("start :: starting tracing for callid=" + callId)
	s := adapter.Sip{"127.0.0.1", 5060}
	s.Trace(callId)
}

func (c Controller) Stop(callId string) {
	log.Println("stop :: starting tracing for callid=" + callId)
}

func (c Controller) Stream(client string, callId string) {
	log.Println("stream :: starting tracing for callid=" + callId)
}
