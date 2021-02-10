# Create Article

Used to create a comment.

**URL** : `/api/comments/`

**Method** : `POST`

**Auth required** : YES

**Data constraints**

```json
{
    "content_comment": "This is the content of the comment",
    "date_comment": "2020-10-12 11:11:11",
    "id_user": 15,
    "id_article": 123
}
```

## Success Response

**Condition** : If user is logged in and no fields are missing.

**Code** : `201 CREATED`

## Error Response

**Condition** : If any field is missed.

**Code** : `400 BAD REQUEST`

**Content example**

```json
{
    "content_comment": [
        "This field is required."
    ]
}
```
### Or

**Condition** : If authentication failed.

**Code** : `403 FORBIDDEN`

**Content** : {}
