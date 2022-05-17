# Go-Bookstore

This GitHub repo has my attempt to build simple REST APIs for practicing go.


## Introduction
The project contains simple APIs for adding, editting, retrieving and deleting books.

This is used for educational and practicing purpose to learn design principals, Go lang and other DevOps toolchain like Docker.


## Design

The code base follows Hexagonal Architecture and Domain-Driven Design principals, where the code is broken down to:

- Domain
  - Model
  - Service

- Ports
- Adapters

## APIs:
### Book
 - Get all books:
 
    - **URL** : `/book`

    - **Method** : `GET`

  - Add book:
 
    - **URL** : `/book`

    - **Method** : `POST`

    - **Data constraints** : `{ "name": 'test',author" : 'test', "publication": 'test' }`

  - Get book:

     - **URL** : `/book/[id]`

     - **Method** : `GET`
     
  - Update book:

     - **URL** : `/book/[id]`

     - **Method** : `PUT`

      - **Data constraints** : `{ "name": 'test',author" : 'test', "publication": 'test' }`

  - Delete book:

      - **URL** : `/book/[id]`

      - **Method** : `DELETE`


## Build and Run
- Install Go and Mysql

- Clone the repo on your machine

- Run
  - From root of the repo on your terminal run the following command
  ```
  go run main.go
  ```

  - (Optional) Use docker compose
  ```
  docker compose build
  docker compose up
  ```

## To be done:
- Authentication
