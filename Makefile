build-linux:
	GOOS=linux go build -o main .

docker-build:
	docker build . -t vk-activity-checker

docker-run:
	docker run --publish 7070:8080 -t vk-activity-checker