FROM golang:1.17.6
WORKDIR /code
COPY . /code/
WORKDIR /code/cmd
CMD ["go","run","main.go"]