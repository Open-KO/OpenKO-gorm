package kogen

type Binary []byte

type DbType uint8

const (
	LOGIN DbType = 0
	GAME  DbType = 1
	LOG   DbType = 2
)

var (
	// These can be changed by the importing program for customized output at runtime
	_loginDbName = "ACCOUNT"
	_gameDbName  = "GAME"
	_logDbName   = "LOG"
)

func SetDbNames(loginDbName string, gameDbName string, logDbName string) {
	_loginDbName = loginDbName
	_gameDbName = gameDbName
	_logDbName = logDbName
}

func GetDatabaseName(dbType DbType) string {
	switch dbType {
	case LOGIN:
		return _loginDbName
	case GAME:
		return _gameDbName
	case LOG:
		return _logDbName
	default:
		return ""
	}
}

// Model is an interface that implements all expected helper functions
type Model interface {
	GetDatabaseName() string
	GetTableName() string
	GetInsertString() string
}

// GetOptionalStringVal Returns the sql format for an optional string
func GetOptionalStringVal(val *string) string {
	if val == nil {
		return "NULL"
	}
	return "N'%s'"
}

// GetOptionalDecVal Returns the sql format for an optional numeric value
func GetOptionalDecVal(val interface{}) string {
	if val == nil {
		return "NULL"
	}
	return "%d"
}

// GetOptionalBinaryVal Returns the sql format for an optional binary array
func GetOptionalBinaryVal(val *Binary) string {
	if val == nil {
		return "NULL"
	}
	return "N'%s'"
}