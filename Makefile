.PHONY: setup expenses test

all: expenses
# all: expenses go-swagger

# setup:
# 	yarn
# 	make expenses

devinit: main.go
	GO111MODULE=on go mod init github.com/zgoldy/expenses
	go mod tidy

expenses: main.go
	GO111MODULE=on go build github.com/zgoldy/expenses/cmd/expenses

# swaggerlocal: go-swagger redoc

# go-swagger:
# 	go build -o ./swagger -i github.com/go-swagger/go-swagger/cmd/swagger
# 	./swagger generate spec -o ./docs/swagger.json


test:
	go test -cover -count=1