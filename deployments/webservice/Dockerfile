FROM golang:1.16.3-alpine as builder
RUN apk add build-base
# create a working directory
WORKDIR /go/src/app

# add source code
ADD . .

# build the source
#RUN go build -a -installsuffix cgo -o susswebservice cmd/webservice/main.go
RUN go build -o susswebservice cmd/webservice/main.go

# use a minimal alpine image
FROM alpine:3.7

# set working directory
WORKDIR /root
# copy the binary from builder
COPY --from=builder /go/src/app/susswebservice .
COPY --from=builder /go/src/app/configs/. configs/.
COPY --from=builder /go/src/app/openapi/. openapi/.
COPY --from=builder /go/src/app/static/. static/.
COPY --from=builder /go/src/app/geolitedb/. geolitedb/.
COPY --from=builder /go/src/app/db/. db/.
COPY --from=builder /go/src/app/certs/. certs/.

CMD ["./susswebservice"]

