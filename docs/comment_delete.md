# Delete Comment

Used to delete a comment.

**URL** : `/api/comments/:id_comment`

**Method** : `DELETE`

**Auth required** : YES

**Data constraints** : {}

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
