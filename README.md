# product-elasticsearch

## How to Run

1. Create env file and fill the value or run

   * cp .env.example .env
2. Simply run

   * docker compose up
3. then if you want to migrate sample data, i already put it in product.json, so you can run

   * curl -X POST "http://localhost:9200/_bulk" -H 'Content-Type: application/json' --data-binary @product.json
4. if you want to run locally (without docker), please make sure the value env file is correct

## Swagger Doc

http://localhost:8080/swagger/index.html

Thank you
