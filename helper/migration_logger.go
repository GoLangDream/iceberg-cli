package helper

type MigrationLogger struct {
}

func (log *MigrationLogger) Printf(format string, v ...any) {
	SuccessPuts("migrate", format, v...)
}

func (log *MigrationLogger) Verbose() bool {
	return false
}
