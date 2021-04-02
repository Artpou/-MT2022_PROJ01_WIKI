# Update Article

Used to update an article.

**URL** : `/api/articles/:id_article`

**Method** : `PUT`

**Auth required** : YES

**Data constraints** : 

```json
{
    "Title": "[title of the article]",
    "Content": "[content of the article]"
}
```

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

```json
{
    "error": "you need to be authenticated to do this"
}
```

## Or 

**Condition** : If any field is missing.

**Code** : `400 BAD REQUEST`

**Content** :

```json
{
    "error": "The field 'Content' is missing"
}
```
