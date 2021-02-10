# Article

Used to create an article.

**URL** : `/api/articles/`

**Method** : `POST`

**Auth required** : YES

**Data constraints**

```json
{
    "titre_article": "[title of the article]",
    "content_article": "[content of the article]"
    "id_user": "[id of the current user]"
}
```

## Success Response

**Code** : `201 CREATED`

## Error Response

**Condition** : If any field is missed

**Code** : `400 BAD REQUEST`
