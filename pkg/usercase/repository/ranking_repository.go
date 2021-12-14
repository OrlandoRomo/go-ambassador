package repository

type RankingRepository interface {
	GetRanking() (map[string]float64, error)
}
