# Statistics Service for PancakeSwap

## Start

-  Install golang, require version >= 1.13.
-  run `make build &&  nohup ./build/pancake-statas --config-path config/config.json &`

or simple use `service stat start` or `service stat restart`

## How it works

All price is deduced from chain, the price info may not accurate when liquidity is bad.

## Database
Restart docker after reboo.
`docker run -p 3307:3306 --name stat-mysql -v /opt/sql-datadir:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=123123 -d mysql`

## Endpoints

- 127.0.0.1:8080/api/v1/stat
- 127.0.0.1:8080/api/v1/price
- 127.0.0.1:8080/api/v1/syrup

## WorkSpace 
`/home/ubuntu/stats`
