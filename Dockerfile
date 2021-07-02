FROM golang:1.15-alpine

# Install psql
RUN apk add --no-cache postgresql gcc g++

# Install golangci-lint
RUN wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.28.3

ENV GO111MODULE=on
# Set work directory
WORKDIR /app

ADD . /app

# Install goreman
RUN go get "github.com/mattn/goreman"

CMD [ "go", "test", "-count=1", "./..." ]
