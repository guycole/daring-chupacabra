#
# daring-chupacabra
#
# docker build . -t daring-chupacabra:1
#
FROM golang:1.19-alpine3.16
LABEL build_date="2022-10-12"
LABEL description="daring-chupacabra"
LABEL maintainer="guycole@gmail.com"
#
WORKDIR /app
#
COPY go.mod .
COPY go.sum .
RUN go mod download
#
COPY *.go ./
COPY *.html ./
#
RUN go build -o /app/manager

# Run
CMD [ "/app/manager" ]