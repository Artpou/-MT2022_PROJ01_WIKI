# Show Articles

Used to show all comments for an article.

**URL** : `/api/comments/`

**Method** : `GET`

**Auth required** : NO

**Data constraints** :

{
    "id_article": "[id of the article you want the comments from]"
}

## Success Response

**Condition** : If article is not missing and exists.

**Code** : `200 OK`

**Content Example** :

```json
[
    {
        "comment": {
            "id": 53,
            "content_comment": "This is the content of the first comment",
            "date_comment": "2020-10-12 11:11:11",
            "id_user": 15,
            "id_article" : 123
        }
    },
    {
        "comment": {
            "id": 54,
            "content_comment": "This is the content of the second comment",
            "date_comment": "2020-11-11 10:11:12",
            "id_user": 17,
            "id_article" : 124
        }
    },
    {
        "comment": {
            "id": 55,
            "content_comment": "This is the content of the third comment",
            "date_comment": "2021-02-03 11:10:09",
            "id_user": 15,
            "id_article" : 124
        }
    }
]
```
## Error Response

**Condition** : If article does not exist.

**Code** : `404 NOT FOUND`

**Content** : {}
