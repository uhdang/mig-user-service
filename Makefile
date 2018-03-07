build:
	protoc -I/usr/local/include -I. \
		--go_out=plugins=micro:. \
		proto/auth/auth.proto

	docker build --no-cache -t mig-user-service .

createdockernetwork:
	docker network create mig

rundb:
	docker run -d --name mig-postgres --net mig -p 5432:5432 postgres

run:
	docker run --name mig-user-service \
		--net mig \
		-p 50051:50051 \
		-e DB_HOST=mig-postgres \
		-e DB_PASS=password \
		-e DB_USER=postgres \
		-e MICRO_SERVER_ADDRESS=:50051 \
		-e MICRO_REGISTRY=mdns \
		mig-user-service
