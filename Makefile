repo=$(shell basename "`pwd`")
gopher:
	@git init
	@go mod init github.com/mikejeuga/$(repo)
	@go mod tidy
	@touch .gitignore
	@go get -u github.com/gorilla/mux
	@go install github.com/matryer/moq@latest
	@go install github.com/alecthomas/assert/v2
	@go get github.com/jackc/pgx
	@go get golang-migrate/migrate
	@go get github.com/spf13/viper


t: test
test:
	@go test -v ./...


ut: unit-test
unit-test:
	@go test -v --tags=unit ./...

at: acceptance-test
acceptance-test:
	@docker-compose -f docker-compose.yml up -d
	@go test -v --tags=acceptance ./...
	@docker-compose down

c: commit
commit:
	@git commit -am "$m"
	@git pull --rebase
	git push
