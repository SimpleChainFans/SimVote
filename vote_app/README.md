
#### 生成接口文档
swag init --generalInfo cmd/vote/main.go --output docs/

#### gorm-tools工具

```bash
1. mkdir gormstructs && cd gormstructs

2. Download and Install gorm:
    wget https://github.com/xxjwxc/gormt/releases/download/v2.0.0/gormt_linux.zip && unzip gormt_linux.zip
    @See https://github.com/xxjwxc/gormt/releases
    
3. vi gormstructs/config.yaml
    
4. 导出和导入表结构
$ mysqldump -h192.168.3.212 -uroot -p123456 -d exchange trans_half_topic trans_message > trans.sql
$ cat trans.sql | mysql -h192.168.3.212 -uroot -p123456 gormstructs

5. Mysql database to golang gorm struct:
    gormstructs/gormt
    @See:https://xxjwxc.github.io/post/gormtools/
```

#### abi文件
``` bash
$ abigen --abi=foundation.abi --bin=foundation.bin --pkg=foundation --out=foundation/foundation.go
```
