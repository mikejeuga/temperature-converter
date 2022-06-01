repo=$(shell basename "`pwd`")
gopher:
	@git init
	@go mod init github.com/mikejeuga/$(repo)
	@go mod tidy
	@touch .gitignore
	@go install github.com/matryer/moq@latest
	@go get github.com/jackc/pgx
	@go get golang-migrate/migrate
	@go get github.com/spf13/viper
	@go get github.com/alecthomas/assert/v2


t: test
test:
	@go test -v ./...

ut: unit-test
unit-test:
	@go test -v --tags=unit ./...

c: commit
commit:
	@git commit -am "$m"
	@git pull --rebase
	git push
