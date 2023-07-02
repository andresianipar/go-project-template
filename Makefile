# Usage:
# make run		# compiles and runs the named main Go package
# make run-dev	# compiles and runs the named main Go package with live reload capability
# ...

.DEFAULT_GOAL := run

project_name := nuke-deployer-system-app
image_name := $(project_name):latest

run:
	@go run cmd/main.go

run-dev:
	@air

tidy:
	@go mod tidy

clean:
	@go clean -modcache

lint:
	@go fmt ./...

test:
	@CGO_ENABLED=1 go test -race ./tests

swag:
	@swag init -g router.go -d ./internal/router -o ./api
	@swag fmt

build-up:
	@make build-up-bg
	@make shell

build-up-bg:
	@docker build -t $(image_name) .
	@make remove-container
	@docker run -d -p 52603:52603 --name $(project_name) $(image_name)

compose-up:
	@make compose-up-bg
	@make shell

compose-up-bg:
	@make remove-container
	@docker-compose -f docker-compose-development.yaml up --build -d

remove-container:
	@chmod 700 ./scripts/remove-container.sh
	@./scripts/remove-container.sh $(project_name)

shell:
	@docker ps -a | grep $(project_name) | awk '{ print $1 }' | xargs -I % sh -c 'docker exec -it % /bin/sh'

start:
	@docker ps -a | grep $(project_name) | awk '{ print $1 }' | xargs -I % sh -c 'docker start %'

stop:
	@docker ps -a | grep $(project_name) | awk '{ print $1 }' | xargs -I % sh -c 'docker stop %'

convert-compose-dev:
	@ls deploy/ | wc -l | xargs -I % sh -c '[ % -gt 0 ] && rm deploy/*'
	@kompose convert -f docker-compose-development.yaml -o ./deploy

convert-compose-prod:
	@ls deploy/ | wc -l | xargs -I % sh -c '[ % -gt 0 ] && rm deploy/*'
	@kompose convert -f docker-compose-production.yaml -o ./deploy
