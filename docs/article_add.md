# Create Article

Used to create an article.

**URL** : `/api/articles/`

**Method** : `POST`

**Auth required** : YES

**Data constraints**

```json
{
    "Title": "[title of the article]",
    "Content": "[content of the article]"
}
```

## Success Response

**Condition** : If user is logged in and no fields are missing.

**Code** : `201 CREATED`

## Error Response

**Condition** : If any field is missed.

**Code** : `400 BAD REQUEST`

**Content example**

```json
{
    "error": "The field 'Content' is missing"
}
```
### Or

**Condition** : If authentication failed.

**Code** : `403 FORBIDDEN`

**Content** : 

```json
{
    "error": "you need to be authenticated to do this"
}
```