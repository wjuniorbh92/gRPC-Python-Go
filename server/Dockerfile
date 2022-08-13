FROM golang:alpine as builder

ENV APP_HOME /go/src/grpc
WORKDIR "$APP_HOME"
COPY . .

RUN go mod download
RUN go mod verify
RUN go build -o grpc_server

FROM golang:alpine 

ENV APP_HOME /go/src/grpc
RUN mkdir -p "$APP_HOME"
WORKDIR "$APP_HOME"

COPY --from=builder "$APP_HOME"/grpc_server $APP_HOME
CMD ["./grpc_server"]
