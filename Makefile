NAME ?= go

BINNAME := go

BINS := $(BINNAME) $(BINNAME)-linux-amd64

$(BINNAME): main.go
	go build -v -o $(BINNAME)

$(BINNAME)-linux-amd64: main.go
	GOOS=linux GOARCH=amd64 go build -v -o $(BINNAME)-linux-amd64

image: $(BINNAME)-linux-amd64
	docker build -t $(NAME) . --build-arg binname=$(BINNAME)-linux-amd64

publish: image
	docker tag $(NAME) $(REPO)/$(NAME)
	docker push $(REPO)/$(NAME)

clean:
	rm $(BINS) 2>/dev/null || true

.PHONY: clean
