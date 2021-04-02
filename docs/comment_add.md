# Create Comment

Used to create a comment.

**URL** : `/api/comments/`

**Method** : `POST`

**Auth required** : YES

**Data constraints** :

```json
{
    "ArticleID": 1,
    "Content": "This is the content of the comment"
}
```

## Success Response

**Condition** : If user is logged in and no fields are missing.

**Code** : `201 CREATED`

```json
{
    "ID": 2,
    "AuthorID": 1,
    "ArticleID": 2,
    "Content": "Test comment",
    "CreationDate": "Fri, 02 Apr 2021 18:14:28",
    "LatestUpdate": "Fri, 02 Apr 2021 18:14:28"
}
```

## Error Response

**Condition** : If any field is missed.

**Code** : `400 BAD REQUEST`

**Content example**

```json
{
    "error": "The field 'Content' is missing"
}
```
### Or

**Condition** : If authentication failed.

**Code** : `403 FORBIDDEN`

**Content** : 

```json
{
    "error": "you need to be authenticated to do this"
}
```
