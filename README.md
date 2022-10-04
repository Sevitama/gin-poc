# gin-poc

``` bash
go run .
```

## SQL Injection

Get Article with ID (secure): **Voting result**

Get Article with id (insecure!): **' OR 1=1 OR title = '**

Get Article with id (insecure!): **'; DROP TABLE article; -- ')**

Further information about SQL Injections: https://go.dev/doc/database/sql-injection

## XSS

Get Article with ID (secure): **<script>alert("hi")</script>**

Get Article with id (secure): **alert("hi")**

Get Article with id (insecure!): **<script>alert("hi")</script>**

Get Article with id (insecure!): **alert("hi")**

