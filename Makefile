local-deps:
	kubectl apply -f k8s/local/ -n nba-stats

clean-local-deps:
	kubectl delete -f k8s/local/ -n nba-stats