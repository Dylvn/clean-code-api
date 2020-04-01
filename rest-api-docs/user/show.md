# Show User

Show a specific user

**URL** : `/api/users/{user_id}`

**Method** : `GET`

**Auth required** : NO

**Permissions required** : None

## Success Response

**Code** : `200 Created`

**Content examples**

```json
{
    "id": 4,
    "email": "test@test.com",
    "password": "password",
    "first_name": "fname",
    "last_name": "lname"
}
```

## Error Response

**Code** : `500 Internal Server Error`

**Content examples**

```json
{
    "error": "user with id {user_id} not found"
}
```