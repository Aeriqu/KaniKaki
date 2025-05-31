# Elevating someone to admin

This is currently done through connecting directly to the mongodb instance and running queries to change it.

```
use auth
db.users.updateOne(
  { username: "usernameToElevate" },
  { $set:
    {
      type: 1
    }
  }
)
```

For local development, you can get shell access to the container and then run the following:

```
mongosh --username mongodb-auth-test-username
use auth
db.users.updateOne({username: "testUser" }, {$set:{type: 1}})
```