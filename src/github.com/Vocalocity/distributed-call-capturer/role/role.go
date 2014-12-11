package role

type role interface {
	Start(callId string) bool
	Stop(callId string) bool
	Stream(client string, callId string) bool
}
