- Admin data
```shell
create /ecommerce
create /ecommerce/admin
create /ecommerce/admin/server_port
create /ecommerce/admin/server_port/server "ecommerce"
create /ecommerce/admin/server_port/port ":9000"
```

- Common data
```shell
create /ecommerce/common
create /ecommerce/common/mongodb
create /ecommerce/common/mongodb/uri "mongodb://localhost:27017"
create /ecommerce/common/mongodb/db_name "ecommerce"
```

```shell
create /ecommerce/common/auth
create /ecommerce/common/auth/secret_key "secret_key_admin"
```