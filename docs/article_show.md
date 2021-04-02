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
        "ID": 1,
        "AuthorID": 1,
        "Title": "jk",
        "Content": "trutyru",
        "CreationDate": "Thu, 01 Apr 2021 17:47:51",
        "LatestUpdate": "Thu, 01 Apr 2021 17:47:51",
        "Comments": []
    }
```

## Error Response :

**Condition** : If article does not exist.

**Code** : `404 NOT FOUND`

**Content** : 

```json
    {
        "error": "This Article cannot be found"
    }
```