all: compile

target      ?=  pulumi-import-state

compile: # locally
	@echo "Compiling..."
	go build 	-o $(target) cmd/*.go
