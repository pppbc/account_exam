#模拟用户模块


###列出指定工厂的员工列表


request
```
GET /plant/<:plantId>/staffs

```
success
```
HTTP1.1  200
Content-Type: application/json
{
    "status": 1,
    "desc": "succeed",
    "data": [
        {
            "id": 1,
            "plantId": 11111,
            "uid": 5,
            "name": "staff01",
            "sex": 1,
            "createdAt": "2019-12-19T17:50:23.493194+08:00",
            "updatedAt": "2019-12-19T17:50:23.493194+08:00",
            "avatar": "",
            "department": {
                "id": 16,
                "name": "department01",
            },
            "post": {
                "id": 1,
                "plantId": 11111,
                "departmentId": 16,
                "name": "post01",
            },
            "user": null
        },
        {
            ...
        }]
    }
```
###列出指定员工信息

request
```
GET /plant/<:plantId>/staffs/<:staffId>

```
success
```
HTTP1.1  200
Content-Type: application/json
{
    "status": 1,
    "desc": "succeed",
    "data": 
        {
            "id": 1,
            "plantId": 11111,
            "uid": 5,
            "name": "staff01",
            "sex": 1,
            "createdAt": "2019-12-19T17:50:23.493194+08:00",
            "updatedAt": "2019-12-19T17:50:23.493194+08:00",
            "avatar": "",
            "department": {
                "id": 16,
                "name": "department01",
            },
            "post": {
                "id": 1,
                "plantId": 11111,
                "departmentId": 16,
                "name": "post01",
            },
            "user": null
        }
    }
```
###添加一条员工记录

request

```
POST /plant/<:plantId>/staffs
{
    name:xxx
    sex :1
    ....
}
```

success

```
HTTP1.1  200
Content-Type: application/json
{
    "status": 1,
    "desc": "success"
}
```

###修改一条员工记录
request

```
PUT /plant/<:plantId>/staffs/<:staffId>
{
    name:xxx
    sex :1
    ....
}

```
success

```
HTTP1.1  200
Content-Type: application/json
{
    "status": 1,
    "desc": "success"
}
```

###删除一条员工记录

request

```
DELETE /plant/<:plantId>/staffs/<:staffId>
```
success

```
HTTP1.1  200
Content-Type: application/json
{
    "status": 1,
    "desc": "success"
}
```