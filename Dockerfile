FROM golang:1.17.7-alpine3.15
WORKDIR /code
COPY . /code/
WORKDIR /code/cmd
CMD ["go","run","main.go"]