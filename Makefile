build:
	docker build --target=production -t poc . --no-cache
up:
	docker-compose up -d --force-recreate
run:
	go build -o main && \
	./main
reup: build up
down:
	docker-compose down && \
	docker system prune -f