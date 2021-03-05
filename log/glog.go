package log

import (
	"strconv"
	"sync/atomic"
)

type severity int32 //定义日志级别

//log level from 0 to 3
// high-severity is also writtten to each loger-severity log file
const (
	infoLog severity =iota
	waringLog
	errorLog
	fatalLog
	numSeverity = 4
)

const severityChar="IWEF" // info waring error fatal

// get returns the value of severity,add atomic to protect the level
func (s *severity) get() severity{
	return severity(atomic.LoadInt32((*int32)(s)))
}

func (s *severity) set(level severity){
	atomic.StoreInt32((*int32)(s),int32(level))
}

func (s *severity) String() string{
	return strconv.FormatInt((int64)(*s),10)
}

func (s *severity) Get() interface{}{
	return *s
}


