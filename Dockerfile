# syntax=docker/dockerfile:1

FROM golang:1.19-alpine as builder

WORKDIR /go/src/github.com/ablancas22/messenger-frontend
ENV GO111MODULE on
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go install github.com/42wim/go.rice && \
    go install github.com/42wim/go.rice/rice && \
    go install github.com/ablancas22/messenger-frontend && \
    cd /go/src/github.com/ablancas22/messenger-frontend/views && \
    rice embed-go && \
    cd .. && \
    go generate 'rice embed-go' && \
    go build

# Multi-Stage production build
FROM alpine AS production
RUN apk --no-cache add ca-certificates

WORKDIR /github.com/ablancas22/messenger-frontend
# Retrieve the binary from the previous stage
COPY --from=builder /go/src/github.com/ablancas22/messenger-frontend /github.com/ablancas22/messenger-frontend
EXPOSE 8080

CMD [ "./messenger-frontend" ]