# Hero API Challenge

This is a Golang API to store super heroes/villains, based on the [SuperHero API](https://superheroapi.com/). It allows you to create a new super, list by all heroes, list by all villains, and search by name and UUID.

It uses Docker to containerize the application.

## Go Version
Golang v1.17 was used on this project.
## External Dependencies
### External Data Sources:
- [SuperHero API](https://superheroapi.com/)
### External Packages
- [gorilla/mux](http://github.com/gorilla/mux)
- [urfave/negroni](http://github.com/urfave/negroni)
- [lib/pq](http://github.com/lib/pq)
- [google/uuid](http://github.com/google/uuid)
- [caarlos0/env](http://github.com/caarlos0/env/v6)
- [DATA-DOG/go-sqlmock](http://github.com/DATA-DOG/go-sqlmock)

## How to run

There is a `docker-compose.yml` file and a `Makefile`to automate application deployment in local environment. 

To run the application, simply:

- Start the database: `$ make up-db`;
- Create database schema: `$ make create-db-schema`;
- Build the API: `$ make build`;
- Start the API: `$ make up-api`;

Note: If any environment variable from `.env` file is changed, the database container must be deleted and recreated. Also, if `POSTGRES_USER`, `POSTGRES_HOST` or `POSTGRES_DB`, you MUST reflect those changes on `./utils/create_database.sh`.

## API Documentation

### /api/create
Creates a new super based in the SuperHero API data. 
Parameters: `name`.

API call: `/api/create?name=<super-name>`
Example API response: 

    // http://localhost:3000/api/create?name=joker
    
    {
      "created-count": "1",
      "status": "success"
    }

### /api/listAll
Lists all supers stored in the database.
Parameters: -

API Call: `/api/listAll`
Example API response:

    // 20220203133451
    // http://localhost:3000/api/listAll
    
    {
      "status": "success",
      "results": [
        {
          "uuid": "7b9e7d5d-44ec-46a2-bd22-6a4b9604beb9",
          "name": "Superman",
          "powerstats": {
            "intelligence": "94",
            "power": "100"
          },
          "biography": {
            "full-name": "Clark Kent",
            "alignment": "good"
          },
          "occupation": "Reporter for the Daily Planet and novelist",
          "relatives-count": 10,
          "group-affiliations": [
            "Justice League of America",
            "The Legion of Super-Heroes (pre-Crisis as Superboy); Justice Society of America (pre-Crisis Earth-2 version); All-Star Squadron (pre-Crisis Earth-2 version)"
          ],
          "image": "https://www.superherodb.com/pictures2/portraits/10/100/791.jpg"
        },
        {
          "uuid": "af5671e5-2ceb-4065-8c9f-f8a326fc4d6a",
          "name": "Joker",
          "powerstats": {
            "intelligence": "100",
            "power": "43"
          },
          "biography": {
            "full-name": "Jack Napier",
            "alignment": "bad"
          },
          "occupation": "-",
          "relatives-count": 1,
          "group-affiliations": [
            "Black Glove",
            "Injustice Gang",
            "Injustice League",
            "Joker League of Anarchy"
          ],
          "image": "https://www.superherodb.com/pictures2/portraits/10/100/719.jpg"
        },
        {
          "uuid": "6cb4a24b-cd82-40b2-b805-44bf13cf4abb",
          "name": "Joker",
          "powerstats": {
            "intelligence": "100",
            "power": "43"
          },
          "biography": {
            "full-name": "Jack Napier",
            "alignment": "bad"
          },
          "occupation": "-",
          "relatives-count": 1,
          "group-affiliations": [
            "Black Glove",
            "Injustice Gang",
            "Injustice League",
            "Joker League of Anarchy"
          ],
          "image": "https://www.superherodb.com/pictures2/portraits/10/100/719.jpg"
        },
        {
          "uuid": "7a5b739e-23ef-4d3c-9597-e7059afb0e48",
          "name": "Cyborg Superman",
          "powerstats": {
            "intelligence": "75",
            "power": "100"
          },
          "biography": {
            "full-name": "Henry Henshaw",
            "alignment": "bad"
          },
          "occupation": "-",
          "relatives-count": 1,
          "group-affiliations": [
            "Alpha Lantern Corps",
            "Manhunters",
            "Warworld",
            "formerly Apokolips and Sinestro Corps"
          ],
          "image": "https://www.superherodb.com/pictures2/portraits/10/100/667.jpg"
        },
        {
          "uuid": "f0cd4137-5de4-411e-9e3c-108ca8ef671b",
          "name": "Superman",
          "powerstats": {
            "intelligence": "94",
            "power": "100"
          },
          "biography": {
            "full-name": "Clark Kent",
            "alignment": "good"
          },
          "occupation": "Reporter for the Daily Planet and novelist",
          "relatives-count": 10,
          "group-affiliations": [
            "Justice League of America",
            "The Legion of Super-Heroes (pre-Crisis as Superboy); Justice Society of America (pre-Crisis Earth-2 version); All-Star Squadron (pre-Crisis Earth-2 version)"
          ],
          "image": "https://www.superherodb.com/pictures2/portraits/10/100/791.jpg"
        }
      ]
    }

### /api/listHeroes
Lists all super heroes stored in database.
Parameters: -

API Call: `/api/listHeroes`
Example API response:

    // 20220203133712
    // http://localhost:3000/api/listHeroes
    
    {
      "status": "success",
      "results": [
        {
          "uuid": "7b9e7d5d-44ec-46a2-bd22-6a4b9604beb9",
          "name": "Superman",
          "powerstats": {
            "intelligence": "94",
            "power": "100"
          },
          "biography": {
            "full-name": "Clark Kent",
            "alignment": "good"
          },
          "occupation": "Reporter for the Daily Planet and novelist",
          "relatives-count": 10,
          "group-affiliations": [
            "Justice League of America",
            "The Legion of Super-Heroes (pre-Crisis as Superboy); Justice Society of America (pre-Crisis Earth-2 version); All-Star Squadron (pre-Crisis Earth-2 version)"
          ],
          "image": "https://www.superherodb.com/pictures2/portraits/10/100/791.jpg"
        },
        {
          "uuid": "f0cd4137-5de4-411e-9e3c-108ca8ef671b",
          "name": "Superman",
          "powerstats": {
            "intelligence": "94",
            "power": "100"
          },
          "biography": {
            "full-name": "Clark Kent",
            "alignment": "good"
          },
          "occupation": "Reporter for the Daily Planet and novelist",
          "relatives-count": 10,
          "group-affiliations": [
            "Justice League of America",
            "The Legion of Super-Heroes (pre-Crisis as Superboy); Justice Society of America (pre-Crisis Earth-2 version); All-Star Squadron (pre-Crisis Earth-2 version)"
          ],
          "image": "https://www.superherodb.com/pictures2/portraits/10/100/791.jpg"
        }
      ]
    }

### /api/listVillains
Lists all villains stored in the database.
Parameters: -

API Call: `/api/listVillains`
Example API response:

    // 20220203133843
    // http://localhost:3000/api/listVillains
    
    {
      "status": "success",
      "results": [
        {
          "uuid": "af5671e5-2ceb-4065-8c9f-f8a326fc4d6a",
          "name": "Joker",
          "powerstats": {
            "intelligence": "100",
            "power": "43"
          },
          "biography": {
            "full-name": "Jack Napier",
            "alignment": "bad"
          },
          "occupation": "-",
          "relatives-count": 1,
          "group-affiliations": [
            "Black Glove",
            "Injustice Gang",
            "Injustice League",
            "Joker League of Anarchy"
          ],
          "image": "https://www.superherodb.com/pictures2/portraits/10/100/719.jpg"
        },
        {
          "uuid": "6cb4a24b-cd82-40b2-b805-44bf13cf4abb",
          "name": "Joker",
          "powerstats": {
            "intelligence": "100",
            "power": "43"
          },
          "biography": {
            "full-name": "Jack Napier",
            "alignment": "bad"
          },
          "occupation": "-",
          "relatives-count": 1,
          "group-affiliations": [
            "Black Glove",
            "Injustice Gang",
            "Injustice League",
            "Joker League of Anarchy"
          ],
          "image": "https://www.superherodb.com/pictures2/portraits/10/100/719.jpg"
        },
        {
          "uuid": "7a5b739e-23ef-4d3c-9597-e7059afb0e48",
          "name": "Cyborg Superman",
          "powerstats": {
            "intelligence": "75",
            "power": "100"
          },
          "biography": {
            "full-name": "Henry Henshaw",
            "alignment": "bad"
          },
          "occupation": "-",
          "relatives-count": 1,
          "group-affiliations": [
            "Alpha Lantern Corps",
            "Manhunters",
            "Warworld",
            "formerly Apokolips and Sinestro Corps"
          ],
          "image": "https://www.superherodb.com/pictures2/portraits/10/100/667.jpg"
        }
      ]
    }

### /api/search
Search for matching supers based on UUID or Name.
Parameters: `uuid`, `name`.

API Call: `/api/search?uuid=<uuid>` or `/api/search?name=<name>`
Example API Response:

    // 20220203134116
    // http://localhost:3000/api/search?name=joker
    
    {
      "status": "success",
      "results": [
        {
          "uuid": "af5671e5-2ceb-4065-8c9f-f8a326fc4d6a",
          "name": "Joker",
          "powerstats": {
            "intelligence": "100",
            "power": "43"
          },
          "biography": {
            "full-name": "Jack Napier",
            "alignment": "bad"
          },
          "occupation": "-",
          "relatives-count": 1,
          "group-affiliations": [
            "Black Glove",
            "Injustice Gang",
            "Injustice League",
            "Joker League of Anarchy"
          ],
          "image": "https://www.superherodb.com/pictures2/portraits/10/100/719.jpg"
        },
        {
          "uuid": "6cb4a24b-cd82-40b2-b805-44bf13cf4abb",
          "name": "Joker",
          "powerstats": {
            "intelligence": "100",
            "power": "43"
          },
          "biography": {
            "full-name": "Jack Napier",
            "alignment": "bad"
          },
          "occupation": "-",
          "relatives-count": 1,
          "group-affiliations": [
            "Black Glove",
            "Injustice Gang",
            "Injustice League",
            "Joker League of Anarchy"
          ],
          "image": "https://www.superherodb.com/pictures2/portraits/10/100/719.jpg"
        }
      ]
    }

### /api/delete
Deletes one or more supers stored in the database.
Parameters: `uuid`, separated by `,`(comma) if more than one.

API Call: `/api/delete?uuid=<uuid>,<uuid>`
Example API response:

    // 20220203134425
    // http://localhost:3000/api/delete?uuid=7b9e7d5d-44ec-46a2-bd22-6a4b9604beb9,af5671e5-2ceb-4065-8c9f-f8a326fc4d6a
    
    {
      "status": "success",
      "deleted-count": 2
    }

## Unit Tests

To run all unit tests, simply run `$ make tests`.