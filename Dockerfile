FROM golang:1.21.11 AS build

WORKDIR /app

COPY go.mod go.sum ./
COPY . .

RUN go mod download

RUN CGO_ENABLED=0 go build -o amquiz

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/amquiz .

EXPOSE 8402

CMD [ "./amquiz" ]