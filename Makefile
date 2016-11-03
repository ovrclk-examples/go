NAME ?= go

BINS := $(NAME) $(NAME)-linux-amd64

$(NAME): main.go
	go build -v -o $(NAME)

$(NAME)-linux-amd64: main.go
	GOOS=linux GOARCH=amd64 go build -v -o $(NAME)-linux-amd64

image: $(NAME)-linux-amd64
	docker build -t $(NAME) . --build-arg app=$(NAME)-linux-amd64

clean:
	rm $(BINS) 2>/dev/null || true

.PHONY: clean
