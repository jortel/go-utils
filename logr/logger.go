package logr

import (
	"fmt"
	"github.com/go-logr/logr"
	liberr "github.com/jortel/go-utils/error"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
)

const (
	EnvDevelopment = "LOG_DEVELOPMENT"
	EnvJson        = "LOG_JSON"
	EnvLevel       = "LOG_LEVEL"
)

//
// Sink -.
type Sink struct {
	development bool
	delegate    *log.Logger
	fields      log.Fields
	name        string
	json        bool
	level       int
}

//
// WithName returns a named logger.
func WithName(name string, kvpair ...interface{}) logr.Logger {
	return logr.New(&Sink{name: name})
}

func (s *Sink) Init(_ logr.RuntimeInfo) {
	s.delegate = log.New()
	v := os.Getenv(EnvDevelopment)
	b, _ := strconv.ParseBool(v)
	s.development = b
	if s.development {
		fmt := new(log.TextFormatter)
		fmt.TimestampFormat = "2006-01-02 15:04:05"
		fmt.FullTimestamp = true
		s.delegate.SetFormatter(fmt)
	} else {
		fmt := new(log.TextFormatter)
		fmt.FullTimestamp = true
		fmt.DisableColors = true
		fmt.DisableQuote = true
		s.delegate.SetFormatter(fmt)
	}
	v = os.Getenv(EnvJson)
	s.json, _ = strconv.ParseBool(v)
	if s.json {
		fmt := new(log.JSONFormatter)
		fmt.PrettyPrint = true
		s.delegate.Formatter = fmt
	}
	v = os.Getenv(EnvLevel)
	n, _ := strconv.Atoi(v)
	s.level = n
}

//
// Info logs at info.
func (s *Sink) Info(_ int, message string, kvpair ...interface{}) {
	fields := s.asFields(kvpair...)
	entry := s.delegate.WithFields(fields)
	entry.Info(s.named(message))
}

//
// Error logs an error.
func (s *Sink) Error(err error, message string, kvpair ...interface{}) {
	if err == nil {
		return
	}
	xErr, cast := err.(*liberr.Error)
	if cast {
		err = xErr.Unwrap()
		if context := xErr.Context(); context != nil {
			context = append(
				context,
				kvpair...)
			kvpair = context
		}
		if s.json {
			fields := s.asFields(kvpair...)
			fields["error"] = xErr.Error()
			fields["stack"] = xErr.Stack()
			fields["logger"] = s.name
			entry := s.delegate.WithFields(fields)
			entry.Error(s.named(message))
		} else {
			fields := s.asFields(kvpair...)
			entry := s.delegate.WithFields(fields)
			if message != "" {
				entry.Error(s.named(message), "\n", xErr.Error(), xErr.Stack())
			} else {
				entry.Error(s.named(xErr.Error()), xErr.Stack())
			}
		}
	} else {
		if wErr, wrapped := err.(interface {
			Unwrap() error
		}); wrapped {
			err = wErr.Unwrap()
		}
		err = liberr.Wrap(err)
		s.Error(err, message, kvpair...)
	}
}

//
// Enabled returns whether logger is enabled.
func (s *Sink) Enabled(level int) bool {
	return s.level >= level
}

//
// WithName returns a logger with name.
func (s *Sink) WithName(name string) logr.LogSink {
	return &Sink{name: name}
}

//
// WithValues returns a logger with values.
func (s *Sink) WithValues(kvpair ...interface{}) logr.LogSink {
	return &Sink{
		name:   s.name,
		fields: s.asFields(kvpair),
	}
}

func (s *Sink) asFields(kvpair ...interface{}) log.Fields {
	fields := log.Fields{}
	for i := range kvpair {
		if i%2 != 0 {
			key := fmt.Sprintf("%v", kvpair[i-1])
			fields[key] = kvpair[i]
		}
	}
	return fields
}

func (s *Sink) named(message string) (m string) {
	if s.name != "" {
		m = "[" + s.name + "] "
	}
	m = m + message
	return
}
