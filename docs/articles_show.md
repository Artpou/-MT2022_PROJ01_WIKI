# Show Articles

Used to show all articles.

**URL** : `/api/articles/`

**Method** : `GET`

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
            "date_article": "2020-10-12 11:11:11",
            "id_user": 11
        }
    },
    {
        "article": {
            "id": 124,
            "title_article": "I am the second article",
            "content_article": "This is my content",
            "date_article": "2020-11-12 09:09:09",
            "id_user": 18
        }
    },
    {
        "article": {
            "id": 125,
            "title_article": "I am the third article",
            "content_article": "This is my content",
            "date_article": "2020-12-12 10:10:10",
            "id_user": 18
        }
    }
]
```


