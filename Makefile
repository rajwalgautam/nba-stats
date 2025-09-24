export POSTGRES_USER=nbastatsuser
export POSTGRES_PASSWORD=nbastatspassword
export POSTGRES_DB=nba_stats
export POSTGRES_PORT=30001
export POSTGRES_HOST=localhost

export SPORTS_API_KEY=your_api_key_here
# export BOX_SCORES_OUTPUT_FILENAME=output.json

local-deps:
	kubectl apply -R -f k8s/local/ -n nba-stats

clean-local-deps:
	kubectl delete -R -f k8s/local/ -n nba-stats

run-gamestatsjob:
	go run ./cmd/gamestatsjob/main.go
