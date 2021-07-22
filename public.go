package logging

// MyLog is my default log messages
type MyLog struct{}

func (l MyLog) Error(errAction string, err error) {
	text := ""
	if err != nil {
		text = err.Error()
	} else {
		text = "..."
	}

	errorLog(errAction, text)
}

// ErrorLog is an error log message
func (l MyLog) ErrorLog(a ...interface{}) {
	errorsLog(a...)
}

// Time is time.Now() log message
func (l MyLog) Time() {
	nowTimeLog()
}

// SuccessStatus is a positive status log message
func (l MyLog) SuccessStatus(action, status string) {
	statusLog(action, status, true)
}

// FailureStatus is a positive status log message
func (l MyLog) FailureStatus(action, status string) {
	statusLog(action, status, false)
}

// Info is a log INFO messge
func (l MyLog) Info(a ...interface{}) {
	infoLog(a...)
}

// Warning is a log WARNING messge
func (l MyLog) Warning(a ...interface{}) {
	warningLog(a...)
}

// Debug is a log DEBUG message
func (l MyLog) Debug(a ...interface{}) {
	debugLog(a...)
}

func (l MyLog) LargeDebug(a ...interface{}) {
	largeDebugLog(a...)
}
