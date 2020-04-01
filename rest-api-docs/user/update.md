# Update User

Update a specific user.

**URL** : `/api/user/{user_id}`

**Method** : `PUT`

**Auth required** : NO

**Permissions required** : None

**Data constraints**

```json
{
    "email": "[1 to 255 chars]",
    "password": "[1 to 255 chars]",
    "first_name": "[1 to 255 chars]",
    "last_name": "[1 to 255 chars]"
}
```

Note that `id` are currently read only field.

**Data examples**

Partial data is allowed.

```json
{
    "first_name": "John"
}
```

## Success Responses

**Condition** : Data provided is valid and User is Authenticated.

**Code** : `200 OK`

**Content example**

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

**Condition** : If provided data is invalid, e.g. a name field is too long.

**Code** : `500 Internal Server Error`

**Content example** :

```json
{
    "error": "user with id {user_id} not found"
}
```
