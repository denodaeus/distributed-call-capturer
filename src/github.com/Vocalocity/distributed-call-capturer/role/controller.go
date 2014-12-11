package role

import (
	"log"
)

type Controller struct{}

func (c Controller) Start(callId string) {
	log.Println("start :: starting tracing for callid=" + callId)
}

func (c Controller) Stop(callId string) {
	log.Println("stop :: starting tracing for callid=" + callId)
}

func (c Controller) Stream(client string, callId string) {
	log.Println("stream :: starting tracing for callid=" + callId)
}
