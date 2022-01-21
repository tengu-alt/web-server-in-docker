# web-server-in-docker
## a simple database in your container
*this service requires **docker & docker-compose***

*please check out **docker-compose.yml** where are described enviroment variables,
which are used in your database(use it to connect to the database)* 

### for start-up of the database and web service run your docker and type the following command in a root directory of the project, which you cloned
```docker-compose up --build```
### after this you will see the following message
>Starting server for testing HTTP POST in 8081...
### the next step is allowing the access to your firewall, the web-server will be available on your local port
**localhost:8081** / **127.0.0.1:8081**