package domain

// Global: Name Flags
const (
	ServerFlagName string = "server"
)

// Global: Usage Flags
const (
	ServerFlagUsage string = "URI to ATMOS API"
)

// Global: Alias Flags
var (
	ServerFlagAlias = []string{"s"}
)

// Global: Value Flags
const (
	ServerFlagDefaultValue string = "http://localhost/"
)

// Name Flags
const (
	WebListenAddressFlagName  string = "web.listen-address"
	WebListenPortFlagName     string = "web.listen-port"
	WebDebugFlagName          string = "web.debug"
	ConfigAtmosServerFlagName string = "config.atmos.server"
	ConfigAtmosTokenFlagName  string = "config.atmos.token"
)

// Usage Flags
const (
	WebListenAddressFlagUsage  string = "Set IP address"
	WebListenPortFlagUsage     string = "Set port number"
	WebDebugFlagUsage          string = "Set debug mode"
	ConfigAtmosServerFlagUsage string = "Set FQDN for ATMOS Platform"
	ConfigAtmosTokenFlagUsage  string = "Set access-tokens to use ATMOS Platform API"
)

// Value Flags
const (
	WebListenAddressFlagDefaultValue string = "0.0.0.0"
	WebListenPortFlagDefaultValue    int    = 8080
	WebDebugFlagDefaultValue         bool   = false
)

// Required Flags
const (
	ConfigAtmosServerFlagRequired bool = true
	ConfigAtmosTokenFlagRequired  bool = true
)

// Flag aliases
var (
	WebListenAddressFlagAlias  = []string{"L"}
	WebListenPortFlagAlias     = []string{"P"}
	WebDebugFlagAlias          = []string{"D"}
	ConfigAtmosServerFlagAlias = []string{"s"}
	ConfigAtmosTokenFlagAlias  = []string{"t"}
)

// Flag envvars
var (
	ConfigAtmosServerFlagEnvVars = []string{"LOGBOOK_ATMOS_API"}
	ConfigAtmosTokenFlagEnvVars  = []string{"LOGBOOK_ATMOS_TOKEN"}
)
