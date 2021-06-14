# Simple Checkout

Simple Checkout is a simple backend service using GraphQL handling a simple ecommerce-like service.

This service supports:
1. Product Management (Add - Lists - Update - Delete), first two features is already implemented
2. Promo Management (Add - Lists - Update - Delete), first two features is already implemented
3. Checkout Process 

The background of this service is handling following promotions when a checkout process initiated.
1. Buy 1 Product X Get 1 Free Product Y
2. Buy 3 Products X, Pay Only 2 Product X
3. Buy Min 3 Products X, Get 10% Discount for purchases of Products X

Based on these requirements, I define three general rules.
1. Buy 1 Get 1 - Buy1Get1
2. Buy X Pay Y - BuyXPayY
3. Buy Min X, OFF Y% - BuyXPercentageOffY

So the variables X & Y can be anything and be stored in a database.

The flow process of this services, it is easy to understand with images.

<img src="https://github.com/sigitisme/simplecheckout/blob/main/img/flow.jpg">



## Build

  make

## Run tests

  make test

## API requests 

### Add Product #1

```
curl --location --request POST 'http://localhost:8080/query' \
--header 'Content-Type: application/json' \
--data-raw '{"query":"mutation AddProduct1{\n  createProduct(input:{sku:\"120P90\", name:\"Google Home\", price:49.99, quantity:10}){\n    ID\n    sku\n    name\n  }\n}","variables":{}}'

```

### Add Product #2

```
curl --location --request POST 'http://localhost:8080/query' \
--header 'Content-Type: application/json' \
--data-raw '{"query":"mutation AddProduc2{\n  createProduct(input:{sku:\"43N23P\", name:\"Macbook Pro\", price:5399.99, quantity:5}){\n    ID\n    sku\n    name\n  }\n}\n","variables":{}}'
```

### Add Product #3

```
curl --location --request POST 'http://localhost:8080/query' \
--header 'Content-Type: application/json' \
--data '{"query":"mutation AddProduct3{\n  createProduct(input:{sku:\"234234\", name:\"Raspberry Pi B\", price:30.00, quantity:2}){\n    ID\n    sku\n    name\n  }\n}","variables":{}}'
```

### Add Product #4

```
curl --location --request POST 'http://localhost:8080/query' \
--header 'Content-Type: application/json' \
--data-raw '{"query":"mutation AddProduct4{\n  createProduct(input:{sku:\"A304SD\", name:\"Alexa Speaker\", price:109.50, quantity:10}){\n    ID\n    sku\n    name\n  }\n}","variables":{}}'
```

### Add Promo #1

```
curl --location --request POST 'http://localhost:8080/query' \
--header 'Content-Type: application/json' \
--data-raw '{"query":"mutation AddPromo1 {\n  createPromo(input:{sku:\"43N23P\", scheme:\"Buy1Get1\",freebiesku:\"234234\",minqty:0,payonly:0,percentageoff:0}){\n    ID\n    sku\n    scheme\n    freebiesku\n  }\n}","variables":{}}'
```

### Add Promo #2

```
curl --location --request POST 'http://localhost:8080/query' \
--header 'Content-Type: application/json' \
--data-raw '{"query":"mutation AddPromo2 {\n  createPromo(input:{sku:\"43N23P\", scheme:\"BuyXPayY\",freebiesku:\"\",minqty:3,payonly:2,percentageoff:0}){\n    ID\n    sku\n    scheme\n    minqty\n  payonly\n  }\n}","variables":{}}'
```

### Add Promo #3

```
curl --location --request POST 'http://localhost:8080/query' \
--header 'Content-Type: application/json' \
--data-raw '{"query":"mutation AddPromo3 {\n  createPromo(input:{sku:\"A304SD\", scheme:\"BuyMinXOffY\",freebiesku:\"\",minqty:3,payonly:0,percentageoff:10}){\n    ID\n    sku\n    scheme\n    minqty\n  percentageoff\n  }\n}","variables":{}}'
```


### Checkout #1

```
curl --location --request POST 'http://localhost:8080/query' \
--header 'Content-Type: application/json' \
--data-raw '{"query":"query testCheckout {\n  checkout(input:{contents:[{sku:\"43N23P\",name:\"Macbook Pro\", quantity:1},{sku:\"234234\",name:\"Raspberry Pi B\", quantity:1}]}){\n    total\n  },\n}","variables":{}}'
```

### Checkout #2

```
curl --location --request POST 'http://localhost:8080/query' \
--header 'Content-Type: application/json' \
--data-raw '{"query":"query testCheckout2 {\n  checkout(input:{contents:[{sku:\"120P90\",name:\"Google Home\", quantity:3]}){\n    total\n  },\n}","variables":{}}'
```


### Checkout #3

```
curl --location --request POST 'http://localhost:8080/query' \
--header 'Content-Type: application/json' \
--data-raw '{"query":"query testCheckout2 {\n  checkout(input:{contents:[{sku:\"A304SD\",name:\"Alexa Speaker\", quantity:3]}){\n    total\n  },\n}","variables":{}}'
```

### Postman Collection

```
https://www.getpostman.com/collections/2a4c44a949af973c9255
```
