#
# daring-chupacabra-backend
#
# docker build . -t daring-chupacabra-be:1
#
# docker run --rm -it daring-chupacabra-be:1 /bin/sh
# docker run --rm daring-chupacabra-be:1
#
FROM golang:1.16-alpine
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
#
RUN go build -o /app/manager

# This is for documentation purposes only.
# To actually open the port, runtime parameters
# must be supplied to the docker command.
#EXPOSE 8080

# (Optional) environment variable that our dockerised
# application can make use of. The value of environment
# variables can also be set via parameters supplied
# to the docker command on the command line.
#ENV HTTP_PORT=8081

# Run
CMD [ "/app/manager" ]