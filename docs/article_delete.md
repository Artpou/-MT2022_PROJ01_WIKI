# Delete Article

Used to delete an article.

**URL** : `/api/articles/:id_article`

**Method** : `DELETE`

**Auth required** : YES

**Data constraints** :

```json
{
    "id_user": "[current user id]"
}
```

## Success Response

**Condition** : If Article exists and current user is the author of the article.

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

**Condition** : If user is not the author of the article.

**Code** : `403 FORBIDDEN`

```json
{"detail": "You do not have permission to perform this action."}
```
