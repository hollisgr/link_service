SRC := cmd/app/main.go
EXEC := link_service

CLEANENV := github.com/ilyakaznacheev/cleanenv
GIN := github.com/gin-gonic/gin

all: build run

build: clean
	go build -o $(EXEC) $(SRC)

run:
	./$(EXEC)

clean:
	rm -f $(EXEC)

mod:
	go mod init $(EXEC)

get:
	go get $(GIN) \
		$(CLEANENV) \
		$(LOGRUS) 