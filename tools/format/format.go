package format

type Format interface {
	Format(input string) (string, error)
}
