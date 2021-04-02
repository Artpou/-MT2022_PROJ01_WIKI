# Login

Used to collect a Token for a registered User.

**URL** : `/api/login/`

**Method** : `POST`

**Auth required** : NO

**Data constraints**

```json
{
    "Username": "[valid username]",
    "Password": "[password in plain text]"
}
```
If you use postman, you must complete the request in json in body>raw


## Success Response

**Condition** : If 'username' and 'password' combination is correct.

**Code** : `200 OK`

**Content example**

```json
{
    "Name": "token",
    "Value": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NSwidXNlcm5hbWUiOiJnbWFyY2VsIiwicm9sZSI6MCwiZXhwIjoxNjE3MzgyNjY3fQ.2ziZ5UH8tU8sftYy1yKQ81Q4kFswKIx0vSKVm8GNGKo",
    "Path": "",
    "Domain": "",
    "Expires": "2021-04-02T18:57:47.470642+02:00",
    "RawExpires": "",
    "MaxAge": 0,
    "Secure": false,
    "HttpOnly": false,
    "SameSite": 0,
    "Raw": "",
    "Unparsed": null
}
```

## Error Response

**Condition** : If 'username' and 'password' combination is wrong.

**Code** : `400 BAD REQUEST`

**Content** :

```json
{
    "non_field_errors": [
        "Unable to login with provided credentials."
    ]
}
```
