# Delete Article

Used to delete an article.

**URL** : `/api/articles/:id_article`

**Method** : `DELETE`

**Auth required** : YES

**Data constraints** : {}

**Permissions required** :

User is in relation to the article requested:

* Owner `OO`

## Success Response

**Condition** : If article exists and current user is the owner of the article.

**Code** : `200 OK`

## Error Response

**Condition** : If article does not exist.

**Code** : `404 NOT FOUND`

**Content** : 

```json
    {
        "error": "This Article cannot be found"
    }
```

## Or

**Condition** : If article exists but Authorized User does not have required permissions.

**Code** : `403 FORBIDDEN`

***Content***

```json
    {
        "error": "you need to be authenticated to do this"
    }
```
