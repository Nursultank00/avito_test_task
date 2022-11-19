# Микросервис для работы с балансом пользователей
## Тестовое задание 

[Ссылка на задание](https://github.com/avito-tech/internship_backend_2022/blob/main/README.md)

## Эндпоинты
+ POST /accounts/create-account - создает нового пользователя. Тело запроса:
    - account_id - уникальный идентификатор пользователя
+ GET /accounts/all - возвращает всех пользователей.
+ GET /balance - получение баланса пользователя. Тело запроса:
    - account_id - идентификатор пользователя
+ POST /balance/accrue - пополнение баланса пользователя. Тело запроса:
    - account_id - идентификатор пользователя
    - amount - сумма начисления
    - description - описание транзакции
+ POST /balance/debit - снятие средств с баланса пользователя. Тело запроса:
    - account_id - идентификатор пользователя
    - amount - сумма начисления
    - description - описание транзакции
+ POST /balance/reserve - резервирование средств пользователя для услуги. Тело запроса:
    - transaction_id - идентификатор транзакции
    - account_id - идентификатор пользователя
    - service_id - идентификатор услуги
    - order_id - идентификатор заказа
    - amount - сумма начисления
    - description - описание транзакции
+ POST /balance/confirmation - снятие средств с резервированного баланса для услуги. Тело запроса:
    - transaction_id - идентификатор транзакции
    - account_id - идентификатор пользователя
    - service_id - идентификатор услуги
    - order_id - идентификатор заказа
    - amount - сумма начисления
    - description - описание транзакции
+ POST /balance/transfer - перевод средств от одного пользователя к другому. Тело запроса:
    - receiver_id - идентификатор получателя
    - sender_id - идентификатор отправителя
    - amount - сумма перевода
    - description - описание транзакции
+ GET /balance/transactions - вывод всех транзакций пользователя. Тело запроса:
    - account_id - уникальный идентификатор пользователя

## Запуск
Создаём образ 
```
docker compose build
```
Запускаем контейнер
```
docker compose up
```
Если запускается первый раз, нужно применить миграции:
```
migrate -path ./schema -database 'postgres://postgres:1234@localhost:5432/postgres?sslmode=disable' up
```


