include .env
export $(shell sed 's/=.*//' .env)

.SILENT:

echo-config:
	echo POSTGRES_USER: ${POSTGRES_USER}
	echo POSTGRES_PASSWORD: ${POSTGRES_PASSWORD}
	echo POSTGRES_DB: ${POSTGRES_DB}
	echo POSTGRES_HOST: ${POSTGRES_HOST}
	echo POSTGRES_PORT: ${POSTGRES_PORT}
	echo POSTGRES_HOST_DOCKER: ${POSTGRES_HOST_DOCKER}
	echo SPORTSBLAZE_API_KEY: ${SPORTSBLAZE_API_KEY}

local-deps:
	kubectl apply -R -f k8s/local/ -n nba-stats

clean-local-deps:
	kubectl delete -R -f k8s/local/ -n nba-stats

run-gamestatsjob:
	go run ./cmd/gamestatsjob/main.go

docker-build-gamestatsjob:
	docker build --no-cache -f build/gamestatsjob.Dockerfile . -t nba-stats --debug

# If running in Docker on Mac, use host.docker.internal to connect to host machine
docker-run-gamestatsjob:
	docker run -e POSTGRES_USER=${POSTGRES_USER} -e POSTGRES_PASSWORD=${POSTGRES_PASSWORD} -e POSTGRES_HOST=${POSTGRES_HOST_DOCKER} -e POSTGRES_PORT=${POSTGRES_PORT} -e POSTGRES_DB='${POSTGRES_DB}' -e SPORTSBLAZE_API_KEY=${SPORTSBLAZE_API_KEY} nba-stats