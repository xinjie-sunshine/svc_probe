FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /build/zero

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
RUN go build -ldflags="-s -w" -o /app/Svc_Probe main.go

FROM alpine

# RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/Svc_Probe /app/Svc_Probe

CMD ["./Svc_Probe"]