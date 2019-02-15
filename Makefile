ARG = foo

print:
	go run main.go printchain

add:
	go run main.go addblock -data ${ARG}
