.PHONY: ent
ent:
	go generate ./infrastructure/ent

.PHONY: wire
wire:
	cd ../../cmd/payment/ && wire

.PHONY: run
run:
	cd  ../../cmd/payment/ && go run .