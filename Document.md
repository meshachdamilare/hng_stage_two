### This API handles Create, Read, Update, and Delete operations related to a person's name through incoming requests

Create Name 
- Request Body
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
- Request => With a path parameter in the URL e.g (localhost:4000/api/1)

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

Update Name 
- Request => With a path parameter in the URL e.g (localhost:4000/api/1)
- Request Body
```json
{
    "name":"Johnson"
}
```
- Response
```json
{
    "ID": 1,
    "CreatedAt": "2023-09-14T18:49:45.935058+01:00",
    "UpdatedAt": "2023-09-14T19:14:03.772178+01:00",
    "DeletedAt": null,
    "name": "Johnson",
    "track": "Marketing"
}
```

Delete Name
- Request => With a path parameter in the URL e.g (localhost:4000/api/1)

- Response
- Status Code 204 
  ```json
  {}
  ```

