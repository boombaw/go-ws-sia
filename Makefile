appname := go-ws-sia

run:
	@go run cmd/server/main.go

build:
	@cd cmd/server && go build -o ../../bin/$(appname) && chmod +x ../../bin/$(appname)

exec:
	@./bin/$(appname)

startapp: build exec


start:
	@supervisorctl start $(appname)

stop:
	@supervisorctl stop $(appname)

createlog:
	@touch /var/log/$(appname).log

log:
	@tail -f /var/log/$(appname).log

bin:
	@mkdir bin

rmbin:
	@rm -rf bin

deploy: bin stop build start log

redeploy: stop build start log