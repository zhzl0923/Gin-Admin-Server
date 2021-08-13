package logger

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	log *zap.Logger
	ctx context.Context
}

// New create logger by opts which can custmoized by command arguments.
func New(opts *Options) *Logger {
	if opts == nil {
		opts = NewOptions()
	}

	var zapLevel zapcore.Level
	if err := zapLevel.UnmarshalText([]byte(opts.Level)); err != nil {
		zapLevel = zapcore.InfoLevel
	}
	core := zapcore.NewCore(getEncoder(opts), getWriteSyncer(opts), zapLevel)

	log := zap.New(core, getZapOptions(opts)...)

	logger := &Logger{
		log: log,
	}

	return logger
}

func getEncoder(opts *Options) zapcore.Encoder {
	encodeLevel := zapcore.CapitalLevelEncoder

	// when output to local path, with color is forbidden
	if opts.Development && opts.EnableColor {
		encodeLevel = zapcore.CapitalColorLevelEncoder
	}

	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "timestamp",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    encodeLevel,
		EncodeTime:     timeEncoder,
		EncodeDuration: milliSecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	var encoder zapcore.Encoder

	if opts.Development {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	}
	return encoder
}

func getWriteSyncer(opts *Options) zapcore.WriteSyncer {
	var writeSyncer []zapcore.WriteSyncer
	for _, w := range opts.WriteSyncer {
		writeSyncer = append(writeSyncer, zapcore.AddSync(w))
	}
	return zapcore.NewMultiWriteSyncer(writeSyncer...)
}

func getZapOptions(opts *Options) []zap.Option {
	var zapOpts []zap.Option

	if !opts.DisableCaller {
		zapOpts = append(zapOpts, zap.AddCaller())
		zapOpts = append(zapOpts, zap.AddCallerSkip(1))
	}

	if !opts.DisableStacktrace {
		zapOpts = append(zapOpts, zap.AddStacktrace(zapcore.PanicLevel))
	}
	return zapOpts
}

func (l *Logger) Debug(msg string, fields ...Field) {
	l.log.Debug(msg, fields...)
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.log.Sugar().Debugf(format, v...)
}

func (l *Logger) Debugw(msg string, keysAndValues ...interface{}) {
	l.log.Sugar().Debugw(msg, keysAndValues...)
}

func (l *Logger) Info(msg string, fields ...Field) {
	l.log.Info(msg, fields...)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.log.Sugar().Infof(format, v...)
}

func (l *Logger) Infow(msg string, keysAndValues ...interface{}) {
	l.log.Sugar().Infow(msg, keysAndValues...)
}

func (l *Logger) Warn(msg string, fields ...Field) {
	l.log.Warn(msg, fields...)
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	l.log.Sugar().Warnf(format, v...)
}

func (l *Logger) Warnw(msg string, keysAndValues ...interface{}) {
	l.log.Sugar().Warnw(msg, keysAndValues...)
}

func (l *Logger) Error(msg string, fields ...Field) {
	l.log.Error(msg, fields...)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.log.Sugar().Errorf(format, v...)
}

func (l *Logger) Errorw(msg string, keysAndValues ...interface{}) {
	l.log.Sugar().Errorw(msg, keysAndValues...)
}

func (l *Logger) Panic(msg string, fields ...Field) {
	l.log.Panic(msg, fields...)
}

func (l *Logger) Panicf(format string, v ...interface{}) {
	l.log.Sugar().Panicf(format, v...)
}

func (l *Logger) Panicw(msg string, keysAndValues ...interface{}) {
	l.log.Sugar().Panicw(msg, keysAndValues...)
}

func (l *Logger) Fatal(msg string, fields ...Field) {
	l.log.Fatal(msg, fields...)
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.log.Sugar().Fatalf(format, v...)
}

func (l *Logger) Fatalw(msg string, keysAndValues ...interface{}) {
	l.log.Sugar().Fatalw(msg, keysAndValues...)
}

func (l *Logger) Flush() {
	_ = l.log.Sync()
}

func (l *Logger) clone() *Logger {
	nl := *l
	return &nl
}

func (l *Logger) WithContext(ctx context.Context) *Logger {
	nl := l.clone()
	nl.ctx = ctx
	return nl
}
