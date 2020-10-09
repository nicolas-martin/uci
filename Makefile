BINARY:=uci
BUILDOPTS:=-v -ldflags="-s -w"
GOPATH?=$(HOME)/go


.PHONY: build
router:
	GOOS=linux GOARCH=mipsle GOMIPS=softfloat go build $(BUILDOPTS) -tags $@ -o $(BINARY).mipsle

.PHONY: run
run: build
	scp -i ~/.ssh/ph -o StrictHostKeyChecking=no ./$(BINARY).mipsle root@192.168.177.1:/tmp/
