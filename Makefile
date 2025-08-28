local-deps:
	kubectl apply -f k8s/local/postgres-adminer.yaml -n nba-stats

clean-local-deps:
	kubectl delete -f k8s/local/postgres-adminer.yaml -n nba-stats

daily-box-scores:
	BOX_SCORES_OUTPUT_FILENAME=test/sportsblaze/daily-box-scores-2024-10-24.json go run ./cmd/gamestatsjob/main.go