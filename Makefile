#!make

up:
	@docker-compose up -d

up-attach:
	@docker-compose up

down:
	@docker-compose down

destroy:
	@docker-compose down
	@docker volume rm $(shell docker volume ls -q)