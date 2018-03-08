build:
	protoc -I/usr/local/include -I. \
		--go_out=plugins=micro:. \
		proto/auth/auth.proto

	docker build --no-cache -t mig-user-service .

createdockernetwork:
	docker network create mig

rundb:
	docker run -d --name mig-postgres -p 5432:5432 postgres

runmicro:
	docker run -p 8080:8080 \
		-e MICRO_REGISTRY=mdns \
		microhq/micro api \
		--handler=rpc \
		--address=:8080 \
		--namespace=mig 


run:
	docker run --name mig-user-service \
		--net="host" \
		-p 50051 \
		-e DB_HOST=localhost \
		-e DB_PASS=password \
		-e DB_USER=postgres \
		-e MICRO_SERVER_ADDRESS=:50051 \
		-e MICRO_REGISTRY=mdns \
		mig-user-service

createuser:
	curl -XPOST -H 'Content-Type: application/json' \
    -d '{ "service": "shippy.auth", "method": "Auth.Create", "request": { "user": { "email": "ewan.valentine89@gmail.com", "password": "testing123", "name": "Ewan Valentine", "company": "BBC" } } }' \
    http://localhost:8080/rpc

authenticateuser:
	curl -XPOST -H 'Content-Type: application/json' \
    -d '{ "service": "shippy.auth", "method": "Auth.Auth", "request":  { "email": "your@email.com", "password": "SomePass" } }' \
    http://localhost:8080/rpc
