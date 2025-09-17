# Builds
build-cmd-service:
	docker build -f Commands/dockerfile -t cmd-service .

build-DB-service:
	docker build -f DB/dockerfile -t db-service .

build-temperature-service:
	docker build -f Temperature/dockerfile -t temp-service .

build-website-service:
	docker build -f Website/dockerfile -t website-service .

build-mcserver-service:
	docker build -f MCServer/dockerfile -t mcserver-service ./MCServer

# Runs
run-mcserver:
	docker run -d -p 25565:25565 --name mcs mcserver-service:latest

run-cmd:
	docker run -d -p 8082:8082 --name cmd-serv cmd-service:latest

run-DB:
	docker run -d -p 8083:8083 --name db-serv db-service:latest

run-temperature:
	docker run -d -p 8081:8081--name temp-serv temp-service:latest

run-website:
	docker run -d -p 8080:8080 --name web-serv website-service:latest

run-consul-dev-server:
	docker run -d -p 8500:8500 -p 8600:8600/udp --name=dev-consul consul:1.15.4 agent -server -ui -node=server-1 -bootstrap-expect=1 -client=0.0.0.0

# Stops
#ill do it later
