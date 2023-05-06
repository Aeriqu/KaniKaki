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

```
db.users.updateOne({username: "usernameToElevate" },{$set:{type: 1}})
```