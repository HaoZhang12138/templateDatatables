#templateDatatables
datatables and editor to display table and how to handle it in the back-end

使用datatables和editor在web前端显示表格，和相应的web后端程序处理请求，使用下面的步骤去显示你自定义的表。

#### 不上传文件的表
1. 创建新的go文件，用于自定义结构体，实现相应的GetId方法，注意把主键的bson别名定义为`_id`
2. 修改datatablesDaofactory.go，加入相应的case
3. 创建js文件，以tables.js为模板，要修改的地方主要是editor中的idSrc属性和fields属性，以及table中的columns属性
4. 创建html文件，以view.html为模板，定义显示的内容和表名，注意显示的内容顺序要和columns属性的顺序相同

#### 上传文件的表
- 步骤和不上传文件的表类似，不同之处有：
1. 自定义结构体里加入
		FileId string `json:"fileid" json:"fileid"`
2. 根据tables.js文件在自定义的js文件的fields属性和columns属性中加入相应的upload域

#### 创建自己的handler
- 上面使用的都是默认的handler在后台对前端的请求进行处理，如果你对数据的处理有什么特殊要求，可自定义handler，步骤如下：
1. 创建go文件，用于自定义handler， 实现DataTableHandler接口里的方法，例子可看petshandler.go
2. 修改datatablesfactory.go，加入相应的case
3. 修改js文件ajax的url值， 例子看pets.js


#### 相应的配置
1. 连接mongo数据库的信息在databaseConnect.go文件
2. 上传文件的配置信息在datatablesDao_uploadfile.go文件





