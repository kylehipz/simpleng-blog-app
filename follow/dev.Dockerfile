FROM golang:1.22

WORKDIR /app

COPY go.mod go.sum ./
RUN go install github.com/cosmtrek/air
RUN go mod download && go mod verify

COPY . .

CMD ["air"]
