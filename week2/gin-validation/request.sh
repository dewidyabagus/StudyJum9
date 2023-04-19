curl --location 'http://localhost:8080/registry' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "email@example.com",
    "first_name": "widya ade",
    "last_name": "bagus",
    "password": "123456",
    "confirm_password": "123456",
    "options": [1]
}'
