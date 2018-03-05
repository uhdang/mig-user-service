build:
	protoc -I/usr/local/include -I. \
		--go_out=plugins=micro:. \
		proto/auth/auth.proto

	docker build -t mig-user-service .

rundb:
	docker run -d --name mig-postgres -p 5432:5432 postgres

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
