# douban-books-crawler
golang crawler demo

# Description
学习golang过程中写的一个很简单的爬虫,爬取豆瓣图书top250  

以后的人生中要好好读书



# Run  
运行前要在Mysql中建表，文件在
```go
/config/book.sql
```
建表后执行

```shell
$ go run main.go
```

# FrontEnd Analysis Tool

用前端写一个统计分析书籍数据的小工具
[vue-book-tool](https://github.com/fenghaojiang/vue-books-tool)  
后端用的gin实现,爬取豆瓣Top250书籍信息到mysql就可以跑起来  

