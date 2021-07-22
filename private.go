package logging

import (
	"time"

	c "github.com/fatih/color"
)

func statusLog(statusText, text string, sign bool) {
	timeLog := c.New(c.FgHiBlack, c.Bold).PrintfFunc()
	status := c.New(c.FgHiWhite, c.Bold).PrintfFunc()
	message := c.New(c.FgHiGreen, c.Bold).PrintfFunc()
	if !sign {
		status = c.New(c.FgHiRed, c.Bold).PrintfFunc()
		message = c.New(c.FgHiRed, c.Bold).PrintfFunc()
	}

	timeLog("[%s] ", time.Now().Format("2006-01-02 15:04:05"))
	status("%s: [ ", statusText)
	message(text)
	status(" ]\n")
}

func errorLog(errorText, text string) {
	timeLog := c.New(c.FgHiBlack, c.Bold).PrintfFunc()
	errLog := c.New(c.FgRed, c.Bold).PrintfFunc()
	err := c.New(c.FgHiRed).PrintfFunc()

	timeLog("[%s] ", time.Now().Format("2006-01-02 15:04:05"))
	errLog("%s: [ ", errorText)
	err(text)
	errLog(" ]\n")
}

func nowTimeLog() {
	tmLog := c.New(c.FgHiBlack, c.Bold).PrintfFunc()
	tm := c.New(c.FgHiBlack, c.Bold).PrintFunc()

	tm("*TIME*")
	tmLog(" { %s } ", time.Now().Format("2006-01-02 15:04:05"))
	tm("*TIME*\n")
}

func errorsLog(a ...interface{}) {
	timeLog := c.New(c.FgHiBlack, c.Bold).PrintfFunc()
	errLog := c.New(c.FgHiRed, c.Bold).PrintFunc()
	errors := c.New(c.FgHiRed, c.FgHiWhite, c.Bold).PrintFunc()

	timeLog("[%s] ", time.Now().Format("2006-01-02 15:04:05"))
	errLog("*[ERROR]*")
	for i := range a {
		errors(" ", a[i])
	}
	errLog(" *[ERROR]*\n")
}

func infoLog(a ...interface{}) {
	timeLog := c.New(c.FgHiBlack, c.Bold).PrintfFunc()
	info := c.New(c.FgHiCyan, c.Bold).PrintFunc()
	text := c.New(c.FgCyan).PrintFunc()

	timeLog("[%s] ", time.Now().Format("2006-01-02 15:04:05"))
	info("*[INFO]*")
	for i := range a {
		text(" ", a[i])
	}
	info(" *[INFO]*\n")
}

func debugLog(a ...interface{}) {
	debug := c.New(c.FgHiMagenta, c.Bold).PrintFunc()
	text := c.New(c.FgMagenta).PrintFunc()

	debug("*[DEBUG]*")
	for i := range a {
		text(" ", a[i])
	}
	debug(" *[DEBUG]*\n")
}

func largeDebugLog(a ...interface{}) {
	debug := c.New(c.FgHiMagenta, c.Bold).PrintFunc()
	text := c.New(c.FgMagenta).PrintFunc()

	debug("******************************* [ DEBUG ] *******************************\n\n")
	for i := range a {
		text(" ", a[i])
	}
	debug("\n\n******************************* [ DEBUG ] *******************************\n")
}

func warningLog(a ...interface{}) {
	timeLog := c.New(c.FgHiBlack, c.Bold).PrintfFunc()
	warning := c.New(c.FgHiYellow, c.Bold).PrintFunc()
	text := c.New(c.FgYellow).PrintFunc()

	timeLog("[%s] ", time.Now().Format("2006-01-02 15:04:05"))
	warning("*[WARNING]*")
	for i := range a {
		text(" ", a[i])
	}
	warning(" *[WARNING]*\n")
}
