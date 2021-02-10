# Create User

Used to create/register a User.

**URL** : `/api/users/`

**Method** : `POST`

**Auth required** : NO

**Data constraints**

```json
{
    "username": "[valid email address or username]",
    "password": "[password in plain text]"
}
```

## Success Response

**Code** : `201 CREATED`

## Error Response

**Condition** : If 'username' is already taken, or password length is inferior to 6, or any field is missing.

**Code** : `400 BAD REQUEST`
