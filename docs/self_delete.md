
# Delete Comment

Used to delete an user.

**URL** : `/api/users/:id_user`

**Method** : `DELETE`

**Auth required** : YES

**Data constraints** : {}

## Success Response

**Condition** : If user exists and current user is administrator.

**Code** : `200 OK`

## Error Response

**Condition** : If comment does not exist.

**Code** : `404 NOT FOUND`

**Content** :

```json
    {
        "error": "This User cannot be found"
    }
```

## Or

**Condition** : If comment exists but Authorized User does not have required permissions.

**Code** : `403 FORBIDDEN`

```json
{
    "error": "you need to be administrator to do this"
}
```
