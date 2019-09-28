To run the server on your system:

1. Make sure you have [dep](https://github.com/golang/dep) installed
2. Run `dep ensure` to install dependencies
3. Run `go build` to create the binary (`go_web_api`)
4. Run the binary : `./go_web_api`

Create the `candidates` table before running the application :

```sql
create table candidates (
candidate_id serial primary key,
candidate_name varchar(500),
candidate_phonenumber varchar(500),
status varchar (500)
);
```
