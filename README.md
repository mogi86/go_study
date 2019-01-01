# go_study
go言語の勉強用

## Docker set up

```
$ cd /path/to/go_study/
$ docker-compose build
$ docker-compose up -d
```

## ローカル環境構築

### exce provisioning

```
$ cd /path/to/go_study/
$ ansible-playbook -i ansible/inventories/local site.yml
```

### docker exec web container

```
$ cd /path/to/go_study/
$ docker exec -it go_study /bin/bash --login
```

## AWS環境構築(prod)

### exce provisioning

* git上の以下のファイルの値をマスキングしているので、実行時に書き換える。
    * `ansible/inventories/prod/hosts`のホスト名(IPアドレス)
    * `ansible/roles/deploy/vars/main.yml`のgitのパスワード

```
$ cd /path/to/go_study/
$ ansible-playbook -i ansible/inventories/prod prod.yml
```
