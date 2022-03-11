FROM ubuntu:focal as builder
WORKDIR /build
RUN apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y gcc golang ca-certificates librados-dev
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -ldflags -'-w -s' -o radosdsproxy

FROM ubuntu:focal
WORKDIR /app
EXPOSE 8080
RUN apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y librados2
COPY --from=builder /build/radosdsproxy .
ENTRYPOINT /app/radosdsproxy
