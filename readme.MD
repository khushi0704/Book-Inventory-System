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
    
## API Endpoints

The application provides RESTful API endpoints for CRUD operations on book records.

### Base URL

The base URL for all endpoints is:

### List of Endpoints

#### 1. **Get All Books**

- **Endpoint:**
  - `GET /books`

- **Description:**
  - Retrieves a list of all books.

![view_books](https://github.com/khushi0704/Book-Inventory-System/blob/main/assets/view_books.png)

#### 2. **Get Book by Title**

- **Endpoint:**
  - `GET /book`

- **Description:**
  - Retrieves information about a specific student based on the provided ID.


![get_book](https://github.com/khushi0704/Book-Inventory-System/blob/main/assets/get_book.png)

#### 3. **Create a New Book**

- **Endpoint:**
  - `POST /books`

- **Description:**
  - Creates a new book record.
    

![create_book](https://github.com/khushi0704/Book-Inventory-System/blob/main/assets/create_book.png)

#### 4. **Update Book Quantity Available by Title**

- **Endpoint:**
  - `PATCH /book`

- **Description:**
  - Updates information about a specific book based on the provided title.

  ![update_book](https://github.com/khushi0704/Book-Inventory-System/blob/main/assets/update_book.png)
 
#### 5. **Delete Book by ID**

- **Endpoint:**
  - `DELETE /book`

- **Description:**
  - Deletes a specific book record based on the provided ID.

![delete_book](https://github.com/khushi0704/Book-Inventory-System/blob/main/assets/delete_book.png)

### Running Tests

- Use the following command to run the tests:

  ```bash
  go test


## UML diagrams
* Use Case Diagram:
  
![use_case](https://github.com/khushi0704/Book-Inventory-System/blob/main/assets/use_case.png)

* Sequence Diagram: 

![sequence](https://github.com/khushi0704/Book-Inventory-System/blob/main/assets/sequence.png)

