# JWT Token Format

```json
{
  "alg": "HS512",
  "typ": "JWT"
}
{
  "exp": 1682598882,
  "sub": "testUser",
  "type": 0,
  "wanikani_level_limit": 3
}
```
* `exp`: Expiration of the jwt token
* `sub`: The user name of the user who owns the token
* `type`: The type value of the user
  * `0`: Regular user
  * `1`: Admin user
* `wanikani_level_limit`: Indicates what level of wanikani content the user can see up until