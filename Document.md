### This simple API performs CRUD functionality on a person's pass as a request.

Create Name 
- Request
```json
{
    "name":"Peace",
    "track":"Marketing"
}
```
- Response
```json
{
    "id": 1,
    "name": "Peace",
    "track": "Marketing"
}
```

Get Name
- Request is a path parameter

- Response
  ```json
  {
    "ID": 1,
    "CreatedAt": "2023-09-14T18:49:45.935058+01:00",
    "UpdatedAt": "2023-09-14T18:49:45.935058+01:00",
    "DeletedAt": null,
    "name": "Peace",
    "track": "Marketing"
}
  ```

