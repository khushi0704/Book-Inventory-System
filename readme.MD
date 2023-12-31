# A Simple HTTP (REST) API using GoFr
Book Inventory System is a CRUD Application which is built using the Gofr framework(Go Language framework) and mySQL for database connection. This application empowers users to perform fundamental CRUD operations—enabling the creation, retrieval, updating, and deletion of books.

## Prerequisites

Before you begin, ensure you have the following installed on your machine:

- [Go](https://go.dev/doc/install)
- [mySQL](https://www.mysql.com/downloads/)
- [PostMan](https://www.postman.com/downloads/)

## Getting Started

1. Clone the repository:

   ```bash
   git clone https://github.com/khushi0704/Book-Inventory-System.git
   
   cd Book-Inventory-System
2. Download Dependencies:

   ```bash
   go get gofr.dev

3. Verify Dependencies:

    ```bash
   go mod tidy
4. Run Project
    
    ```bash
    go run main.go

5. Open Server
    
    ```bash
    http://localhost:8000/books
    
## API Endpoints check using Postman

```bash
https://www.postman.com/research-cosmologist-13477158/workspace/book-inventory-system/collection/31770885-5a0012d2-b962-41e8-b968-42503741ae8c?action=share&creator=31770885
```

The application provides RESTful API endpoints for CRUD operations on book records.

### Base URL

The base URL for all endpoints is:

### List of Endpoints

#### 1. **Get All Books**

- **Endpoint:**
  - `GET /books`

- **Description:**
  - Retrieves a list of all books.

![view_books](https://github.com/khushi0704/Book-Inventory-System/assets/64957861/1ca28fe5-666e-4341-9db1-f57b05875e80)

#### 2. **Get Book by Title**

- **Endpoint:**
  - `GET /book`

- **Description:**
  - Retrieves information about a specific student based on the provided ID.


![get_book](https://github.com/khushi0704/Book-Inventory-System/assets/64957861/fd7815d0-3e7e-4b5a-99bc-216d3ae233e9)

#### 3. **Create a New Book**

- **Endpoint:**
  - `POST /books`

- **Description:**
  - Creates a new book record.
    

![create_book](https://github.com/khushi0704/Book-Inventory-System/assets/64957861/d678c382-06f7-4ad2-91d6-c3c5cdede49a)

#### 4. **Update Book Quantity Available by Title**

- **Endpoint:**
  - `PATCH /book`

- **Description:**
  - Updates information about a specific book based on the provided title.

![update_book](https://github.com/khushi0704/Book-Inventory-System/assets/64957861/30b5da20-7a7e-42db-99e4-62ada32128cc)
 
#### 5. **Delete Book by ID**

- **Endpoint:**
  - `DELETE /book`

- **Description:**
  - Deletes a specific book record based on the provided ID.

![delete_book](https://github.com/khushi0704/Book-Inventory-System/assets/64957861/bbc45544-353a-4265-b6f8-919432bb8696)

# Database 🛢️

The project utilizes a MySQL database to store information about book records. The MySQL database is set up as a Docker container for easy deployment and management.

## Database Schema 🗃️

- Table: `book`
  - Columns: `book_id` (int) auto-increment primary key, `Title` (varchar), `Author` (varchar), Price (int), QuantityAvailable (int)


## Database Setup 🐳
- Docker Container
![docker](https://github.com/khushi0704/Book-Inventory-System/assets/64957861/1872ffb1-b4fd-42f1-b5ed-ba2f04f4b420)

- To run the MySQL database locally as a Docker container, use the following command:
```bash
docker run --name gofr-mysql -e MYSQL_ROOT_PASSWORD=1234 -e MYSQL_DATABASE=book_db -p 3308:3306 -d mysql:8.0.25
```

### Running Tests

- Use the following command to run the tests:

  ```bash
  go test


## UML diagrams
* Use Case Diagram:
  
![use_case](https://github.com/khushi0704/Book-Inventory-System/assets/64957861/cf5ca9ee-27b5-44fd-90c2-8f5b24bb8536)

* Sequence Diagram: 

![sequence](https://github.com/khushi0704/Book-Inventory-System/assets/64957861/4f400e5d-4ba3-46d8-b51d-ad050efd431f)

