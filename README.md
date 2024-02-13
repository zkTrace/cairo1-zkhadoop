# Go Server

Based heavily off MIT Lab for Computer Science's 6.824 Distributed Systems course. Go wrapper to facilitate Cairo workers for zk MapReduce.

## Testing Map Reduce

Use these tests to verify that the map reduce framework is working correctly. Lowkey will need to change these after get Cairo integrated

```sh
cd src/main
go build -buildmode=plugin ../mrapps/wc.go
bash test-mr.sh
```

## Todo

[] integrate cairo