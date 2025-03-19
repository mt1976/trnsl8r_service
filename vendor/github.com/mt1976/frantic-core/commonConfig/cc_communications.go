package commonConfig

import "strconv"

func (s *Settings) GetCommunicationsPushover_UserKey() string {
	return s.Communications.Pushover.UserKey
}

func (s *Settings) GetCommunicationsPushover_APIToken() string {
	return s.Communications.Pushover.APIToken
}

func (s *Settings) GetCommunicationsEmail_Host() string {
	return s.Communications.Email.Host
}

func (s *Settings) GetCommunicationsEmail_Port() int {
	return s.Communications.Email.Port
}

func (s *Settings) GetCommunicationsEmail_User() string {
	return s.Communications.Email.User
}

func (s *Settings) GetCommunicationsEmail_Password() string {
	return s.Communications.Email.Password
}

func (s *Settings) GetCommunicationsEmail_Sender() string {
	return s.Communications.Email.From
}

func (s *Settings) GetCommunicationsEmail_Footer() string {
	return s.Communications.Email.Footer
}

func (s *Settings) GetCommunicationsEmail_PortString() string {
	return strconv.Itoa(s.GetCommunicationsEmail_Port())
}

func (s *Settings) GetCommunicationsEmail_AdminEmail() string {
	return s.Communications.Email.Admin
}
