local-deps:
	kubectl apply -f k8s/local/postgres-adminer.yaml

clean-local-deps:
	kubectl delete -f k8s/local/postgres-adminer.yaml 