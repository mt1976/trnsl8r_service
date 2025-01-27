package date

import (
	"strings"
	"time"

	"github.com/mt1976/trnsl8r_service/app/support/config"
	"github.com/mt1976/trnsl8r_service/app/support/logger"
)

var name = "DATE"
var Format DateFormat
var cfg *config.Configuration

type DateFormat struct {
	External     string
	DMY          string
	Internal     string
	Detail       string
	YMD          string
	Calendar     string
	BackupDate   string
	BackupFolder string
}

func init() {
	cfg = config.Get()

	Format.External = cfg.DateFormatHuman()
	Format.DMY = cfg.DateFormatDMY2()
	Format.Internal = cfg.DateFormatHuman()
	Format.Detail = cfg.DateFormatDateTime()
	Format.YMD = cfg.DateFormatYMD()
	//Format.Calendar = "2006-01-02T15:04:05"
	Format.BackupDate = cfg.DateFormatBackup()
	Format.BackupFolder = cfg.DateFormatBackupFolder()
}

// func FormatYMD(in time.Time) string {
// 	return in.Format(Format.YMD)
// }

func FormatAudit(in time.Time) string {
	return in.Format(Format.Detail)
}

func FormatDMY(in time.Time) string {
	if in.IsZero() {
		return ""
	}
	return in.Format(Format.DMY)
}

func FormatYMD(in time.Time) string {
	if in.IsZero() {
		return ""
	}
	return in.Format(Format.YMD)
}

func FormatCalendar(in time.Time) string {
	if in.IsZero() {
		return ""
	}
	return in.Format(Format.Calendar)
}

func FormatHumanFromString(in string) (time.Time, error) {
	if in == "" {
		return time.Time{}, nil
	}
	return time.Parse(Format.External, in)
}

func FormatHuman(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format(Format.External)
}

func StartOfDay(t time.Time) time.Time {
	// Purpose: To remove the time from a date
	w := t.Format(Format.DMY)
	r, err := time.Parse(Format.DMY, w)
	if err != nil {
		logger.WarningLogger.Printf("[%v] Error=[%v]", strings.ToUpper(name), err.Error())
		return t
	}
	if cfg.ApplicationModeIs(config.MODE_DEVELOPMENT) {
		logger.InfoLogger.Printf("[%v] [DateStartOfDay] Date=[%v] Result=[%v]", strings.ToUpper(name), t, r)
	}
	return r
}

func EndOfDay(t time.Time) time.Time {
	// Purpose: To remove the time from a date
	w := t.Format(Format.DMY)
	r, err := time.Parse(Format.DMY, w)
	if err != nil {
		logger.WarningLogger.Printf("[%v] Error=[%v]", strings.ToUpper(name), err.Error())
		return t
	}
	r = r.AddDate(0, 0, 1)
	r = r.Add(-time.Second)
	if cfg.ApplicationModeIs(config.MODE_DEVELOPMENT) {
		logger.InfoLogger.Printf("[%v] [DateEndOfDay] Date=[%v] Result=[%v]", strings.ToUpper(name), t, r)
	}
	return r
}

func IsBeforeOrEqualTo(t1, t2 time.Time) bool {
	// Purpose: check if a time is before or equal to a time
	if cfg.ApplicationModeIs(config.MODE_DEVELOPMENT) {
		//	logger.InfoLogger.Printf("HLP: [HELPER] Date=[%v] Check=[%v]", DateStartOfDay(t1), DateStartOfDay(t2))
	}
	check := StartOfDay(t1)
	if check.Before(StartOfDay(t2)) || check.Equal(StartOfDay(t2)) {
		//	logger.InfoLogger.Printf("HLP: [HELPER] Date=[%v] Check=[%v] Result=[%v]", DateStartOfDay(t1), DateStartOfDay(t2), true)
		return true
	}
	//logger.InfoLogger.Printf("HLP: [HELPER] Date=[%v] Check=[%v] Result=[%v]", DateStartOfDay(t1), DateStartOfDay(t2), false)
	return false
}

func IsAfterOrEqualTo(t1, t2 time.Time) bool {
	// Purpose: check if a time is after or equal to a time
	if t1.After(StartOfDay(t2)) || t1.Equal(StartOfDay(t2)) {
		return true
	}
	return false
}
