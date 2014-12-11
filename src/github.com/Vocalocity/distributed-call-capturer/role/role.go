package role

type role interface {
	Start() bool
	Stop() bool
	Stream() bool
}
