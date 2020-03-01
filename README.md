# Neoway Challenge
## Data manipulation and persistence in an RDBMS

This program was built with the purpose of extracting the data from .csv or .txt files using the Golang programming language and inserting them into a PostgreSQL database, ensuring data cleansing and also validation. It all is done using docker-compose.

### - Requirements

#### Docker
Docker must be installed. It's possible to download it here: [link](https://www.docker.com/products/docker-desktop).

#### Git
Git also must be installed. It's possible to download it here: [link](https://git-scm.com/downloads).

### - Git Clone

Clone the application inside a folder:
```console
git clone https://github.com/mateusolvr/neoway.git
```

### - Files

The files wished to be inserted into the database must be .txt or .csv and put inside the folder `./neoway/myfiles/`

### - Execution

Once inside the folder `neoway` execute the command:
```console
docker-compose up
```

After a few seconds, your data will be cleaned and validated inside your database.


### - Useful commands

Here are a few other useful commands if one judges necessary.

`docker-compose up -d`: Runs the docker-compose in the background.

`docker-compose images`: Check all docker images.

`docker-compose ps`: Check all docker containers.

`docker-compose stop`: Stop services.

`docker-compose rm`: Remove stopped containers.

`docker-compose down`: Remove all docker containers.



### Possiveis melhorias:
jogar dados brutos num schema anterior sem nenhum tratamento e outro schema com dados tratados e validados
