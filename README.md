### 配置文件

#### 开发

首先要修改 `config.example.yaml` 为 `config.yaml` 这是开发

启动 `go run main.go` 

监听端口 `8000`


#### 生产

新增加文件 `config.yaml`

### 实现的功能

* [x] 实现Restful风格的接口

* [x] jwt的校验

* [x] validate.v9参数中文校验


### todo

* [ ] 报错信息

* [ ] 完善的log功能

* [ ] docker 启动部署

* [ ] 数据库自动创建

* [ ] 校验的完成报错，提取到公共包

* [ ] 完成注释

* [ ] 命名优化

* [ ] 生成swagger

* [ ] redis引入