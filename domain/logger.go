package domain

type Logger interface {
	Printf(format string, v ...interface{})
}
