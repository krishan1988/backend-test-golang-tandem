# Solution Overview
## Structure of the Project

To facilitate future scalability and maintainability, the single endpoint /api/v1/facts will be implemented using a modular framework. The following is an outline of the project's structure.

   * api/server: The main entry point of the application that initializes the HTTP server and handles routes.
   * config: Contains the all configurations of the application.
   * models: Handles the connection to the MongoDB database and provides data access methods.
   * mongo: Contains the mongo client 
   * mongo-seed: Include db json sample data
   * repo: table structs and query writter
   * service: added main service layer
   * cmd/backend-test-golang: Test functions written under this folder
  

### Technical Choices

   * Language: Go was chosen as the programming language due to its efficiency, simplicity, and excellent support for building web applications.
   * MongoDB: MongoDB was selected as the database for storing facts due to its flexibility, scalability, and ease of use.
   * MongoDB Driver: The official MongoDB driver for Go, go.mongodb.org/mongo-driver, was used to connect to the MongoDB database and perform database operations.
* HTTP Server: The standard library package net/http was utilized to create the HTTP server and handle API routes.
 * Query Parameters: The API endpoint allows query parameters for filtering and limiting the results. Parameters like text, number, found, and type can be used to customize the query and provide a rich user experience.

### Architecture

The implemented solution follows a builder design pattern and  Singleton design pattern to reduce object complexity, separate concerns and promote code maintainability and testability.


### Installation and Setup


```To install and launch the solution, follow these steps:```

* Clone the project repository from GitHub:

git clone https://github.com/KryptoKnight/backend-test-golang.git

* Navigate to the project directory:

bash

cd  backend-test-golang

* Create bin/dev/mongodb folders under path

* Run `docker compose up` command on terminal.

The API server will be available at http://localhost:8083/api/v1/facts.

* Run CURL request and see the results

curl --location 'http://localhost:8083/api/v1/facts?page=4' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTIyMDQyMjh9.TC2fGAQGa8U7sVZXL33y83fnvIEb6wk6c68rSE14B90m0Pa0m-xsKzNRAbZ5T1luN2BEC5v3169yWt0yPsSlO5OgY4EfbsnFiB1qq3KoR8JNXmgAd1TEtFobcrQXH7SSpkCLRT5_QrCedd2SJX6dsCNQiV_Zp90_CHPg71XdZlc'