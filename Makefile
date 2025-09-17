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
	docker run -d -p 25566:25566 --name mcs --network=micro-net mcserver-service:latest

run-cmd:
	docker run -d -p 8082:8082 -e SERVICE_HOST=cmd-serv --name cmd-serv --network=micro-net cmd-service:latest

run-DB:
	docker run -d -p 8083:8083 -e SERVICE_HOST=db-serv --name db-serv --network=micro-net db-service:latest

run-temperature:
	docker run -d -p 8081:8081 -e SERVICE_HOST=temp-serv --name temp-serv --network=micro-net temp-service:latest

run-website:
	docker run -d -p 8080:8080 -e SERVICE_HOST=web-serv --name web-serv --network=micro-net website-service:latest

run-consul-dev-server:
	docker run -d -p 8500:8500 -p 8600:8600/udp --name=dev-consul --network=micro-net consul:1.15.4 agent -server -ui -node=server-1 -bootstrap-expect=1 -client=0.0.0.0

# Stops
#ill do it later
