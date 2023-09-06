SHELL := /bin/bash

## Golang Stuff
GOCMD=go
GORUN=$(GOCMD) run
ENV=local
GOPRIVATE=github.com/waresix/*

SERVICE=project

init:
	$(GOCMD) mod init $(SERVICE)

tidy:
	ENV=local GOPRIVATE=$(GOPRIVATE) $(GOCMD) mod tidy

test:
	$(GOCMD) test ./...

test-cov:
	$(GOCMD) test ph-service/domain/phunter/unit_test -coverpkg ph-service/domain/phunter/usecase -coverprofile cover.out --v -p 1 ./...
	$(GOCMD) tool cover -func cover.out

coverage-report:
	$(GOCMD) tool cover -html cover.out

run:
	echo "for local development, please run: make run ENV=local"
	ENV=$(ENV) $(GORUN) main.go --mode grpc

run-jaeger:
	ENV=$(ENV) $(GORUN) main.go --tracer jaeger

run-consumer:
	ENV=$(ENV) $(GORUN) main.go --mode consumer

run-cron-sample-command:
	ENV=$(ENV) $(GORUN) main.go --mode cron --name sample-command

run-cron-update-price-hunter-status-expired:
	ENV=$(ENV) $(GORUN) main.go --mode cron --name update-price-hunter-status-expired

run-cron-create-price-hunter-approval-history-optional-approver:
	ENV=$(ENV) $(GORUN) main.go --mode cron --name create-price-hunter-approval-history-optional-approver

run-cron-recalculate-expired-pricebook-vendor:
	ENV=$(ENV) $(GORUN) main.go --mode cron --name recalculate-expired-pricebook-vendor

migration-up:
	ENV=$(ENV) $(GORUN) main.go --migrate-up

migration-down:
	ENV=$(ENV) $(GORUN) main.go --migrate-down

# Proto Generator
proto-gen:
	protoc --proto_path=proto --go_out=paths=source_relative,plugins=grpc:./pb proto/*/*.proto

proto-gen2:
	protoc proto/*/*.proto --go_out=pb --go_opt=paths=source_relative --go-grpc_out=require_unimplemented_servers=false:pb --go-grpc_opt=paths=source_relative -I=proto --experimental_allow_proto3_optional

new-proto-gen:
	protoc proto/phunter/*.proto --go_out=pb --go_opt=paths=source_relative --go-grpc_out=require_unimplemented_servers=false:pb --go-grpc_opt=paths=source_relative -I=proto
	protoc proto/quotation/*.proto --go_out=pb --go_opt=paths=source_relative --go-grpc_out=require_unimplemented_servers=false:pb --go-grpc_opt=paths=source_relative -I=proto
	protoc proto/quotation_line_item/*.proto --go_out=pb --go_opt=paths=source_relative --go-grpc_out=require_unimplemented_servers=false:pb --go-grpc_opt=paths=source_relative -I=proto
	protoc proto/location/*.proto --go_out=pb --go_opt=paths=source_relative --go-grpc_out=require_unimplemented_servers=false:pb --go-grpc_opt=paths=source_relative -I=proto
	protoc proto/fleet/*.proto --go_out=pb --go_opt=paths=source_relative --go-grpc_out=require_unimplemented_servers=false:pb --go-grpc_opt=paths=source_relative -I=proto

# Mock Generator
mock-gen:
	mockery --dir  domain/phunter --name PriceHunterRepoInterface --filename PriceHunterRepoInterfaceMock.go --output domain/phunter/mocks
	mockery --dir  domain/phunter --name PriceHunterUsecaseInterface --filename PriceHunterUsecaseInterfaceMock.go --output domain/phunter/mocks
	
mock-gen2:
	mockery --dir  domain/phunter --name PriceHunterRepoInterface --output domain/phunter/mocks
	mockery --dir  domain/phunter --name PriceHunterUsecaseInterface --output domain/phunter/mocks
	mockery --dir  pb/quotation --name QuotationServiceClient --output domain/phunter/mocks
	mockery --dir  pb/quotation_line_item --name QuotationLineItemServiceClient --output domain/phunter/mocks
	mockery --dir  pb/company --name CompanyServiceClient --output domain/phunter/mocks
	mockery --dir  lib/external_lib --name ExternalLibInterface --output domain/phunter/mocks
	mockery --dir  pb/vquotation --name VendorQuotationServiceClient --output domain/phunter/mocks
	mockery --dir  pb/quotation_add_cost --name QuotationAddCostServiceClient --output domain/phunter/mocks
	mockery --dir  pb/auth --name UserMgmtServiceClient --output domain/phunter/mocks
	mockery --dir  pb/user_access --name UserAccessServiceClient --output domain/phunter/mocks
	mockery --dir  pb/fleet --name FleetServiceClient --output domain/phunter/mocks
	mockery --dir  pb/location --name LocationServiceClient --output domain/phunter/mocks
	mockery --dir  pb/notification --name NotificationServiceClient --output domain/phunter/mocks
	mockery --dir  pb/quotation_log --name QuotationLogServiceClient --output domain/phunter/mocks
	mockery --dir  pb/addcost --name AddCostServiceClient --output domain/phunter/mocks
	mockery --dir  pb/map_data_matrix --name MapDataMatrixServiceClient --output domain/phunter/mocks
	mockery --dir  pb/currency --name CurrencyServiceClient --output domain/phunter/mocks
	mockery --dir  pb/location --name LocationServiceClient --output domain/phunter/mocks
	mockery --dir  pb/type_agreement --name TypeAgreementsServiceClient --output domain/phunter/mocks