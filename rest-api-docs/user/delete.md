# Show User

Show a specific user

**URL** : `/api/users/{user_id}`

**Method** : `DELETE`

**Auth required** : NO

**Permissions required** : None

## Success Response

**Code** : `200 Created`

**Content examples**

```json
{
    "status": "deleted"
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