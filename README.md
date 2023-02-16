# RESTAPIDocs Examples

These examples were taken from projects mainly using Golang Gin Frmaework and so the
JSON responses are often similar to the way in which DRF makes responses. 

for postman collection i added in root file 

ERD this project is [Click Here](https://app.diagrams.net/#G1JgT5loWmWMi9XbD3XOEbDyn5LcppH2cg)

first of all you must setting env and then, : 

## Download all dependencies. 
    go mod download

## Migrate the table
    go run .\migrate\migrate.go

## Run the tests 
    go run main.go

# REST API

The REST API to the example app is described below.

## Sign Up

### Request

`POST /api/auth/signup`

    curl --location --request POST 'http://localhost:8000/api/auth/signup' \
    --header 'Content-Type: application/json' \
    --data-raw '{"email" : "aisya1243@gmail.com","fullname" : "Aisya Putri","password" : "asdqwe123","passwordConfirm" : "asdqwe123"}'

### Response

    HTTP/1.1 200 OK
    Date: Thu, 24 Feb 2011 12:36:30 GMT
    Status: 200 OK
    Connection: close
    Content-Type: application/json
    Content-Length: 2
    
    {"access_token": "{{token}}","message": "Success Registration","status": "Success"}
    
## Log In

### Request

`POST /api/auth/login`

    curl --location --request POST 'http://localhost:8000/api/auth/login' \
    --header 'Content-Type: application/json' \
    --data-raw '{"email" : "aisya1243@gmail.com","password" : "asdqwe123"}'

### Response

    HTTP/1.1 200 OK
    Date: Thu, 24 Feb 2011 12:36:30 GMT
    Status: 200 OK
    Connection: close
    Content-Type: application/json
    Content-Length: 2
    
    {"access_token": "{{token}}","status": "Success"}
    
## Request Token 

### Request

`GET /api/auth/token`

    curl --location --request GET 'http://localhost:8000/api/auth/token' \
    --header 'Cookie: access_token={{token}}; logged_in=true; refresh_token={{refresh_token}}' \

### Response

    HTTP/1.1 200 OK
    Date: Thu, 24 Feb 2011 12:36:30 GMT
    Status: 200 OK
    Connection: close
    Content-Type: application/json
    Content-Length: 2
    
    {"access_token": "{{token}}","status": "Success"}
    
## Log Out 

### Request

`GET /api/auth/logout`

    curl --location --request GET 'http://localhost:8000/api/auth/logout' \ 
    --header 'Cookie: access_token={{token}}; logged_in=true; refresh_token={{refresh_token}}' \

### Response

    HTTP/1.1 200 OK
    Date: Thu, 24 Feb 2011 12:36:30 GMT
    Status: 200 OK
    Connection: close
    Content-Type: application/json
    Content-Length: 2
    
    {"status": "Success"}

## Get Products

### Request

`POST /api/product/listproducts`

    curl --location --request POST 'http://localhost:8000/api/product/listproducts' \
    --header 'Content-Type: application/json' \
    --header 'Cookie: access_token={{token}}; logged_in=true; refresh_token={{refresh_token}}' \
    --data-raw '{"page" : "1","perPage" : "10","search" : "","sort" : "","sortColumn" : "","category" : ""}'

### Response

    HTTP/1.1 200 OK
    Date: Thu, 24 Feb 2011 12:36:30 GMT
    Status: 200 OK
    Connection: close
    Content-Type: application/json
    Content-Length: 2
    
    {
    "count": 4,
    "data": [
        {
            "id_product": 1,
            "id_category": 1,
            "name": "Clothes ASD",
            "description": "Example Description Clothes ASD",
            "price": 56000,
            "deleted": "N",
            "id_created": 0,
            "id_updated": 0,
            "created_at": "2023-02-06T10:49:20.041+07:00",
            "updated_at": "2023-02-06T10:49:20.041+07:00",
            "detail_category": {
                "id_category": 1,
                "name": "Clothes",
                "description": "Example Description Clothes",
                "deleted": "N",
                "id_created": 0,
                "id_updated": 0,
                "created_at": "2023-02-06T10:49:20.041+07:00",
                "updated_at": "2023-02-06T10:49:20.041+07:00"
            }
        },
    ],
    "status": "success"
    }
    
## Add Cart

### Request

`POST /api/product/addcart`

    curl --location --request POST 'http://localhost:8000/api/product/addcart' \
    --header 'Content-Type: application/json' \
    --header 'Cookie: access_token={{token}}; logged_in=true; refresh_token={{refresh_token}}' \
    --data-raw '{"product" : 1}'

### Response

    HTTP/1.1 200 OK
    Date: Thu, 24 Feb 2011 12:36:30 GMT
    Status: 200 OK
    Connection: close
    Content-Type: application/json
    Content-Length: 2
    
    {"data": "Success Add Product to Cart","status": "success"}
    
## Get Cart

### Request

`GET /api/product/listcart`

    curl --location --request GET 'http://localhost:8000/api/product/listcart' \
    --header 'Cookie: access_token={{token}}; logged_in=true; refresh_token={{refresh_token}}' \

### Response

    HTTP/1.1 200 OK
    Date: Thu, 24 Feb 2011 12:36:30 GMT
    Status: 200 OK
    Connection: close
    Content-Type: application/json
    Content-Length: 2
    
    {
    "data": [
        {
            "id_product": 1,
            "id_user": 2,
            "qty": 1,
            "detail_product": {
                "id_product": 1,
                "id_category": 1,
                "name": "Clothes ASD",
                "description": "Example Description Clothes ASD",
                "price": 56000,
                "deleted": "N",
                "id_created": 0,
                "id_updated": 0,
                "created_at": "2023-02-06T10:49:20.041+07:00",
                "updated_at": "2023-02-06T10:49:20.041+07:00",
                "detail_category": {
                    "id_category": 1,
                    "name": "Clothes",
                    "description": "Example Description Clothes",
                    "deleted": "N",
                    "id_created": 0,
                    "id_updated": 0,
                    "created_at": "2023-02-06T10:49:20.041+07:00",
                    "updated_at": "2023-02-06T10:49:20.041+07:00"
                }
            },
            "detail_user": {
                "id_user": 2,
                "email": "aisya1243@gmail.com",
                "fullname": "Aisya Putri",
                "role": "customer",
                "status": "Active",
                "deleted": "N",
                "id_created": 0,
                "id_updated": 0,
                "created_at": "2023-02-06T10:50:04.953+07:00",
                "updated_at": "2023-02-06T10:50:04.953+07:00"
            }
        }
    ],
    "status": "success"
    }
    
## Update Cart

### Request

`POST /api/product/updatecart`

    curl --location --request POST 'http://localhost:8000/api/product/updatecart' \
    --header 'Content-Type: application/json' \
    --header 'Cookie: access_token={{token}}; logged_in=true; refresh_token={{refresh_token}}' \
    --data-raw '{"product" : 1,"qty" : 3,"method" : "update"}'
    or
    --data-raw '{"product" : 1,"qty" : 3,"method" : "delete"}'

### Response

    HTTP/1.1 200 OK
    Date: Thu, 24 Feb 2011 12:36:30 GMT
    Status: 200 OK
    Connection: close
    Content-Type: application/json
    Content-Length: 2
    
    {"data": "Success Add Product to Cart","status": "success"}
    
## Get Payment Method

### Request

`GET /api/product/updatecart`

    curl --location --request GET 'http://localhost:8000/api/transaction/listpaymentmethod'
    --header 'Cookie: access_token={{token}}; logged_in=true; refresh_token={{refresh_token}}' \

### Response

    HTTP/1.1 200 OK
    Date: Thu, 24 Feb 2011 12:36:30 GMT
    Status: 200 OK
    Connection: close
    Content-Type: application/json
    Content-Length: 2
    
    {
    "data": [
        {
            "iIdPaymentMethod": 1,
            "name": "BRI",
            "description": "18236791263 a/n BRI",
            "deleted": "N",
            "id_created": 0,
            "id_updated": 0,
            "created_at": "2023-02-06T10:49:20.155+07:00",
            "updated_at": "2023-02-06T10:49:20.155+07:00"
        },
    ],
    "status": "success"
    }
    
## Get List Transaction

### Request

`POST /api/transaction/listtransaction`

    curl --location --request POST 'http://localhost:8000/api/transaction/listtransaction' \
    --header 'Content-Type: application/json' \
    --header 'Cookie: access_token={{token}}; logged_in=true; refresh_token={{refresh_token}}' \
    --data-raw '{"page" : "1","perPage" : "10","search" : "","sort" : "","sortColumn" : ""}'

### Response

    HTTP/1.1 200 OK
    Date: Thu, 24 Feb 2011 12:36:30 GMT
    Status: 200 OK
    Connection: close
    Content-Type: application/json
    Content-Length: 2
    
    {
    "count": 1,
    "data": [
        {
            "id_transaction_header": 1,
            "name_customer": "Aisya Putri",
            "no_transaction": "NIycGdKtEgoMCbsEedrROtJhNkEPbN",
            "payment_method_name": "BCA",
            "payment_method_desc": "BCA a/n 19237912873",
            "total": 17000000,
            "deleted": "N",
            "id_created": 2,
            "id_updated": 2,
            "created_at": "2023-02-06T11:28:06.762+07:00",
            "updated_at": "2023-02-06T11:28:06.762+07:00",
            "detail": [
                {
                    "id_transaction_detail": 1,
                    "id_transaction_header": 1,
                    "name_product": "Laptop KLJ",
                    "description_product": "Example Description Laptop KLJ",
                    "qty": 2,
                    "price_product": 5000000,
                    "total": 10000000,
                    "deleted": "N",
                    "id_created": 2,
                    "id_updated": 2,
                    "created_at": "2023-02-06T11:28:06.764+07:00",
                    "updated_at": "2023-02-06T11:28:06.764+07:00"
                },
                {
                    "id_transaction_detail": 2,
                    "id_transaction_header": 1,
                    "name_product": "Laptop POI",
                    "description_product": "Example Description Laptop POI",
                    "qty": 1,
                    "price_product": 7000000,
                    "total": 7000000,
                    "deleted": "N",
                    "id_created": 2,
                    "id_updated": 2,
                    "created_at": "2023-02-06T11:28:06.768+07:00",
                    "updated_at": "2023-02-06T11:28:06.768+07:00"
                }
            ]
        }
    ],
    "status": "success"
    }
    
## Post Transaction

### Request

`POST /api/transaction/transaction`

    curl --location --request POST 'http://localhost:8000/api/transaction/addtransaction' \
    --header 'Content-Type: application/json' \
    --header 'Cookie: access_token={{token}}; logged_in=true; refresh_token={{refresh_token}}' \
    --data-raw '{
        "payment_method_name" : "BCA",
        "payment_method_desc" : "BCA a/n 19237912873",
        "total" : 17000000,
        "detail" : [
            {
                "name_product" : "Laptop KLJ",
                "description_product" : "Example Description Laptop KLJ",
                "price_product" : 5000000,
                "qty" : 2
            },
            {
                "name_product" : "Laptop POI",
                "description_product" : "Example Description Laptop POI",
                "price_product" : 7000000,
                "qty" : 1
            }
        ]
    }'

### Response

    HTTP/1.1 200 OK
    Date: Thu, 24 Feb 2011 12:36:30 GMT
    Status: 200 OK
    Connection: close
    Content-Type: application/json
    Content-Length: 2
    
    {"message": "Transaction has been successfully","status": "success"}
    

    
