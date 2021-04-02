# Create User

Used to create/register a User.

**URL** : `/api/users/`

**Method** : `POST`

**Auth required** : NO

**Data constraints**

```json
{
    "Username": "[valid  username]",
    "Password": "[password in plain text]"
}
```

## Success Response

**Condition** : If no fields are missing.

**Code** : `201 CREATED`

## Error Response

**Condition** : If any field is missing.

**Code** : `400 BAD REQUEST`

**Content example**

```json
{
    "error": "The field 'Username' is missing"
}
```

### Or


**Condition** : If username is already taken.

**Code** : `400 BAD REQUEST`
