package std

import (
	"os"
	"runtime/debug"

	"github.com/sirupsen/logrus"
)

// LogFields indicates the log's tags
type LogFields map[string]interface{}

var (
	logFile *os.File
)

const (
	// TagTopic flags the topic
	TagTopic = "topic"

	// TopicCodeTrace traces the running of code
	TopicCodeTrace = "code_trace"

	// TopicBugReport indicates the bug report topic
	TopicBugReport = "bug_report"

	// TopicCrash indicates the program's panics
	TopicCrash = "crash"

	// TopicUserActivity indicates the user activity like web access, user login/logout
	TopicUserActivity = "user_activity"

	// TagCategory tags the log category
	TagCategory = "category"

	// TagError tags the error category
	TagError = "error"

	// CategoryRPC indicates the rpc category
	CategoryRPC = "rpc"

	// CategoryRedis indicates the redis category
	CategoryRedis = "redis"

	// CategoryMySQL indicates the MySQL category
	CategoryMySQL = "mysql"

	// CategoryElasticsearch indicates the Elasticsearch category
	CategoryElasticsearch = "elasticsearch"
)

// InitLog initializes the logger
func InitLog(conf ConfigLog) {
	logrus.SetLevel(logrus.Level(conf.Level))
	// logrus.SetFormatter(&logrus.JSONFormatter{})
}

// LogInfo records Info level information which helps trace the running of program and
// moreover the production infos
func LogInfo(fields LogFields, message string) {
	logrus.WithFields(logrus.Fields{
		TagTopic: TopicCodeTrace,
	}).WithFields(map[string]interface{}(fields)).Info(message)
}

// LogInfoc records the running infos
func LogInfoc(category string, message string) {
	logrus.WithFields(logrus.Fields{
		TagTopic:    TopicCodeTrace,
		TagCategory: category,
	}).Info(message)
}

// LogWarn records the warnings which are expected to be removed, but not influence the
// running of the program
func LogWarn(fields LogFields, message string) {
	logrus.WithFields(logrus.Fields{
		TagTopic: TopicBugReport,
	}).WithFields(map[string]interface{}(fields)).Warn(message)
}

// LogError records the running errors which are expected to be solved soon
func LogError(fields LogFields, message string) {
	logrus.WithFields(logrus.Fields{
		TagTopic: TopicBugReport,
	}).WithFields(map[string]interface{}(fields)).Error(message)
}

// LogErrorc records the running errors which are expected to be solved soon
func LogErrorc(category string, err error, message string) {
	logrus.WithFields(logrus.Fields{
		TagTopic:    TopicBugReport,
		TagCategory: category,
		TagError:    err,
	}).Error(message)
}

// LogPanic records the running errors which are expected to be severe soon
func LogPanic(fields LogFields, message string) {
	logrus.WithFields(logrus.Fields{
		TagTopic: TopicBugReport,
	}).WithFields(map[string]interface{}(fields)).Panic(message)
}

// LogInfoLn records Info level information which helps trace the running of program and
// moreover the production infos
func LogInfoLn(args ...interface{}) {
	logrus.Infoln(args)
}

// LogWarnLn records the program warning
func LogWarnLn(args ...interface{}) {
	logrus.WithFields(logrus.Fields{
		TagTopic: TopicBugReport,
	}).Warnln(args)
}

// LogErrorLn records the program error, go to fix it!
func LogErrorLn(args ...interface{}) {
	logrus.WithFields(logrus.Fields{
		TagTopic: TopicBugReport,
	}).Errorln(args)
}

// LogFatalLn records the program fatal error, developer should follow immediately
func LogFatalLn(args ...interface{}) {
	logrus.WithFields(logrus.Fields{
		TagTopic: TopicBugReport,
	}).Fatalln(args)
}

// LogPanicLn records the program fatal error, developer should fix otherwise the company dies
func LogPanicLn(args ...interface{}) {
	logrus.WithFields(logrus.Fields{
		TagTopic: TopicBugReport,
	}).Panicln(args)
}

// LogDebugLn records debug information which helps trace the running of program
func LogDebugLn(args ...interface{}) {
	logrus.Debugln(args)
}

// LogDebugc records the running infos
func LogDebugc(category string, message string) {
	logrus.WithFields(logrus.Fields{
		TagTopic:    TopicCodeTrace,
		TagCategory: category,
	}).Debug(message)
}

// LogUserActivity records user activity, like user access page, login/logout
func LogUserActivity(fields LogFields, message string) {
	logrus.WithFields(logrus.Fields{
		TagTopic: TopicUserActivity,
	}).WithFields(map[string]interface{}(fields)).Infoln(message)
}

// LogRecover records when program crashes
func LogRecover(e interface{}) {
	logrus.WithFields(logrus.Fields{
		TagTopic:     TopicCrash,
		"error":      e,
		"stacktrace": string(debug.Stack()),
	}).Errorln("Recovered panic")
}
