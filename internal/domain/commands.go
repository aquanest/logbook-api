package domain

// Command Names
const (
	ProfileCommandName string = "profile"
	DivelogCommandName string = "divelog"
)

// Command Usages
const (
	ProfileCommandUsage string = "Fetch my profile."
	DivelogCommandUsage string = "Fetch my divelogs."
)

// Command Aliases
var (
	ProfileCommandAlias = []string{"p", "pr", "pro", "prof"}
	DivelogCommandAlias = []string{"d", "di", "div", "dive"}
)
