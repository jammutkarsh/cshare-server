# ⚠️ cShare Server

Not complety functional for testing puposes only.

_This project consumes port `5675` && `5432`_

## Project Setup

Download Repo

```bash
git clone https://github.com/JammUtkarsh/cshare-server
```

Run Docker compose to initialize DB and application.

```bash
docker compose up -d
```

Open PostgreSQL shell

```bash
docker exec -it postgrestest psql -U cshare
```

Now you can test the API's

## Functional API Endpoints

All the endpoints are up and running. They provide the basic correctness of operations.

There are still some caveats to this API.

- There is no auth system enabled here. If you know any username from the database, then you can retrieve their data. Though, there is no known method (up until now)  by which username can be extracted.

- The single clip delete operation doesn't update the clip_id in DB.
Example: By user 1, there are clips 1,2,3,4. User decides to delete clip 3. The order will be following 1,2,4.

- User login (Data in JSON required)
`{url}/login/`

```json
{
"username": "username",
"password": "password"
}
```

- User Signup (Data in JSON required)
`{url}/signup`

```json
{
"username": "username",
"password": "password"
}
```

- Insert a new Clip(Data in JSON required)
`{url}/postclip`

```json
{
"username": "username",
"message": "text message as long as you want",
"secret": false
}
```

- Get all clips
`{url}/getclip/:username`

- Get all clips from for a user
`{url}/getclip/:username/:clip_id`

- Update username of a user
`{url}/updateusername/:username/:updated`

- Delete a clip
`{url}/deleteclip/:username/:clip_id`

- Clear all your clip data
`{url}/deleteall/:username`

## Check the data in PostgreSQL shell

Check users

```sql
SELECT * FROM users;
```

Check clips added

```sql
SELECT * FROM clip_stack;
```
