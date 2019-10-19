FROM golang:alpine AS builder
LABEL maintainer="Niraj Georgian <nirajgeorgian@gmail.com>"

# define system module and update them
RUN apk update && apk add --no-cache \
    git \
    dep

# Create user for accessing job service.
RUN adduser -D -g '' accountuser
WORKDIR /go/src/github.com/nirajgeorgian/account

# install dependencies
COPY Gopkg.* ./
RUN dep ensure -vendor-only -v
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main main.go

# use scratch (base for a docker image)
FROM scratch
COPY --from=builder /etc/passwd /etc/passwd

# set working directory
WORKDIR /root/
COPY --from=builder /go/src/github.com/nirajgeorgian/account/config.yaml .
COPY --from=builder /go/src/github.com/nirajgeorgian/account/main .

# Use an unprivileged user.
USER accountuser
EXPOSE 3000
CMD ["./main", "serve"]