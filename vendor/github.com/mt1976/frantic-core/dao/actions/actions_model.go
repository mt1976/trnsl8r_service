package actions

import "strings"

type Action struct {
	Name        string `storm:"Name" csv:"Name"`
	userDefined bool   `storm:"userDefined" csv:"userDefined"`
	Code        string `storm:"Code" csv:"Code"`
}

var LIST Action = Action{Name: "List", userDefined: false, Code: "LIST"}
var VIEW Action = Action{Name: "View", userDefined: false, Code: "VIEW"}
var EDIT Action = Action{Name: "Edit", userDefined: false, Code: "EDIT"}
var UPDATE Action = Action{Name: "Update", userDefined: false, Code: "UPDATE"}
var DELETE Action = Action{Name: "Delete", userDefined: false, Code: "DELETE"}
var CREATE Action = Action{Name: "Create", userDefined: false, Code: "CREATE"}
var ENABLE Action = Action{Name: "Enable", userDefined: false, Code: "ENABLE"}
var DISABLE Action = Action{Name: "Disable", userDefined: false, Code: "DISABLE"}
var GRANT Action = Action{Name: "Grant", userDefined: false, Code: "GRANT"}
var REVOKE Action = Action{Name: "Revoke", userDefined: false, Code: "REVOKE"}
var RESET Action = Action{Name: "Reset", userDefined: false, Code: "RESET"}
var ROUTE Action = Action{Name: "Route", userDefined: false, Code: "ROUTE"}
var MESSAGE Action = Action{Name: "Message", userDefined: false, Code: "MESSAGE"}
var LOGIN Action = Action{Name: "Login", userDefined: false, Code: "LOGIN"}
var LOGOUT Action = Action{Name: "Logout", userDefined: false, Code: "LOGOUT"}
var TIMEOUT Action = Action{Name: "Timeout", userDefined: false, Code: "TIMEOUT"}
var API Action = Action{Name: "API", userDefined: false, Code: "API"}
var GET Action = Action{Name: "GET", userDefined: false, Code: "GET"}
var GETALL Action = Action{Name: "GETALL", userDefined: false, Code: "GETALL"}
var EXPORT Action = Action{Name: "EXPORT", userDefined: false, Code: "EXPORT"}
var IMPORT Action = Action{Name: "Import", userDefined: false, Code: "IMPORT"}
var PROCESS Action = Action{Name: "Process", userDefined: false, Code: "PROCESS"}
var REPAIR Action = Action{Name: "Repair", userDefined: false, Code: "REPAIR"}
var AUDIT Action = Action{Name: "Audit", userDefined: false, Code: "AUDIT"}
var CONNECT Action = Action{Name: "Connect", userDefined: false, Code: "CONNECT"}
var DISCONNECT Action = Action{Name: "Disconnect", userDefined: false, Code: "DISCONNECT"}
var BACKUP Action = Action{Name: "Backup", userDefined: false, Code: "BACKUP"}
var VALIDATE Action = Action{Name: "Validate", userDefined: false, Code: "VALIDATE"}
var INITIALISE Action = Action{Name: "Initialise", userDefined: false, Code: "INITIALISE"}
var SHUTDOWN Action = Action{Name: "Shutdown", userDefined: false, Code: "SHUTDOWN"}
var RESTART Action = Action{Name: "Restart", userDefined: false, Code: "RESTART"}
var RELOAD Action = Action{Name: "Reload", userDefined: false, Code: "RELOAD"}
var CLEAR Action = Action{Name: "Clear", userDefined: false, Code: "CLEAR"}
var LOOKUP Action = Action{Name: "Lookup", userDefined: false, Code: "LOOKUP"}
var COUNT Action = Action{Name: "Count", userDefined: false, Code: "COUNT"}
var MAINTENANCE Action = Action{Name: "Maintenance", userDefined: false, Code: "MAINTENANCE"}
var START Action = Action{Name: "Start", userDefined: false, Code: "START"}
var STOP Action = Action{Name: "Stop", userDefined: false, Code: "STOP"}
var PAUSE Action = Action{Name: "Pause", userDefined: false, Code: "PAUSE"}
var SCHEDULE Action = Action{Name: "Schedule", userDefined: false, Code: "SCHEDULE"}
var RUN Action = Action{Name: "Run", userDefined: false, Code: "RUN"}
var EXECUTE Action = Action{Name: "Execute", userDefined: false, Code: "EXECUTE"}
var SUBMIT Action = Action{Name: "Submit", userDefined: false, Code: "SUBMIT"}
var APPROVE Action = Action{Name: "Approve", userDefined: false, Code: "APPROVE"}
var REJECT Action = Action{Name: "Reject", userDefined: false, Code: "REJECT"}
var CANCEL Action = Action{Name: "Cancel", userDefined: false, Code: "CANCEL"}
var CLOSE Action = Action{Name: "Close", userDefined: false, Code: "CLOSE"}
var OPEN Action = Action{Name: "Open", userDefined: false, Code: "OPEN"}
var CLONE Action = Action{Name: "Clone", userDefined: false, Code: "CLONE"}
var NOTIFY Action = Action{Name: "Notify", userDefined: false, Code: "NOTIFY"}

func New(name string) Action {
	return Action{Name: name, userDefined: true}
}

func (bt Action) GetName() string {
	return strings.ToUpper(bt.Name)
}

func (bt Action) GetDescription(in string) string {
	if in == "" {
		return bt.Name
	}
	return bt.Name + " " + in
}

func (bt Action) GetShortName() string {
	return bt.Name
}

func (bt Action) GetCode() string {
	// if len is less than 4, return a suffixed code
	if len(bt.Code) < 4 {
		return (bt.Code + "___")[0:4]
	}
	// return the first 4 characters of the code
	return bt.Code[0:4]
}

func (bt Action) IsUserDefined() bool {
	return bt.userDefined
}

func (bt Action) Is(in Action) bool {
	return bt.Name == in.Name && bt.userDefined == in.userDefined
}
