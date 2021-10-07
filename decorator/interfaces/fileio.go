package interfaces

type (
	IOWriter interface {
		WriteFile(string, string) error
	}
	IOReader interface {
		ReadFile(string) (string, error)
	}
)
