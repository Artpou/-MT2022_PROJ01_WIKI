# Update Comment

Used to update a comment.

**URL** : `/api/comments/:id_comment`

**Method** : `PUT`

**Auth required** : YES

**Data constraints** : 

```json
{
    "content_comment": "[content of the comment]"
}
```

**Permissions required** :

User is in relation to the comment requested:

* Owner `OO`

## Success Response

**Condition** : If comment exists and current user is the owner of the article.

**Code** : `200 OK`

## Error Response

**Condition** : If comment does not exist.

**Code** : `404 NOT FOUND`

**Content** : {}

## Or

**Condition** : If comment exists but Authorized User does not have required permissions.

**Code** : `403 FORBIDDEN`

```json
{"detail": "You do not have permission to perform this action."}
```

## Or

**Condition** : If 'content_comment' field is missing.

**Code** : `400 BAD REQUEST`

```json
{
    "content_comment": [
        "This field is required."
    ]
}
```
