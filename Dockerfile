FROM opensuse/tumbleweed:latest

RUN zypper --non-interactive in git go

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -a -tags netgo -ldflags '-extldflags "-static"' -o echo-server cmd/echo-server/main.go

FROM opensuse/leap:latest

RUN zypper -n up && \
    zypper -n in shadow && \
    zypper -n clean

COPY --from=0 /build/echo-server /usr/local/bin

RUN groupadd -r echo-server && \
    useradd -r -g echo-server -s /sbin/nologin -c "Docker image user" echo-server

USER echo-server

EXPOSE 8080

CMD ["/usr/local/bin/echo-server"]
