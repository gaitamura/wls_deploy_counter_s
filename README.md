# wls_deploy_counter

WebLogicの管理サーバログに記載されたデプロイメント・タスク 配布の行数を数える

管理対象サーバ名を手動で変更する必要がある。  

```
	host := []string{"managed_server1", "managed_server2", "managed_server3"}
```

## Usage

```
$ go run wls_deploy_counter.go <AdminServer.log>
```
