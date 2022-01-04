# payment-app
## Для запуска сервера:
1. склонируйте проект
2. docker-compose up --build
3. ctrl-c
4. docker-compose up &nbsp;

to see endpoints you can use Postman collection: https://www.getpostman.com/collections/a9a535a4ae583da2cc20

this is the payment app with two microservices:
### payment_app_auth - this service is responsible for all auth related services. clone this service from repository https://github.com/DYernar/payment-system-auth
  available api's:
  1. localhost:8111/signup &nbsp;
    Method POST,
    Body : 
      {
          "iin": "000614500238",
          "login": "bobby",
          "password": "12345"
      },
  2. localhost:8111/login &nbsp;
    Method: POST,
    Body: 
      {
          "login": "admin",
          "password": "12345"
      },
    
  3. localhost:8111/user/userLogin  -  возвращает детали пользователя с login - ом userLogin. &nbsp;
  4. localhost:8111/create/admin - используется для создания пользователся с ролью администратор. Для создания клиент тоже должен быть администратором. &nbsp;
    METHOD: POST
    Header: {
      "Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"
    }
  5. localhost:8111/get/all/users - возвращает всех пользователей в системе &nbsp;


### payment_app_parsing - this service is responsible for transactions, creating a wallet and filling the wallet with money. Clone it from repository https://github.com/DYernar/payment-system-parsing &nbsp;
  available api's: &nbsp;
  1. localhost:8112/create/wallet - создание кошелька для пользователя, в теле запроса нужно отправить имя кошелька &nbsp;
    METHOD: POST,
    Header: {
      "Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"
    },
    Body: {
        "name": "visa"
    },
  2. localhost:8112/transfer - перевод денег с одного кошелька в другую &nbsp;
    METHOD: POST,
    Header: {
      "Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"
    },
    Body: {
        "from": 1,
        "to": 5,
        "amount": 30
    },
  3. localhost:8112/add/money - пополнение кошелька &nbsp;
    METHOD: POST,
    Body: {
        "id": 1,
        "amount": 4500
    },
  4. localhost:8112/get/all/transactions  - возвращает список всех транзакции.пользователь должен быть админом. &nbsp;
    METHOD: GET,
    Header: {
      "Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"
    },
