# mini-project-order-api
Rest API mini project build with golang, gorm, jwt, gin and postgre SQL using model view controller architecture.
## Key Featrues
- Auth with Login and Register
- JWT Token
- CRUD Orders
## Project Structure
```
pkg    
│
└───config
│   │db.go   
│   
└───controller
|   │authcontroller.go
|   |ordercontroller.go
|
└───dto
|   |orderpagination.go
└───helper
|   |middlewares.go
|   |pagination.go   
|
└───logger
|   |logger.go
|
└───models
|   |order.go
|   |user.go
| 
└───routes
|   |routes.go
|
└───tests
|   |orderController_test.go
|   
└───utils
|   |token.go
|
|env
|go.mod
|go.sum
|main.go
|README.md
```
## SQL 
This mini project using postgreSQL
### Users
```
DROP TABLE IF EXISTS "users";
CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "user_name" varchar(100) NOT NULL,
  "password" varchar NOT NULL
);
```
### Orders
```
DROP TABLE IF EXISTS "orders";
CREATE TABLE "orders" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "product_name" varchar(100) NOT NULL,
  "status" varchar NOT NULL
);

INSERT INTO "orders"  VALUES
	(DEFAULT,'Lenovo Ideapad','1'),
	(DEFAULT,'Pocophone','1'),
	(DEFAULT,'Dual Sense','1'),
	(DEFAULT,'Aripods 2','1'),
	(DEFAULT,'MBP M1','1'),
	(DEFAULT,'LG Monitor','1'),
	(DEFAULT,'MBA M2 ','0'),
	(DEFAULT,'Airpods Pro','0'),
	(DEFAULT,'IWatch 7','0');
```
## End Point
GET /orders ----> get all orders

GET /orders/:id ----> get orders by id

POST /orders ----> create new orders

PUT /orders/:id ----> update orders by id

DELETE /orders/:id ----> delete orders by id

GET /orders?page=0&limit=2 ----> Pagination with page 0 and limit 2 (can customize the page and limit)

POST /register ----> create new user

POST /login ----> login

## Example Body and Result Create Order
### Body create order POST localhost:8080/orders
```
{
    "product_name": "Macbook Air M2 2022",
    "status": "0"
}
```
### Result
```
{
    "data": {
        "id": 14,
        "product_name": "Macbook Air M2 2022",
        "status": "0"
    }
}
```
## Example Result Get Orders, Orders by Id and Using Pagination
### Get all orders GET localhost:8080/orders
```
{
    "data": {
        "limit": 10,
        "page": 0,
        "total_rows": 12,
        "rows": [
            {
                "id": 1,
                "product_name": "Lenovo Ideapad",
                "status": "1"
            },
            {
                "id": 2,
                "product_name": "Pocophone",
                "status": "1"
            },
            {
                "id": 3,
                "product_name": "Dual Sense",
                "status": "1"
            },
            {
                "id": 4,
                "product_name": "Aripods 2",
                "status": "1"
            },
            {
                "id": 5,
                "product_name": "MBP M1",
                "status": "1"
            },
            {
                "id": 6,
                "product_name": "LG Monitor",
                "status": "1"
            },
            {
                "id": 7,
                "product_name": "MBA M2 ",
                "status": "0"
            },
            {
                "id": 8,
                "product_name": "Airpods Pro",
                "status": "0"
            },
            {
                "id": 12,
                "product_name": "Macbook Pro M1 Pro",
                "status": "0"
            },
            {
                "id": 9,
                "product_name": "MBA 2022",
                "status": "0"
            }
        ]
    }
}
```
### Get orders by Id GET localhost:8080/orders/6
```
{
    "data": {
        "id": 6,
        "product_name": "LG Monitor",
        "status": "1"
    }
}
```
### Get orders using pagination GET localhost:8080/orders?page=0&limit=4
```
{
    "data": {
        "limit": 4,
        "page": 0,
        "total_rows": 12,
        "rows": [
            {
                "id": 1,
                "product_name": "Lenovo Ideapad",
                "status": "1"
            },
            {
                "id": 2,
                "product_name": "Pocophone",
                "status": "1"
            },
            {
                "id": 3,
                "product_name": "Dual Sense",
                "status": "1"
            },
            {
                "id": 4,
                "product_name": "Aripods 2",
                "status": "1"
            }
        ]
    }
}
```
## Example Body and Result Update Order
### Body update order PUT localhost:8080/orders/12
```
{
    "product_name": "Macbook Pro M1 Max 2022",
    "status": "1"
}
```
### Result
```
{
    "data": {
        "id": 12,
        "product_name": "Macbook Pro M1 Max 2022",
        "status": "1"
    }
}
```
## Example Result Delete Order
### Result delete order DELETE localhost:8080/orders/12
```
{
    "data": true
}
```

## Example Body and Result Register and Login
### Body Register POST localhost:8080/register
```
{
    "user_name": "user23",
    "password": "rahasia123"
}
```
### Result Register
```
{
    "message": "registration success"
}
```
### Body Login POST localhost:8080/login
```
{
    "user_name": "user23",
    "password": "rahasia123"
}
```
### Result Login with token
```
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NjAwNjIyOTgsInVzZXJfaWQiOjR9.rwTka3M69fLOQ9XOce1CZ-go_KZw1GSnpqi3Dt5oegY"
}
```
