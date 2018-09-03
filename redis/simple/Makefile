NAME=redis-test

start:
	docker run -d -p 6379:6379 \
	--rm --name $(NAME) \
	redis:latest

stop:
	docker stop $(NAME)

logs:
	docker logs $(NAME)

run:
	go run main.go
