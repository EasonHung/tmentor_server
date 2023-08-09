ENV=${mentor_env}

all: build_user_system build_case_system build_chatroom build_classroom build_finance_system

build_user_system: 
	export mentor_env=develop
	sudo docker build --build-arg mentor_env=$(ENV) --network=mentor_server_backend-bridge -t mentor_user_system ./user_system
	sudo docker image prune --filter="dangling=true" -f
	sudo docker stop mentor_user_system || true
	sudo docker run --rm -p 2306:2306 -d --network=mentor_server_backend-bridge --name=mentor_user_system mentor_user_system

build_case_system: 
	export mentor_env=develop
	sudo docker build --build-arg mentor_env=$(ENV) --network=mentor_server_backend-bridge -t mentor_case_system ./case_system
	sudo docker image prune --filter="dangling=true" -f
	sudo docker stop mentor_case_system || true
	sudo docker run --rm -p 2305:2305 -d --network=mentor_server_backend-bridge --name=mentor_case_system mentor_case_system

build_chatroom: 
	export mentor_env=develop
	sudo docker build --build-arg mentor_env=$(ENV) --network=mentor_server_backend-bridge -t mentor_chatroom ./chatroom
	sudo docker image prune --filter="dangling=true" -f
	sudo docker stop mentor_chatroom || true
	sudo docker run --rm -p 2303:2303 -d --network=mentor_server_backend-bridge --name=mentor_chatroom mentor_chatroom

build_classroom: 
	export mentor_env=develop
	sudo docker build --build-arg mentor_env=$(ENV) --network=mentor_server_backend-bridge -t mentor_classroom ./classroom
	sudo docker image prune --filter="dangling=true" -f
	sudo docker stop mentor_classroom || true
	sudo docker run --rm -p 2304:2304 -d --network=mentor_server_backend-bridge --name=mentor_classroom mentor_classroom

build_finance_system: 
	export mentor_env=develop
	sudo docker build --build-arg mentor_env=$(ENV) --network=mentor_server_backend-bridge -t mentor_finance_system ./finance_system
	sudo docker image prune --filter="dangling=true" -f
	sudo docker stop mentor_finance_system || true
	sudo docker run --rm -p 2307:2307 -d --network=mentor_server_backend-bridge --name=mentor_finance_system mentor_finance_system

build_evaluation_system: 
	export mentor_env=develop
	sudo docker build --build-arg mentor_env=$(ENV) --network=mentor_server_backend-bridge -t mentor_evaluation_system ./evaluation_system
	sudo docker image prune --filter="dangling=true" -f
	sudo docker stop mentor_evaluation_system || true
	sudo docker run --rm -p 2308:2308 -d --network=mentor_server_backend-bridge --name=mentor_evaluation_system mentor_evaluation_system

build_redis:
	sudo docker run -d -p 6379:6379 --network=mentor_server_backend-bridge --name=redis -e ALLOW_EMPTY_PASSWORD=yes -e DISABLE_COMMANDS=FLUSHDB,FLUSHALL,CONFIG redis