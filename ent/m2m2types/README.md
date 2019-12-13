# Groups-Users M2M Relation

In this groups-users example, we have a M2M relation between groups and their users.
Each group **has many** users, and each user can be joined to **many** groups.

### Generate Assets

```console
go generate ./...
```

### Run Example

```console
go test
```

### sql

```go
// Traverse the graph.
users, err := a8m.
    QueryGroups().                                           // [hub, lab]
    Where(group.Not(group.HasUsersWith(user.Name("nati")))). // [lab]
    QueryUsers().                                            // [a8m]
    QueryGroups().                                           // [hub, lab]
    QueryUsers().                                            // [a8m, nati]
    All(ctx)
```

```sql
SELECT DISTINCT "users"."id", "users"."age", "users"."name" FROM "users" 
    JOIN (SELECT "group_users"."user_id" FROM "group_users" 
        JOIN (SELECT "groups"."id" FROM "groups" 
            JOIN (SELECT "group_users"."group_id" FROM "group_users" 
                JOIN (SELECT "users"."id" FROM "users" 
                    JOIN (SELECT "group_users"."user_id" FROM "group_users" 
                        JOIN (SELECT "groups"."id" FROM "groups" 
                            JOIN (SELECT "group_users"."group_id" FROM "group_users" WHERE "group_users"."user_id" = $1) AS "t1" 
                            ON "groups"."id" = "t1"."group_id" WHERE NOT ("groups"."id" IN (SELECT "group_users"."group_id" FROM "group_users" 
                                JOIN "users" AS "t0" 
                                ON "group_users"."user_id" = "t0"."id" WHERE "t0"."name" = $2))) AS "t1" 
                        ON "group_users"."group_id" = "t1"."id") AS "t1" 
                    ON "users"."id" = "t1"."user_id") AS "t1" 
                ON "group_users"."user_id" = "t1"."id") AS "t1" 
            ON "groups"."id" = "t1"."group_id") AS "t1" 
        ON "group_users"."group_id" = "t1"."id") AS "t1" 
    ON "users"."id" = "t1"."user_id" args=[1 nati]
```