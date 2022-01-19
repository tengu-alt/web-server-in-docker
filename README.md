# web-server-in-docker
## a simply database in your container
*this service require **docker & docker-compose***

*please check out **docker-compose.yml** where described enviroment variables
which are used in your database(use it for connect to database)* 

### for startup the database run your docker and type the following command inside of root directory of project which you clone
```docker-compose up --build```
### and for startup the web service type next command inside of ./cmd/ directory
```go run .```
### after this you will see the followinng message
>Starting server for testing HTTP POST in 8081...
### next step is allow the access to your firewall web-server will be available on your local port
### to use it go to your port by using localhost IP not "localhost"
~~localhost:8081~~ **127.0.0.1:8081**