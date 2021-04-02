
# Update Article

Used to update an User.

**URL** : `/api/users/:id_user`

**Method** : `PUT`

**Auth required** : YES

**Data constraints** : 

```json
{
    "Username": "[username of the user]",
    "Password": "[password of the user]"
}
```

## Success Response

**Condition** : If users exists and current user is administrator.

**Code** : `200 OK`

## Error Response

**Condition** : If user does not exist.

**Code** : `404 NOT FOUND`

**Content** : 
```json
    {
        "error": "This User cannot be found"
    }
```

## Or

**Condition** : If User exists but Authorized User does not have required permissions.

**Code** : `403 FORBIDDEN`

```json
{
    "error": "you need to be administrator to do this"
}
```

## Or 

**Condition** : If any field is missing.

**Code** : `400 BAD REQUEST`

**Content** :

```json
{
    "error": "The field 'Username' is missing"
}
```
