FROM golang:1.18

WORKDIR /usr/src/app

COPY go.* ./
RUN go mod download && go mod verify

COPY . .

CMD ["go", "test"]
