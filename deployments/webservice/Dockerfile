FROM golang:1.16.3-alpine as builder

# create a working directory
WORKDIR /go/src/app

# add source code
ADD ../../build/deploy .

# build the source
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o susswebservice main.go

# use a minimal alpine image
FROM alpine:3.7

# set working directory
WORKDIR /root
# copy the binary from builder
COPY --from=builder /go/src/app/susswebservice .
ENTRYPOINT ["./susswebservice"]
