package timing

import (
	"strings"
	"time"

	"github.com/mt1976/trnsl8r_service/app/support/logger"
)

var name = "Timing"

type Stopwatch struct {
	table   string
	action  string
	start   time.Time
	msg     string
	end     time.Time
	duraton time.Duration
}

func Start(table, action, msg string) Stopwatch {
	start := time.Now()

	//logger.InfoLogger.Printf("TIM: %v %v: [%v]", name, msg, start)
	return Stopwatch{table, action, start, msg, time.Now(), time.Duration(0)}
}

func (w *Stopwatch) Stop(count int) {
	w.end = time.Now()
	w.duraton = w.end.Sub(w.start)
	logger.TimingLogger.Printf("Object=[%v] Action=[%v] Msg=[%v] Count=[%v] Duration=[%v]", w.table, strings.ToUpper(w.action), w.msg, count, w.duraton)
}
