# Article

Used to show all articles.

**URL** : `/api/articles/`

**Method** : `POST`

**Auth required** : NO

**Data constraints** : {}

## Success Response

**Condition** : None.

**Code** : `200 OK`

**Content Example** :

```json
[
    {
        "article": {
            "id": 123,
            "title_article": "I am the first article",
            "content_article": "This is my content",
            "id_user": 11
        }
    },
    {
        "article": {
            "id": 124,
            "title_article": "I am the second article",
            "content_article": "This is my content",
            "id_user": 18
        }
    },
    {
        "article": {
            "id": 125,
            "title_article": "I am the third article",
            "content_article": "This is my content",
            "id_user": 18
        }
    }
]
```


