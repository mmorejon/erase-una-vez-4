# golang alpine 1.13.5-alpine
FROM golang:1.14.2-alpine AS builder
# Create workspace
WORKDIR /opt/app/
COPY go.mod .
# fetch dependancies
RUN go mod download
RUN go mod verify
# copy the source code as the last step
COPY . .
# build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o /go/bin/erase-una-vez-4 .

# build a small image
FROM alpine:3.11.6
LABEL language="golang"
# Create appuser.
RUN adduser -D -g '' elf
# copy the static executable
COPY --from=builder /go/bin/erase-una-vez-4 /go/bin/erase-una-vez-4
# use an unprivileged user.
USER elf
# run app
ENTRYPOINT ["/go/bin/erase-una-vez-4"]