local-deps:
	kubectl apply -R -f k8s/local/ -n nba-stats

clean-local-deps:
	kubectl delete -R -f k8s/local/ -n nba-stats