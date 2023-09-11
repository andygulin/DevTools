package format

type Format interface {
	Format(input string) (string, error)
	FormatFile(filename string) (string, error)
}
