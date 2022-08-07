## 生成模型
``go run main.go make model category``
## 生成迁移文件
``go run main.go make migration add_categories_table``
## 执行迁移
``go run main.go migrate up``
## 生成request文件
``go run main.go make request category``
## 生成controller文件
``go run main.go make apicontroller v1/category``

