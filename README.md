# mini-project-order-api
Rest API mini project build with golang, gorm, jwt, gin and postgres using model view controller architecture.
The feature of this mini project are register, login, create, read, update and delete
## Project Structure

## End Point
GET /orders ----> get all orders

GET /orders/:id ----> get orders by id

POST /orders ----> create new orders

PUT /orders/:id ----> update orders by id

DELETE /orders/:id ----> delete orders by id

GET /orders?page=0&limit=2 ----> Pagination with page 0 and limit 2 (can customize the page and limit)

POST /register ----> create new user

POST /login ----> login 

