# Todo List备忘录
此项目使用gin+gorm实现的一个纯后端备忘录
涉及工具：vscode+ApiPost+Navicat   其他依赖：ini+mysql
· 用户注册登录，使用bcrypt对用户密码进行加密存储以及登录是密码验证，用户注册成功后给该用户分发一个token，当对备忘录进行操作时token验证成功才能进行相关操作；
· 使用ApiPost工具对备忘录进行新增/删除/修改/查询，后端用gorm的Create、First等方法来对mysql数据库进行实现，存储每条备忘录的信息；
· 路由分发:创建gin实例后进行相应模块的路由分发
· 分页查询功能
