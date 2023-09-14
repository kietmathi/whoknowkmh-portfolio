package domain

// Logger: Represent the Logger contract 
type Logger interface {
	Printf(format string, v ...interface{})
}
