ARG = foo
FROM = from
TO = to
AMOUNT = 0
ADDRESS = address

print:
	go run main.go printchain

send:
	go run main.go send --from ${FROM} --to ${TO} --amount ${AMOUNT}

getbalance:
	go run main.go getbalance --address ${ADDRESS}

build:
	go build main.go && rm main

clean:
	rm -rf uooo.db
