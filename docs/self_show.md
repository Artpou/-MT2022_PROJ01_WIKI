
# Show Articles

Used to show all Users.

**URL** : `/api/users/`

**Method** : `GET`

**Auth required** : YES

**Data constraints** : {}

## Success Response

**Condition** : If current user is administrator.

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

## Error Response

**Condition** : If comment exists but Authorized User does not have required permissions.

**Code** : `403 FORBIDDEN`

```json
{
    "error": "you need to be administrator to do this"
}
```