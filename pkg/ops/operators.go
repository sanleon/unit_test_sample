package ops

type Operator interface {
	Generate(int, int) string
	Degenerate(string)(int, int, error)
}
