FROM golang as builder
WORKDIR /go/src/github.com/lgsvl/data-marketplace-chaincode-rest
RUN go get github.com/Masterminds/glide
COPY . .

RUN glide up --strip-vendor

RUN go get github.com/rakyll/statik
RUN statik -src swaggerui
RUN CGO_ENABLED=1 GOOS=linux go build -tags netgo -v -a --ldflags '-w -linkmode external -extldflags "-static"' -installsuffix cgo -o bin/chaincode-rest main.go


FROM alpine:3.7
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/lgsvl/data-marketplace-chaincode-rest/bin/chaincode-rest .
COPY --from=hyperledger/fabric-peer:1.3.0 /usr/local/bin/peer /usr/local/bin/
COPY --from=hyperledger/fabric-peer:1.3.0 /etc/hyperledger/fabric  /etc/hyperledger/fabric

ENV PATH="/usr/local/bin/:${PATH}"

CMD ["./chaincode-rest"]

EXPOSE 9090