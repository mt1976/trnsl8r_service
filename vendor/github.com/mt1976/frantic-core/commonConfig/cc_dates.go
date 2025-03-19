package commonConfig

func (s *Settings) GetDateFormat_DateTime() string {
	if s.Dates.Formats.DateTime == "" {
		return "2006-01-02 15:04:05"
	}
	return s.Dates.Formats.DateTime
}

func (s *Settings) GetDateFormat_Date() string {
	if s.Dates.Formats.Date == "" {
		return "02/01/2006"
	}
	return s.Dates.Formats.Date
}

func (s *Settings) GetDateFormat_Time() string {
	if s.Dates.Formats.Time == "" {
		return "15:04:05"
	}
	return s.Dates.Formats.Time
}

func (s *Settings) GetDateFormat_Backup() string {
	if s.Dates.Formats.Backup == "" {
		return "060102"
	}
	return s.Dates.Formats.Backup
}

func (s *Settings) GetDateFormat_BackupDirectory() string {
	if s.Dates.Formats.BackupFolder == "" {
		return "060102150405"
	}
	return s.Dates.Formats.BackupFolder
}

func (s *Settings) GetDateFormat_Human() string {
	if s.Dates.Formats.Human == "" {
		return "02 Jan 2006"
	}
	return s.Dates.Formats.Human
}

func (s *Settings) GetDateFormat_DMY2() string {
	if s.Dates.Formats.DMY2 == "" {
		return "02/01/06"
	}
	return s.Dates.Formats.DMY2
}

func (s *Settings) GetDateFormat_YMD() string {
	if s.Dates.Formats.YMD == "" {
		return "2006-01-02"
	}
	return s.Dates.Formats.YMD
}

func (s *Settings) GetDateFormat_Internal() string {
	if s.Dates.Formats.Internal == "" {
		return "20060102"
	}
	return s.Dates.Formats.Internal
}
