ENV=${mentor_env}

all: build_user_system build_case_system build_chatroom build_classroom build_finance_system

build_user_system: 
	sudo docker build --build-arg mentor_env=$(ENV) --network=mentor_server_backend-bridge -t mentor_user_system ../user_system
	sudo docker push mentor_user_system:latest 35.201.214.194:5000/mentor_user_system
	sudo docker image prune --filter="dangling=true" -f

deploy_evaluation_system:
	sudo docker stack deploy --compose-file evaluation-system-docker-compose.yml evaluation_system_stack