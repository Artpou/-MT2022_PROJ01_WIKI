# Show Comment

Used to show a comment.

**URL** : `/api/comments/:id_comment`

**Method** : `GET`

**Auth required** : NO

**Data constraints** : {}

## Success Response

**Condition** : If comment exists.

**Code** : `200 OK`

**Content Example** :

```json
    {
        "comment": {
            "id": 53,
            "content_comment": "This is the content of the comment",
            "date_comment": "2020-10-12 11:11:11",
            "id_user": 15,
            "id_article" : 123
        }
    }
```

## Error Response :

**Condition** : If comment does not exist.

**Code** : `404 NOT FOUND`

**Content** :

```json
    {
        "error": "This Comment cannot be found"
    }
```
