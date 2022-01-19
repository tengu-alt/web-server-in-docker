FROM golang:1.17.6
WORKDIR /code
COPY . /code/
CMD ["go","run","/code/cmd/"]