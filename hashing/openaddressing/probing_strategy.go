package openaddressing

type ProbingStrategy interface {
	Hash(data string) int
}
