package loggerc

type Logger interface {
	Debug(...any)
	Debugln(...any)
	Debugf(string, ...any)

	Info(...any)
	Infoln(...any)
	Infof(string, ...any)

	Warn(...any)
	Warnln(...any)
	Warnf(string, ...any)

	Error(...any)
	Errorln(...any)
	Errorf(string, ...any)

	Fatal(...any)
	Fatalln(...any)
	Fatalf(string, ...any)

	Panic(...any)
	Panicln(...any)
	Panicf(string, ...any)

	WithField(Field) Logger
	// add source field to log
	WithSrc() Logger
	GetLevel() string
}

type Field map[string]any

type AppLogger interface {
	GetLogger(prefix string) Logger
}
