FROM golang:1.14 as builder

COPY . .

RUN go build -o /go/bin/server.go main.go

# ---

FROM registry.access.redhat.com/ubi8:latest

LABEL io.k8s.display-name="Echo Server"
LABEL name="Echo Server"
LABEL vendor="Engeto"
LABEL version="v0.0.0"
LABEL release="N/A"
LABEL summary="Kubernetes Example Application: Echo Server"
LABEL description="See summary"

COPY --from=builder /go/bin/server.go /opt/app-root/bin/server.go

ENTRYPOINT [ "/opt/app-root/bin/server.go" ]

