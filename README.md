# go_study
go言語の勉強用

## Docker set up

```
$ cd /path/to/go_study/
$ docker-compose build
$ docker-compose up -d
```

## ローカル環境構築

### exce rovisioning

```
$ cd /path/to/go_study/
$ ansible-playbook -i ansible/inventories/local site.yml
```

### docker exec web container

```
$ cd /path/to/go_study/
$ docker exec -it go_study /bin/bash --login
```
