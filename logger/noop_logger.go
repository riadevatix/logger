package logger

// noOpLogger represents a dummy logger
type noOpLogger struct{}

// Logf does not do anything here
func (np *noOpLogger) Logf(string, ...interface{}) {}

// Logln does not do anything here
func (np *noOpLogger) Logln(...interface{}) {}

// NoOpLogger returns a NoOpLogger wrapped into ILogger interface
func NoOpLogger() ILogger {
	return &noOpLogger{}
}
