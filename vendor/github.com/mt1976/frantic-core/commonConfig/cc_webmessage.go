package commonConfig

func (s *Settings) GetMessageKey_Type() string {
	return s.Message.Keys.Type
}

func (s *Settings) GetMessageKey_Title() string {
	return s.Message.Keys.Title
}

func (s *Settings) GetMessageKey_Content() string {
	return s.Message.Keys.Content
}

func (s *Settings) GetMessageKey_Action() string {
	return s.Message.Keys.Action
}
