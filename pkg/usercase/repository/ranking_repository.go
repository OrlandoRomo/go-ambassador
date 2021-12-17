package repository

type RankingCache interface {
	GetRanking() (map[string]float64, error)
}
