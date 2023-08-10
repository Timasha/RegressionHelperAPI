FROM golang:alpine

WORKDIR /app

COPY . .

RUN ls

RUN go mod download

RUN go build cmd/regression/main.go

EXPOSE 8080

CMD [ "./main" ]