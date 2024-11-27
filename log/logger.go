package log

type Logger interface {
	Info(...any)
	Warn(...any)
	Error(...any)
}
