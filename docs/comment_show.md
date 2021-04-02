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
        "ID": 1,
        "AuthorID": 3,
        "ArticleID": 7,
        "Content": "Commentaire test 2",
        "CreationDate": "Mon, 01 Jan 0001 00:00:00",
        "LatestUpdate": "Mon, 01 Jan 0001 00:00:00"
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
