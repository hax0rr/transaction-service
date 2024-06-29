FROM golang:1.20

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build -v -o ./out .
RUN chmod +x ./out

RUN go test ./...

EXPOSE 8080
CMD ./out migrate ; ./out app