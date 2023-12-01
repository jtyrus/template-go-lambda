build:
	go build -C src -o ../bin/portfolio lambda_handler/handler.go

compile:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -C src -ldflags '-s -w' -o ../bin/portfolio_x86 lambda_handler/handler.go
	GOOS=linux GOARCH=arm64 go build -C src -o ../bin/portfolio_arm64 lambda_handler/handler.go
	GOOS=linux GOARCH=arm go build -C src -o ../bin/portfolio_arm lambda_handler/handler.go
