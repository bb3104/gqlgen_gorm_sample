# gqlgen_gorm_sample

## create tables and init data
```
go run lib/init.go
```

## Graphql Query Get Users data
http://localhost:8080/playground

```
query users {
  users {
    id
    name
    email
    items {
      id
      name
    }
  }
}
```

```
{
  "data": {
    "users": [
      {
        "id": 1,
        "name": "hogehoge",
        "email": "hogehoge@gmail.com",
        "items": [
          {
            "id": 1,
            "name": "item"
          }
        ]
      }
    ]
  }
}
```
