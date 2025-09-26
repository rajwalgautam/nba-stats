export POSTGRES_USER=nbastatsuser
export POSTGRES_PASSWORD=nbastatspassword
export POSTGRES_DB=nba_stats
export POSTGRES_PORT=30001
export POSTGRES_HOST=localhost

# If running in docker, use this host to connect to host machine
export POSTGRES_HOST_DOCKER=host.docker.internal

export SPORTSBLAZE_API_KEY=your_api_key_here
# export BOX_SCORES_OUTPUT_FILENAME=output.json

local-deps:
	kubectl apply -R -f k8s/local/ -n nba-stats

clean-local-deps:
	kubectl delete -R -f k8s/local/ -n nba-stats

run-gamestatsjob:
	go run ./cmd/gamestatsjob/main.go

docker-build-gamestatsjob:
	docker build --no-cache -f build/gamestatsjob.Dockerfile . -t nba-stats --debug

docker-run-gamestatsjob:
	docker run -e POSTGRES_USER=${POSTGRES_USER} -e POSTGRES_PASSWORD=${POSTGRES_PASSWORD} -e POSTGRES_HOST=${POSTGRES_HOST_DOCKER} -e POSTGRES_PORT=${POSTGRES_PORT} -e POSTGRES_DB='${POSTGRES_DB}' -e SPORTSBLAZE_API_KEY=${SPORTSBLAZE_API_KEY} nba-stats