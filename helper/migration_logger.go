package helper

type MigrationLogger struct {
}

func (log *MigrationLogger) Printf(format string, v ...any) {
	SuccessPuts("run", format, v...)
}

func (log *MigrationLogger) Verbose() bool {
	return false
}
