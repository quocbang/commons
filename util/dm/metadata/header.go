package metadata

// metadata header.
const (
	Token       = "Token"
	UserID      = "User-Id"
	ProgramType = "Prg-Type"
	ProgramID   = "Prg-No"
	PID         = "Pid"
	MaxRecords  = "Max-Records"
	Timezone    = "Timezone"
)

// IsHeader asserts the header whether the header of the metadata that DM requires is or not.
func IsHeader(h string) bool {
	switch h {
	case Token, UserID, ProgramType, ProgramID, PID,
		MaxRecords, Timezone:
		return true
	}
	return false
}
