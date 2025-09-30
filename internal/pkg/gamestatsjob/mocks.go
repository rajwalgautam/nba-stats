package gamestatsjob

import "github.com/rajwalgautam/nba-stats/internal/pkg/sportsblaze"

type mockStorage struct {
	saveErr error
}

func (ms mockStorage) SaveTeam(sportsblaze.Team) error {
	return ms.saveErr
}

type mockStatsClient struct {
	dailyBoxScoreResp *sportsblaze.DailyBoxScores
	dailyBoxScoreErr  error
}

func (msc mockStatsClient) DailyBoxScores(date string) (*sportsblaze.DailyBoxScores, error) {
	return msc.dailyBoxScoreResp, msc.dailyBoxScoreErr
}
