package testhelper

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"

	logapi "github.com/doug4j/gologr/logapi/v1"
	"github.com/doug4j/gologr/logapi/v1/loggo"
)

// ExpectedErrorMsg is the message to signal to the developer that the test did pass. Typically used for logging to
// signal that a test was indeed successful though an expected error had been caught.
const ExpectedErrorMsg = "OK - Expected error found"

// SetupTest gives a default logging with debug turned on (the 'log') determines the callling function (the 'testName')
func SetupTest() (log logapi.Logging, testName string) {
	testName = GetCallingNameByDepth(2)
	return defaultLogger(logapi.DebugLogging), testName
}

// SetupTest gives a default logging with caller-defined logging level  the 'log') determines the callling function (the 'testName')
func SetupTestWithLogLevel(level logapi.Level) (log logapi.Logging, testName string) {
	testName = GetCallingNameByDepth(2)
	return defaultLogger(level), testName
}

func getCurrentParentPath() string {
	_, filename, _, _ := runtime.Caller(2)
	return filepath.Dir(filename)
}

// GetParentDirPathFromCurSrcFile get the full path to the running source file and returns the parent directory.
// This file has only been tested with "go test" and the results are not known outside of a testing context.
func GetParentDirPathFromCurSrcFile(appendingPaths ...string) string {
	newPath := getCurrentParentPath()
	newPath = filepath.Clean(newPath)
	for _, appendingPath := range appendingPaths {
		newPath = filepath.Join(newPath, appendingPath)
		newPath = filepath.Clean(newPath)
	}
	return newPath
}

func defaultLogger(level logapi.Level) logapi.Logging {
	if level == logapi.InfoLogging {
		return loggo.NewLogAdaptor(level, loggo.NewEmojiMessageHandler(), loggo.NewStdOutPrintln())
	}
	return loggo.NewLogAdaptor(level, loggo.NewLineNumMessageHandler(loggo.LineNumMessageHandlerParms{}), loggo.NewStdOutPrintln())
}

// GetCallingName obtains the name of the calling function from here (that is, using depth of 1)
func GetCallingName() string {
	return GetCallingNameByDepth(1)
}

// GetCallingName obtains the name using a depth
func GetCallingNameByDepth(depth int) string {
	pc, _, _, _ := runtime.Caller(depth)

	fullPCName := runtime.FuncForPC(pc).Name()
	lastIndexOfPc := strings.LastIndex(fullPCName, "/") + 1
	justPcName := fullPCName[lastIndexOfPc:len(fullPCName)]
	lastIndexOfJustName := strings.LastIndex(justPcName, ".") + 1
	justName := justPcName[lastIndexOfJustName:len(justPcName)]

	return justName
}

// StartTest prints out a friendly name for tests at their start.
func StartTest(testName string, log logapi.Logging) {
	log.Info("")
	log.Info("")
	pc, file, line, _ := runtime.Caller(1)

	fullPCName := runtime.FuncForPC(pc).Name()
	lastIndexOfPc := strings.LastIndex(fullPCName, "/") + 1
	// justPcName := fullPCName[lastIndexOfPc:len(fullPCName)]
	justPcName := fullPCName[lastIndexOfPc:]

	lastIndexOfFile := strings.LastIndex(file, "/") + 1
	// justFileName := file[lastIndexOfFile:len(file)]
	justFileName := file[lastIndexOfFile:]

	// log.Printf("INFO [%s:%d] [%s] %v", justFileName, line, justPcName, msg)
	msg := fmt.Sprintf("***START [%s:%d] [%s] %v", justFileName, line, justPcName, testName)
	log.Info(msg)
	// log.Printf("***START " + testName + " [%s:%d] [%s] %v", justFileName, line, justPcName, msg))
	log.Info("")
}

// EndTest prints out a friendly name for tests at their end.
func EndTest(testName string, log logapi.Logging) {
	log.Info("")
	pc, file, line, _ := runtime.Caller(1)

	fullPCName := runtime.FuncForPC(pc).Name()
	lastIndexOfPc := strings.LastIndex(fullPCName, "/") + 1
	// justPcName := fullPCName[lastIndexOfPc:len(fullPCName)]
	justPcName := fullPCName[lastIndexOfPc:]

	lastIndexOfFile := strings.LastIndex(file, "/") + 1
	// justFileName := file[lastIndexOfFile:len(file)]
	justFileName := file[lastIndexOfFile:]

	msg := fmt.Sprintf("***END [%s:%d] [%s] %v", justFileName, line, justPcName, testName)
	log.Info(msg)
	// log.Println("***END " + testName)
	log.Info("")
	log.Info("")
}
