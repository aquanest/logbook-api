FROM golang:1.17-alpine

ARG LOGBOOK_ATMOS_API
ARG LOGBOOK_ATMOS_TOKEN

ENV LOGBOOK_ATMOS_API=$LOGBOOK_ATMOS_API
ENV LOGBOOK_ATMOS_TOKEN=$LOGBOOK_ATMOS_TOKEN

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY / ./
RUN CGO_ENABLED=0 go build -trimpath -o ./logbook ./cmd/main.go

EXPOSE 8080

CMD [ "./logbook" ]

