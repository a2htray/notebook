Notebook: What We Need to Read
=========

> 中英文混用，书写方式比较随意，权当个人学习笔记

本仓库主要包括以下两个部分：

* 书中的代码示例
* 知识点总结

**快速导航**

Book: 《Hands-On RESTful Web Services With Go》

|示例|附加|
|:---|:---|
|查找最近的源|[code](./Hands-On%20RESTFul%20Web%20Services%20with%20Go/finding%20the%20fastest%20mirror%20site%20from%20a%20list/main.go) [note](https://github.com/a2htray/notebook/issues/1)|
|健康检查|[code](./Hands-On%20RESTFul%20Web%20Services%20with%20Go/health%20check/main.go)|
|使用 ServeMux 生成 UUID|[code](./Hands-On%20RESTFul%20Web%20Services%20with%20Go) [note](https://github.com/a2htray/notebook/issues/2)|
|生成随机整数和浮点数的 handlers (ServeMux)|[code](./Hands-On%20RESTFul%20Web%20Services%20with%20Go/multiple%20handlers%20with%20in-build%20ServeMux%20to%20random%20int%20&%20float/main.go) [note](https://github.com/a2htray/notebook/issues/3)|
|使用 httprouter 完成上述功能|[code](./Hands-On%20RESTFul%20Web%20Services%20with%20Go/develop%20two%20apis%20with%20httprouter/main.go)|
|静态文件|[code](./Hands-On%20RESTFul%20Web%20Services%20with%20Go/static%20files/main.go)|
|gorilla/mux - 路径参数匹配|[code](./Hands-On%20RESTFul%20Web%20Services%20with%20Go/gorilla-mux%20-%20path-based%20matching/main.go)|
|gorilla/mux - 查询参数匹配|[code](./Hands-On%20RESTFul%20Web%20Services%20with%20Go/gorilla-mux%20-%20query-based%20matching/main.go)|
|命名路由 gorilla-mux|[code](./Hands-On%20RESTFul%20Web%20Services%20with%20Go/named%20route/main.go)|
|闭包函数|[code](./Hands-On%20RESTFul%20Web%20Services%20with%20Go/a%20closure%20function%20returns%20another%20function/main.go)|
|中间件(net/http)|[code](./Hands-On%20RESTFul%20Web%20Services%20with%20Go/middleware%20with%20build-in%20net-http%20package/main.go)|
|链式中间件|[code](./Hands-On%20RESTFul%20Web%20Services%20with%20Go/multiple%20middlewares%20and%20channing/main.go)|
|链式中间件(Alice)|[code](./Hands-On%20RESTFul%20Web%20Services%20with%20Go/channing%20middlewares%20with%20Alice/main.go)|
|gorilla/handlers 包|[code](./Hands-On%20RESTFul%20Web%20Services%20with%20Go/gorilla-handlers%20package/main.go)|
|RPC 服务|[server code](./Hands-On%20RESTFul%20Web%20Services%20with%20Go/create%20an%20RPC%20server%20that%20returns%20the%20UTC%20server%20time/server.go) [client code](./Hands-On%20RESTFul%20Web%20Services%20with%20Go/create%20an%20RPC%20server%20that%20returns%20the%20UTC%20server%20time/client.go)|
|JSON-RPC(Gorilla)|[server code](./Hands-On%20RESTFul%20Web%20Services%20with%20Go/JSON-RPC%20using%20Gorilla%20RPC/server.go) [curl](./Hands-On%20RESTFul%20Web%20Services%20with%20Go/JSON-RPC%20using%20Gorilla%20RPC/request.sh)|
|go-restful 示例|[code](./Hands-On%20RESTFul%20Web%20Services%20with%20Go/go-restful%20-%20return%20the%20server%20time/main.go)|
|CURD operations with sqlite3 and go-sqlite3|[code](./Hands-On%20RESTFul%20Web%20Services%20with%20Go/CURD%20operations%20with%20sqlite3%20and%20go-sqlite3/main.go)|
|基于资源的 go-restful API 服务|[code](./Hands-On%20RESTFul%20Web%20Services%20with%20Go/a%20Metro%20Rail%20API%20with%20go-restful%20based%20on%20resource/main.go)|
|Gin 的小程序|[code](./Hands-On%20RESTFul%20Web%20Services%20with%20Go/a%20simple%20program%20with%20Gin/main.go)|
|Gin API 服务|[code](./Hands-On%20RESTFul%20Web%20Services%20with%20Go/Metro%20Rail%20API%20with%20Gin/main.go)|
|MongoDB 的基本操作|[code](./Hands-On%20RESTFul%20Web%20Services%20with%20Go/MongoDB%20Operations/main.go)|

Book：《深入浅出 React 和 Redux》

`原有代码中没有使用到 TypeScript，出于学习 TypeScript 的目的，示例代码中加入 TypeScript`

|示例|附加|
|:---|:---|
|显示点击次数的组件|[ClickCounter.tsx](./深入浅出%20React%20和%20Redux/tutorial/src/components/ClickCounter.tsx)|
|Props 和 State 组件属性的使用|[ControlPanel.tsx](./深入浅出%20React%20和%20Redux/tutorial/src/components/ControlPanel/index.tsx) [Counter.tsx](深入浅出%20React%20和%20Redux/tutorial/src/components/ControlPanel/Counter.tsx)|
|组件向外传递数据|[ControlPanel.tsx](./深入浅出%20React%20和%20Redux/tutorial/src/components/ControlPanel2/index.tsx) [Counter.tsx](./深入浅出%20React%20和%20Redux/tutorial/src/components/ControlPanel2/Counter.tsx)|
|基于 Redux 全局状态的组件|[动作类型](./深入浅出%20React%20和%20Redux/tutorial/src/components/ReduxControlPanel/ActionTypes.ts) [动作生成函数](./深入浅出%20React%20和%20Redux/tutorial/src/components/ReduxControlPanel/actions.ts) [reducer](./深入浅出%20React%20和%20Redux/tutorial/src/components/ReduxControlPanel/reducer.ts) [store](./深入浅出%20React%20和%20Redux/tutorial/src/components/ReduxControlPanel/store.ts) [Counter.tsx](./深入浅出%20React%20和%20Redux/tutorial/src/components/ReduxControlPanel/Counter.tsx) [ControlPanel.tsx](./深入浅出%20React%20和%20Redux/tutorial/src/components/ReduxControlPanel/index.tsx)|
|聪明组件与傻瓜组件(与教程不同)|[动作](./深入浅出%20React%20和%20Redux/tutorial/src/components/ReduxControlPanel2/actions.ts) [reducer](./深入浅出%20React%20和%20Redux/tutorial/src/components/ReduxControlPanel2/reducer.ts) [store](./深入浅出%20React%20和%20Redux/tutorial/src/components/ReduxControlPanel2/store.ts) [Counter.tsx](./深入浅出%20React%20和%20Redux/tutorial/src/components/ReduxControlPanel2/Counter.tsx) [ControlPanel.tsx](./深入浅出%20React%20和%20Redux/tutorial/src/components/ReduxControlPanel2/index.tsx)|
|聪明组件与傻瓜组件2(与教程不同)|[动作](./深入浅出%20React%20和%20Redux/tutorial/src/components/ReduxControlPanel3/actions.ts) [reducer](./深入浅出%20React%20和%20Redux/tutorial/src/components/ReduxControlPanel3/reducer.ts) [store](./深入浅出%20React%20和%20Redux/tutorial/src/components/ReduxControlPanel3/store.ts) [Counter.tsx](./深入浅出%20React%20和%20Redux/tutorial/src/components/ReduxControlPanel3/Counter.tsx) [ControlPanel.tsx](./深入浅出%20React%20和%20Redux/tutorial/src/components/ReduxControlPanel3/index.tsx)|
|Redux with Context|[StoreContext.tsx](./深入浅出%20React%20和%20Redux/tutorial/src/components/ReduxWithContext/StoreContext.tsx)|
|connect 方法的使用|[index.tsx](./深入浅出%20React%20和%20Redux/tutorial/src/components/ReduxConnect/index.tsx)|

Book：《深入理解 ES6》

* [var 声明及变量提升机制](./深入理解%20ES6/var%20声明及变量提升机制.md)
* [块级声明 - let 代替 var](./深入理解%20ES6/块级声明%20-%20let%20代替%20var.md)
* [禁止重声明](./深入理解%20ES6/禁止重声明.md)
* [使用 const 声明](./深入理解%20ES6/使用%20const%20声明.md)
* [用 const 声明对象](./深入理解%20ES6/用%20const%20声明对象.md)
* [临时死区 TDZ](./深入理解%20ES6/临时死区%20TDZ.md)
* [循环中的块作用域绑定](./深入理解%20ES6/循环中的块作用域绑定.md)
* [循环中的函数](./深入理解%20ES6/循环中的函数.md)
* [循环中使用立即调用函数](./深入理解%20ES6/循环中使用立即调用函数.md)
* [32 位检测函数](./深入理解%20ES6/32%20位检测函数.md)
* [字符的等效性](./深入理解%20ES6/字符的等效性.md)
* [正则 - 支持 Unicode 的 u 修饰符](./深入理解%20ES6/正则%20-%20支持%20Unicode%20的%20u%20修饰符.md)
* [字符串的子串识别](./深入理解%20ES6/字符串的子串识别.md)
* [正则表达式作为正则表达式构造方法的参数问题](./深入理解%20ES6/正则表达式作为正则表达式构造方法的参数问题.md)
* [正则表达式的字符串表示](./深入理解%20ES6/正则表达式的字符串表示.md)
* [模板字面量](./深入理解%20ES6/模板字面量.md)
* [函数的默认参数值](./深入理解%20ES6/函数的默认参数值.md)
* [默认参数值对 arguments 对象的影响](./深入理解%20ES6/默认参数值对%20arguments%20对象的影响.md)
* [处理无命名参数](./深入理解%20ES6/处理无命名参数.md)
* [增强的 Function 构造函数](./深入理解%20ES6/增强的%20Function%20构造函数.md)
* [不定参数与展开运算符](./深入理解%20ES6/不定参数与展开运算符.md)
* [函数的 name 属性](./深入理解%20ES6/函数的%20name%20属性.md)
* [函数的两个不同的方法](./深入理解%20ES6/函数的两个不同的方法.md)
* [函数的元属性 target](./深入理解%20ES6/函数的元属性%20target.md)
* [提升到顶部的块级函数](./深入理解%20ES6/提升到顶部的块级函数.md)
* [箭头函数的特性](./深入理解%20ES6/箭头函数的特性.md)
* [尾调用优化](./深入理解%20ES6/尾调用优化.md)
* [对象字面量属性方法的简写](./深入理解%20ES6/对象字面量属性方法的简写.md)
* [Object 的新增方法](./深入理解%20ES6/Object%20的新增方法.md)
* [自有属性枚举顺序的基本规则](./深入理解%20ES6/自有属性枚举顺序的基本规则.md)
* [Object.setPrototypeOf 设置原型](./深入理解%20ES6/Object.setPrototypeOf.md)
* [简化原型访问的  super 引用](./深入理解%20ES6/简化原型访问的%20%20super%20引用.md)
* [方法 supper 的调用过程](./深入理解%20ES6/方法%20supper%20的调用过程.md)

博文：“深入设计模式”

`基于 Go 的设计模式示例。设计模式这种东西，有了解且会用即可，不推荐强制使用。`

* 创建型模式
  * [抽象工厂](./Post%20Code/设计模式/创建型/抽象工厂/main.go) [UML](./Post%20Code/设计模式/创建型/抽象工厂/抽象工厂模式.png)
* 结构型模式
* 行为型模式

* stackoverflow
  * 后端
    * [GO - 如何在模板中遍历数组并使用下标](./stackoverflow/GO%20-%20如何在模板中遍历数组并使用下标.md)
    * [GO - 如何调试运行命令行 exit status 1](./stackoverflow/GO%20-%20如何调试运行命令行%20exit%20status%201.md)
    * [GO - 如何格式化 JSON](./stackoverflow/GO%20-%20如何格式化%20JSON.md)
  * 前端
    * [Web - What does enctype='multipart/form-data' mean?](./stackoverflow/form-enctype.md)
    * [Web - 在页面重定向时自定义设置请求头](./stackoverflow/在页面重定向时自定义设置请求头.md)
    * [JavaScript - 数组去重的多种实现](./stackoverflow/数组去重的多种实现.md)
    * [jQuery - 如何失效按钮](./stackoverflow/如何失效按钮.md)
    * [TS - Typescript: How to export a variable](./stackoverflow/ts-export-a-variable.md)
    * [TS - jQuery 中的 与 TypeScript 中的 string 类型不匹配问题](./stackoverflow/jquery-typescript-parameter-not-match.md)
    * [TS - 对象字面量隐式含 any 类型的读取](./stackoverflow/TS%20-%20对象字面量隐式含%20any%20类型的读取.md)
    * [TS - 对象可以为 null](./stackoverflow/)
    * [TS&React - select 的 onChange 事件的类型安全](./stackoverflow/select%20的%20onChange%20事件的类型安全.md)
    * [TS&React - TypeScript 中 React 如何初始化 State](./stackoverflow/TS%20-%20TypeScript%20中%20React%20如何初始化%20State.md)
    * [React - 如何将 React 的 Component 作为参数进行传递 - TypeScript](./stackoverflow/如何将%20React%20的%20Component%20作为参数进行传递%20-%20TypeScript.md)
    * [React - react-dom-router - Route 中 component 和 render 的区别](./stackoverflow/react-dom-router%20-%20Route%20中%20component%20和%20render%20的区别.md)
    * [React - FontAwesomeIcon 的 spinner 不 spin](./stackoverflow/Font%20awesome%20spinner%20not%20spinning.md)
    * [React - pre 标签格式化代码问题](./stackoverflow/数组去重的多种实现.md)
  * Database
    * [MySQL - 如何修改 max_allowed_packet 的大小](./stackoverflow/How%20to%20change%20max_allowed_packet%20size.md)
    * [MySQL - 随机选择行](./stackoverflow/Selecting%20Random%20Rows%20in%20MySQL.md)
  * Nginx
    * [如何个性 max_allowed_packet 的大小](./stackoverflow/How%20to%20change%20max_allowed_packet%20size.md)
  * Linux
    * [Linux - 单行文本转列](./stackoverflow/How%20to%20convert%20from%20row%20to%20column.md)
* 工作贴士
  * [Windows 下搭建 React 项目 - TypeScript](./work%20tips/Windows%20下搭建%20React%20项目%20-%20TypeScript.md)
  * [Mac 下修改 docker 的源 - 客户端](./work%20tips/Mac%20下修改%20docker%20的源%20-%20客户端.md)
  * [Mac 下清理 DNS 缓存](./work%20tips/Mac%20下清理%20DNS%20缓存.md)
  * [MySQL 的基本操作](./work%20tips/MySQL%20的基本操作.md)
  * [React 下 HTMLAnchorElement 的 onClick 函数声明](./work%20tips/React%20下%20HTMLAnchorElement%20的%20onClick%20函数声明.md)
  * [Ubuntu 下根据端口关闭进程](./work%20tips/ubuntu_close_thread_based_on_port.sh)
  * [Jetbrains 全家桶](./work%20tips/Jetbrains/README.md)
  * 通用模板
    * [.editorconfig](./work%20tips/通用模板/.editorconfig)
    * [tsconfig.json](./work%20tips/通用模板/tsconfig.json)
    * [rollup.config.js](./work%20tips/通用模板/rollup.config.js)
    * [Ubuntu 20.04 源 source.list](./work%20tips/通用模板/Ubuntu%2020.04%20源%20source.list)
* Linux Commands
  * [watch](./Linux%20Commands/watch.md)