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

### - Validation

To be able to validate, one can use the following commands.

Run the following to get the PostgreSQL container.
```console
docker-compose ps
```
Substitute the container name in the next code.
```console
docker exec -it $container psql -U postgres -d neoway
```

Now one is able to check the data. For example:
```console
SELECT * FROM public.analise_compra_usuario;
```

### - Useful commands

Here are a few other useful commands if one judges necessary.

`docker-compose up -d`: Runs the docker-compose in the background.

`docker-compose images`: Check all docker images.

`docker-compose ps`: Check all docker containers.

`docker-compose stop`: Stop services.

`docker-compose rm`: Remove stopped containers.

`docker-compose down`: Remove all docker containers.



### - Improvements to be made:

- Delete the file which has already been inserted, so when one runs `docker-compose up` again the same data is not duplicated.

- Insert all raw data into a previous schema without any treatment or validation to have all the records. Use another schema with the clean and validated data.
