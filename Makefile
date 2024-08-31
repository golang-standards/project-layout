# note: call scripts from /scripts
.PHONY: vendor
vendor:
	go mod tidy && go mod vendor

.PHONY: run
run:
	docker-compose build
	docker-compose up -d