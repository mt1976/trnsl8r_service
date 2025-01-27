package actions

type Action struct {
	Name        string `storm:"Name" csv:"Name"`
	userDefined bool   `storm:"userDefined" csv:"userDefined"`
}

var LIST Action = Action{Name: "List", userDefined: false}
var VIEW Action = Action{Name: "View", userDefined: false}
var EDIT Action = Action{Name: "Edit", userDefined: false}
var UPDATE Action = Action{Name: "Update", userDefined: false}
var DELETE Action = Action{Name: "Delete", userDefined: false}
var CREATE Action = Action{Name: "Create", userDefined: false}
var ENABLE Action = Action{Name: "Enable", userDefined: false}
var DISABLE Action = Action{Name: "Disable", userDefined: false}
var RESET Action = Action{Name: "Reset", userDefined: false}
var ROUTE Action = Action{Name: "Route", userDefined: false}
var MESSAGE Action = Action{Name: "Message", userDefined: false}
var LOGIN Action = Action{Name: "Login", userDefined: false}
var LOGOUT Action = Action{Name: "Logout", userDefined: false}
var TIMEOUT Action = Action{Name: "Timeout", userDefined: false}
var API Action = Action{Name: "API", userDefined: false}

func New(name string) Action {
	return Action{Name: name, userDefined: true}
}

func (bt Action) GetName() string {
	return bt.Name
}

func (bt Action) IsUserDefined() bool {
	return bt.userDefined
}

func (bt Action) Is(in Action) bool {
	return bt.Name == in.Name && bt.userDefined == in.userDefined
}
