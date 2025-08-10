package commands

type Action int

type ActionHandler func(args ...string)

const (
	SaveAction Action = iota
	RunAction
	ListAction
	FindAction
	DeleteAction
	HelpAction
)

var ActionsMap = map[string]Action{
	"save":   SaveAction,
	"run":    RunAction,
	"list":   ListAction,
	"find":   FindAction,
	"delete": DeleteAction,
	"help":   HelpAction,
}

var ActionHandlersMap = map[Action]ActionHandler{
	SaveAction:   Save,
	RunAction:    Run,
	ListAction:   List,
	FindAction:   Find,
	DeleteAction: Delete,
	HelpAction:   Help,
}
