FROM golang:1.16


RUN curl -fsSL https://raw.githubusercontent.com/pressly/goose/master/install.sh | sh


WORKDIR /api


COPY . /api
RUN cd /api; go mod download

RUN cd /api; make build

ENTRYPOINT ["/api/main"]
