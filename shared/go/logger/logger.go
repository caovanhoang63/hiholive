package logger

type Logger interface {
	Debug(msg string)
	DebugWithFields(msg string, fields Field)
	Info(msg string)
	InfoWithFields(msg string, fields Field)
	Warn(msg string)
	WarnWithFields(msg string, fields Field)
	Error(msg string)
	ErrorWithFields(msg string, fields Field)
	Fatal(msg string)
	FatalWithFields(msg string, fields Field)
}

type Field map[string]interface{}
