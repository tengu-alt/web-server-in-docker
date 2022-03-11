FROM golang:1.17.7-alpine3.14 as build
WORKDIR /code
COPY . /code/
WORKDIR /code/cmd
RUN go build -o /registration main.go

#############
FROM alpine
WORKDIR /code
COPY --from=build /registration /bin/registration
COPY --from=build /code/assets /assets

ENTRYPOINT ["/bin/registration"]

