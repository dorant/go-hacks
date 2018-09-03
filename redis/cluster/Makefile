NETWORKNAME=rediscluster
NAME=rediscluster-instance

# https://medium.com/commencis/creating-redis-cluster-using-docker-67f65545796d

start:
	docker network create $(NETWORKNAME) && \
	\
	for INSTANCE in `seq 1 6`; do \
		docker run -d \
		-p 700$$INSTANCE:6379 \
		--rm --name $(NAME)-$$INSTANCE \
		--net $(NETWORKNAME) \
		-v $$PWD/redis.conf:/usr/local/etc/redis/redis.conf \
		redis:latest \
		redis-server /usr/local/etc/redis/redis.conf ; \
	done && \
	\
	echo 'yes' | docker run -i --rm --net $(NETWORKNAME) ruby sh -c '\
	gem install redis && \
	wget http://download.redis.io/redis-stable/src/redis-trib.rb && \
	ruby redis-trib.rb create --replicas 1 \
	172.19.0.2:6379 172.19.0.3:6379 \
	172.19.0.4:6379 172.19.0.5:6379 \
	172.19.0.6:6379 172.19.0.7:6379'

stop:
	for INSTANCE in `seq 1 6`; do \
		docker stop $(NAME)-$$INSTANCE ; \
	done && \
	docker network rm $(NETWORKNAME)

status:
	docker exec $(NAME)-1 redis-cli cluster nodes

benchmark:
	docker exec $(NAME)-1 redis-benchmark -h 172.19.0.3 -q -t set,get,incr,lpush,lpop,sadd,spop,lpush,lrange -c 100 -p 6379


run:
	go run main.go
run-cli:
	docker exec $(NAME)-1 redis-cli get key
# -c : Enable cluster mode (follow -ASK and -MOVED redirections).
run-cluster-cli:
	docker exec $(NAME)-1 redis-cli -c get key

