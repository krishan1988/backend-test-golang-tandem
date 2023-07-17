# Solution Overview
## Structure of the Project

To facilitate future scalability and maintainability, the single endpoint `/api/v1/facts` will be implemented following layer architecture. The following is an outline of the project's structure.

   * __api/server:__ The main entry point of the application that initializes the HTTP server and handles routes.
   * __config__: Contains the all configurations of the application.
   * __models__: Handles the connection to the MongoDB database and provides data access methods.
   * __mongo__: Contains the mongo client 
   * __mongo-seed__: Instantiate a service which use to migrate given data to the `mongodb` service.
   * __repo__: Contain database entities and the data access layer.
   * __service__: Implemented the business related logics in this layer.
   * __cmd/backend-test-golang__: Contain the main package and application entry point.
  
### Technical Choices

   - __Language:__ Go was chosen as the programming language due to its efficiency, simplicity, and excellent support for building web applications.
   - __MongoDB:__ MongoDB was selected as the database for storing facts due to its flexibility, scalability, and ease of use.
   - __MongoDB Driver:__ The official MongoDB driver for Go, go.mongodb.org/mongo-driver, was used to connect to the MongoDB database and perform database operations.
- __HTTP Server:__ The standard library package net/http was utilized to create the HTTP server and handle API routes.
 - __Query Parameters:__ The API endpoint allows query parameters for filtering and limiting the results. Parameters like text, number, found, and type can be used to customize the query and provide a rich user experience.

### Architecture

The implemented solution follows a _builder design pattern_ and  _Singleton design pattern_ to reduce object complexity, separate concerns and promote code maintainability and testability.


### Installation and Setup

#### To install and launch the solution, follow these steps:

- Clone the project repository from GitHub:
```sh
git clone https://github.com/KryptoKnight/backend-test-golang.git
```
- Navigate to the project directory:
```
cd  backend-test-golang
```
- Create bin/dev/mongodb folders under path
- Run `docker compose up` command on terminal.

The API server will be available at http://localhost:8083/api/v1/facts.

- Run cURL request and see the results
> Note: For generate token for testing  refer [this](./tools/README.md)

```sh
curl --location 'http://localhost:8083/api/v1/facts?page=4' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer <token>'
```