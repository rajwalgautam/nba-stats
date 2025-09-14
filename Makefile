local-deps:
	kubectl apply -f k8s/local/postgres-nba.yml -n nba-stats
	kubectl apply -f k8s/local/adminer-nba.yml -n nba-stats

clean-local-deps:
	kubectl delete -f k8s/local/postgres-nba.yml -n nba-stats
	kubectl delete -f k8s/local/adminer-nba.yml -n nba-stats