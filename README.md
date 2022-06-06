# test pay API
### docker-compose -up для старта
```
POST   /api/create_pay
body
{
    "id_user": 2,
    "email":"test@tet.com",
    "price":1400,
    "price_type":"USD",
    "status":"NEW"
}
```
### Чтоб отправлять запрос на ```/api/pay_system/edit_pay``` необходимо получить Token ```/api/generate_token```
#### после создания платежа, делается случайный ```status```, при попытке изменить когда он в состоянии ```FAILURE, SUCCES, ERROR``` приведет к ошибке. Но если он ```NEW``` то его можно отменить и статус будет ```ABORT```
```
POST   /api/pay_system/edit_pay  
body 
{
    "id_transaction":2,
    "status":"FAILURE"
}
```
```
POST   /api/check_pay
body 
{
    "id_transaction": 2
}
```
```
POST   /api/check_pay_userid
body
{
    "id_user": 1
}
```
```
POST   /api/check_pay_email
body
{
    "email": "test@test.co"
}
```
```
POST   /api/generate_token
body
{
    "pay_system_id":124
}       
```
