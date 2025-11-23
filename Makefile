run:
	go run ./cmd/sim/main.go


build:
	go build -o juliaset ./cmd/sim/main.go

clean:
	rm -f juliaset
