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
        "ID": 1,
        "AuthorID": 1,
        "Title": "Article 1",
        "Content": "Lorem Ipsum",
        "CreationDate": "Thu, 01 Apr 2021 17:47:51",
        "LatestUpdate": "Thu, 01 Apr 2021 17:47:51"
    },
    {
        "ID": 2,
        "AuthorID": 1,
        "Title": "Article 2",
        "Content": "Lorem Ipsum",
        "CreationDate": "Thu, 01 Apr 2021 13:23:12",
        "LatestUpdate": "Thu, 01 Apr 2021 17:47:51"
    }
]
```


