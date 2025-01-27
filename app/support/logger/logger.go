package logger

import (
	"io"
	"log"
	"os"
	"runtime"

	"github.com/mt1976/trnsl8r_service/app/support/colours"
	"github.com/mt1976/trnsl8r_service/app/support/config"
	"github.com/mt1976/trnsl8r_service/app/support/paths"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	WarningLogger     *log.Logger
	InfoLogger        *log.Logger
	ErrorLogger       *log.Logger
	PanicLogger       *log.Logger
	TimingLogger      *log.Logger
	EventLogger       *log.Logger
	ServiceLogger     *log.Logger
	TraceLogger       *log.Logger
	AuditLogger       *log.Logger
	TranslationLogger *log.Logger
	SecurityLogger    *log.Logger
	DatabaseLogger    *log.Logger
	ApiLogger         *log.Logger
)

var Reset string
var Red string
var Green string
var Yellow string
var Blue string
var Magenta string
var Cyan string
var Gray string
var White string

func init() {
	cfg := config.Get()
	//prefix := "data/logs/"
	prefix := paths.Application().String() + paths.Logs().String() + string(os.PathSeparator)
	name := prefix + cfg.ApplicationName() + "-"

	generalWriter := io.MultiWriter(os.Stdout, &lumberjack.Logger{Filename: name + "general.log", MaxSize: 10, MaxBackups: 3, MaxAge: 28, Compress: true})
	timingWriter := io.MultiWriter(&lumberjack.Logger{Filename: name + "timing.log", MaxSize: 10, MaxBackups: 3, MaxAge: 28, Compress: true})
	serviceWriter := io.MultiWriter(os.Stdout, &lumberjack.Logger{Filename: name + "service.log", MaxSize: 10, MaxBackups: 3, MaxAge: 28, Compress: true})
	auditWriter := io.MultiWriter(&lumberjack.Logger{Filename: name + "audit.log", MaxSize: 10, MaxBackups: 3, MaxAge: 28, Compress: true})
	errorWriter := io.MultiWriter(os.Stdout, os.Stderr, &lumberjack.Logger{Filename: name + "error.log", MaxSize: 10, MaxBackups: 3, MaxAge: 28, Compress: true})
	translationWriter := io.MultiWriter(&lumberjack.Logger{Filename: name + "translation.log", MaxSize: 10, MaxBackups: 3, MaxAge: 28, Compress: true})
	traceWriter := io.MultiWriter(&lumberjack.Logger{Filename: name + "trace.log", MaxSize: 10, MaxBackups: 3, MaxAge: 28, Compress: true})
	warningWriter := io.MultiWriter(os.Stdout, &lumberjack.Logger{Filename: name + "warning.log", MaxSize: 10, MaxBackups: 3, MaxAge: 28, Compress: true})
	eventWriter := io.MultiWriter(os.Stderr, &lumberjack.Logger{Filename: name + "event.log", MaxSize: 10, MaxBackups: 3, MaxAge: 28, Compress: true})
	securityWriter := io.MultiWriter(os.Stderr, &lumberjack.Logger{Filename: name + "security.log", MaxSize: 10, MaxBackups: 3, MaxAge: 28, Compress: true})
	databaseWriter := io.MultiWriter(&lumberjack.Logger{Filename: name + "database.log", MaxSize: 10, MaxBackups: 3, MaxAge: 28, Compress: true})
	apiWriter := io.MultiWriter(os.Stdout, &lumberjack.Logger{Filename: name + "api.log", MaxSize: 10, MaxBackups: 3, MaxAge: 28, Compress: true})

	//fmt.Printf("name: %v\n", name)
	//os.Exit(1)
	setColoursNormal()
	if runtime.GOOS == "windows" {
		setColoursWindows()
	}

	msgStructure := log.Ldate | log.Ltime | log.Lshortfile

	InfoLogger = log.New(generalWriter, Cyan+"[INFO       ] "+Reset, msgStructure)

	WarningLogger = log.New(warningWriter, Yellow+"[WARNING !!!] "+Reset, msgStructure)
	ErrorLogger = log.New(errorWriter, Red+"[ERROR !!!!!] "+Reset, msgStructure)
	PanicLogger = log.New(generalWriter, Red+"[PANIC      ] "+Reset, msgStructure)
	TimingLogger = log.New(timingWriter, Blue+"[TIMING     ] "+Reset, msgStructure)
	EventLogger = log.New(eventWriter, Green+"[EVENT      ] "+Reset, msgStructure)
	ServiceLogger = log.New(serviceWriter, Green+"[SERVICE    ] "+Reset, msgStructure)
	TraceLogger = log.New(traceWriter, White+"[TRACE      ] "+Reset, msgStructure)
	AuditLogger = log.New(auditWriter, Yellow+"[AUDIT      ] "+Reset, msgStructure)
	TranslationLogger = log.New(translationWriter, Cyan+"[TRANSLATION] "+Reset, msgStructure)
	SecurityLogger = log.New(securityWriter, Magenta+"[SECURITY   ] "+Reset, msgStructure)
	DatabaseLogger = log.New(databaseWriter, Blue+"[DATABASE   ] "+Reset, msgStructure)
	ApiLogger = log.New(apiWriter, Green+"[API        ] "+Reset, msgStructure)
}

func TestIt() {
	InfoLogger.Println("Starting the application...")
	InfoLogger.Println("Something noteworthy happened")
	WarningLogger.Println("There is something you should know about")
	PanicLogger.Println("Something went wrong")
	ErrorLogger.Println("Something went wrong")
	TimingLogger.Println("Timing")
	EventLogger.Println("Events")
	ServiceLogger.Println("Service")
	TraceLogger.Println("Trace")
	AuditLogger.Println("Audit")
	TranslationLogger.Println("Translation")
	SecurityLogger.Println("Security")
	DatabaseLogger.Println("Database")
	ApiLogger.Println("API")
}

func setColoursNormal() {
	Reset = colours.Reset
	Red = colours.Red
	Green = colours.Green
	Yellow = colours.Yellow
	Blue = colours.Blue
	Magenta = colours.Magenta
	Cyan = colours.Cyan
	Gray = colours.Gray
	White = colours.White
}

func setColoursWindows() {
	Reset = ""
	Red = ""
	Green = ""
	Yellow = ""
	Blue = ""
	Magenta = ""
	Cyan = ""
	Gray = ""
	White = ""
}
