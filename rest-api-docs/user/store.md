# Store User

Create User

**URL** : `/api/users/`

**Method** : `POST`

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

**Data examples**

```json
{
	"email": "test@test.com",
	"password": "password",
	"first_name": "fname_test",
	"last_name": "lname_test"
}
```

## Success Response

**Code** : `201 Created`

**Content examples**

```json
{
	"id": 4,
	"email": "test@test.com",
	"password": "password",
	"first_name": "fname_test",
	"last_name": "lname_test"
}
```

## Error Response

**Code** : `500 Internal Server Error`

**Content examples**

```json
{
    "error": "mail: missing '@' or angle-addr"
}
```
