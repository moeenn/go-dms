## Delivery Management System


### Running postgres
Create file `docker-compose.yml` and add the following content.

```yml
version: '3.3'

services:
  db:
    container_name: "postgres-db"
    image: postgres:15.3-alpine
    restart: always
    environment:
      - POSTGRES_USER=devuser
      - POSTGRES_PASSWORD=devpass
      - POSTGRES_DB=dev
    volumes:
      - ./docker-volume:/var/lib/postgresql/data
    ports:
      - "5432:5432"
```

Start the development DB using the following command

```bash
$ docker-compose up
```