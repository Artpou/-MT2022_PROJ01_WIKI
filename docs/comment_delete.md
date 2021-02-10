# Delete Article

Used to delete a comment.

**URL** : `/api/comments/:id_comment`

**Method** : `DELETE`

**Auth required** : YES

**Data constraints** : {}

**Permissions required** :

User is in relation to the Article requested:

* Owner `OO`

## Success Response

**Condition** : If comment exists and current user is the author of the comment.

**Code** : `200 OK`

## Error Response

**Condition** : If 'id_user' field is missing.

**Code** : `400 BAD REQUEST`

**Content** :

```json
{
    "id_user": [
        "This field is required."
    ]
}
```

## Or

**Condition** : If Article exists but Authorized User does not have required permissions.

**Code** : `403 FORBIDDEN`

```json
{"detail": "You do not have permission to perform this action."}
```
