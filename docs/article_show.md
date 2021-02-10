# Show Article

Used to show an article.

**URL** : `/api/articles/:id_article`

**Method** : `GET`

**Auth required** : NO

**Data constraints** : {}

## Success Response

**Condition** : If article exists.

**Code** : `200 OK`

**Content Example** :

```json
    {
        "article": {
            "id": 123,
            "title_article": "I am the article",
            "content_article": "This is my content",
            "id_user": 11
        }
    }
```

## Error Response :

**Condition** : If article does not exist.

**Code** : `404 NOT FOUND`

**Content** : `{}`
