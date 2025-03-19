package commonConfig

// / STATUS
func (s *Settings) GetStatusList() []string {
	statusList := []string{s.GetStatus_Unknown(), s.GetStatus_Online(), s.GetStatus_Offline(), s.GetStatus_Error(), s.GetStatus_Warning()}
	return statusList
}

func (s *Settings) GetStatus_Unknown() string {
	if s.Status.UNKNOWN == "" {
		return "UNKN"
	}
	return s.Status.UNKNOWN
}

func (s *Settings) GetStatus_Online() string {
	if s.Status.ONLINE == "" {
		return "ONLN"
	}
	return s.Status.ONLINE
}

func (s *Settings) GetStatus_Offline() string {
	if s.Status.OFFLINE == "" {
		return "OFLN"
	}
	return s.Status.OFFLINE
}

func (s *Settings) GetStatus_Error() string {
	if s.Status.ERROR == "" {
		return "ERRO"
	}
	return s.Status.ERROR
}

func (s *Settings) GetStatus_Warning() string {
	if s.Status.WARNING == "" {
		return "WARN"
	}
	return s.Status.WARNING
}