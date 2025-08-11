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
	"s":      SaveAction,
	"run":    RunAction,
	"r":      RunAction,
	"list":   ListAction,
	"l":      ListAction,
	"find":   FindAction,
	"f":      FindAction,
	"delete": DeleteAction,
	"d":      DeleteAction,
	"help":   HelpAction,
	"h":      HelpAction,
}

var ActionHandlersMap = map[Action]ActionHandler{
	SaveAction:   Save,
	RunAction:    Run,
	ListAction:   List,
	FindAction:   Find,
	DeleteAction: Delete,
	HelpAction:   Help,
}
