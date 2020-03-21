package log

import (
	"flag"

	"github.com/golang/glog"
)

type Logger interface {
	Fatal(args ...interface{})
	Panic(args ...interface{})
	Error(args ...interface{})
	Wran(args ...interface{})
	Info(args ...interface{})
	Debug(args ...interface{})

	Fatalf(format string, args ...interface{})
	Panicf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Wranf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Debugf(format string, args ...interface{})
}

var logger Logger

func init() {
	flag.Parse()
	defer glog.Flush()
}

func Init(lg Logger) {
	logger = lg
}

func Panicf(format string, args ...interface{}) {
	// if !strings.HasSuffix(format, "\n") {
	// 	format += "\n"
	// }
	if logger == nil {
		defer glog.Flush()
		glog.Errorf(format, args...)
		return
	}
	logger.Panicf(format, args...)
}

func Panic(args ...interface{}) {
	if len(args) == 0 {
		return
	}
	if logger == nil {
		// log.Panic(args...)
		defer glog.Flush()
		glog.Error(args)
	}
	logger.Panic(args...)
}

func Fatal(args ...interface{}) {
	if len(args) == 0 {
		return
	}
	if logger == nil {
		// log.Fatal(args...)
		defer glog.Flush()
		glog.Fatal(args...)
	}
	logger.Fatal(args...)
}

func Fatalf(format string, args ...interface{}) {
	// if !strings.HasSuffix(format, "\n") {
	// 	format += "\n"
	// }
	if logger == nil {
		// log.Fatalf(format, args...)
		defer glog.Flush()
		glog.Fatalf(format, args...)
		return
	}
	logger.Fatalf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	// if !strings.HasSuffix(format, "\n") {
	// 	format += "\n"
	// }
	if logger == nil {
		defer glog.Flush()
		glog.Errorf(format, args...)
		return
	}
	logger.Errorf(format, args...)
}

func Error(args ...interface{}) {
	if len(args) == 0 {
		return
	}
	if logger == nil {
		// log.Println(args...)
		defer glog.Flush()
		glog.Error(args...)
		return
	}
	logger.Error(args...)
}

func Warnf(format string, args ...interface{}) {
	// if !strings.HasSuffix(format, "\n") {
	// 	format += "\n"
	// }
	if logger == nil {
		// log.Printf(format, args...)
		defer glog.Flush()
		glog.Warningf(format, args...)
		return
	}
	logger.Wranf(format, args...)
}

func Warn(args ...interface{}) {
	if len(args) == 0 {
		return
	}
	if logger == nil {
		// log.Println(args...)
		defer glog.Flush()
		glog.Warning(args...)
		return
	}
	logger.Wran(args...)
}

func Infof(format string, args ...interface{}) {
	// if !strings.HasSuffix(format, "\n") {
	// 	format += "\n"
	// }
	if logger == nil {
		// log.Printf(format, args...)
		defer glog.Flush()
		glog.Infof(format, args...)
		return
	}
	logger.Infof(format, args...)
}
func Info(args ...interface{}) {
	if len(args) == 0 {
		return
	}
	if logger == nil {
		// log.Println(args...)
		defer glog.Flush()
		glog.Info(args...)
		return
	}
	logger.Info(args...)
}

func Debugf(format string, args ...interface{}) {
	// if !strings.HasSuffix(format, "\n") {
	// 	format += "\n"
	// }
	if logger == nil {
		// log.Printf(format, args...)
		defer glog.Flush()
		glog.Infof(format, args...)
		return
	}
	logger.Debugf(format, args...)
}

func Debug(args ...interface{}) {
	if len(args) == 0 {
		return
	}
	if logger == nil {
		// log.Println(args...)
		defer glog.Flush()
		glog.Info(args...)
		return
	}
	logger.Debug(args...)
}
