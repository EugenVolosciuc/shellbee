package commands

type Action int

type ActionHandler func(args ...string) error

const (
	SaveAction Action = iota
	RunAction
	ListAction
	SearchAction
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
	"find":   SearchAction,
	"f":      SearchAction,
	"delete": DeleteAction,
	"d":      DeleteAction,
	"help":   HelpAction,
	"h":      HelpAction,
}

var ActionHandlersMap = map[Action]ActionHandler{
	SaveAction:   Save,
	RunAction:    Run,
	ListAction:   List,
	SearchAction: Search,
	DeleteAction: Delete,
	HelpAction:   Help,
}
