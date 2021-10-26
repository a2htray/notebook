Notebook: What We Need to Read
=========

> 中英文混用，书写方式比较随意，权当个人学习笔记

本仓库主要包括以下两个部分：

* 书中的代码示例
* 知识点总结

**快速导航**

Book: Hands-On RESTful Web Services With Go

|demo|addition|
|:---|:---|
|finding the fastest mirror site from a list|[code](./Hands-On%20RESTFul%20Web%20Services%20with%20Go/finding%20the%20fastest%20mirror%20site%20from%20a%20list/main.go) [note](https://github.com/a2htray/notebook/issues/1)|
|health check|[code](./Hands-On%20RESTFul%20Web%20Services%20with%20Go/health%20check/main.go)|
|uuid generator api using ServeMux|[code](./Hands-On%20RESTFul%20Web%20Services%20with%20Go) [note](https://github.com/a2htray/notebook/issues/2)|
|multiple handlers with in-build ServeMux to random int & float|[code](./Hands-On%20RESTFul%20Web%20Services%20with%20Go/multiple%20handlers%20with%20in-build%20ServeMux%20to%20random%20int%20&%20float/main.go) [note](https://github.com/a2htray/notebook/issues/3)|
|develop two apis with httprouter|[code](./Hands-On%20RESTFul%20Web%20Services%20with%20Go/develop%20two%20apis%20with%20httprouter/main.go)|
|static files|[code](./Hands-On%20RESTFul%20Web%20Services%20with%20Go/static%20files/main.go)|
|gorilla/mux - path-based matching `TODO: 看源码了解匹配的过程`|[code](./Hands-On%20RESTFul%20Web%20Services%20with%20Go/gorilla-mux%20-%20path-based%20matching/main.go)|
|gorilla/mux - query-based matching|[code](./Hands-On%20RESTFul%20Web%20Services%20with%20Go/gorilla-mux%20-%20query-based%20matching/main.go)|
|named route|[code](./Hands-On%20RESTFul%20Web%20Services%20with%20Go/named%20route/main.go)|
|a closure function returns another function `闭包函数`|[code](./Hands-On%20RESTFul%20Web%20Services%20with%20Go/a%20closure%20function%20returns%20another%20function/main.go)|
|middleware with build-in net/http package|[code](./Hands-On%20RESTFul%20Web%20Services%20with%20Go/middleware%20with%20build-in%20net-http%20package/main.go)|
|multiple middlewares and channing|[code](./Hands-On%20RESTFul%20Web%20Services%20with%20Go/multiple%20middlewares%20and%20channing/main.go)|
|channing middlewares with Alice|[code](./Hands-On%20RESTFul%20Web%20Services%20with%20Go/channing%20middlewares%20with%20Alice/main.go)|
|gorilla/handlers package|[code](./Hands-On%20RESTFul%20Web%20Services%20with%20Go/gorilla-handlers%20package/main.go)|
|create an RPC server that returns the UTC server time|[server code](./Hands-On%20RESTFul%20Web%20Services%20with%20Go/create%20an%20RPC%20server%20that%20returns%20the%20UTC%20server%20time/server.go) [client code](./Hands-On%20RESTFul%20Web%20Services%20with%20Go/create%20an%20RPC%20server%20that%20returns%20the%20UTC%20server%20time/client.go)|
|JSON-RPC using Gorilla RPC|[server code](./Hands-On%20RESTFul%20Web%20Services%20with%20Go/JSON-RPC%20using%20Gorilla%20RPC/server.go) [curl](./Hands-On%20RESTFul%20Web%20Services%20with%20Go/JSON-RPC%20using%20Gorilla%20RPC/request.sh)|
|go-restful: return the server time|[code](./Hands-On%20RESTFul%20Web%20Services%20with%20Go/go-restful%20-%20return%20the%20server%20time/main.go)|
|CURD operations with sqlite3 and go-sqlite3|[code](./Hands-On%20RESTFul%20Web%20Services%20with%20Go/CURD%20operations%20with%20sqlite3%20and%20go-sqlite3/main.go)|
|a Metro Rail API with go-restful based on resource|[code](./Hands-On%20RESTFul%20Web%20Services%20with%20Go/a%20Metro%20Rail%20API%20with%20go-restful%20based%20on%20resource/main.go)|
|a simple program with Gin|[code](./Hands-On%20RESTFul%20Web%20Services%20with%20Go/a%20simple%20program%20with%20Gin/main.go)|
|Metro Rail API with Gin|[code](./Hands-On%20RESTFul%20Web%20Services%20with%20Go/Metro%20Rail%20API%20with%20Gin/main.go)|

Book：深入浅出 React 和 Redux

`原有代码中没有使用到 TypeScript，出于学习 TypeScript 的目的，示例代码中加入 TypeScript`

|demo|addition|
|:---|:---|
|显示点击次数的组件|[ClickCounter.tsx](./深入浅出%20React%20和%20Redux/tutorial/src/components/ClickCounter.tsx)|
|Props 和 State 组件属性的使用|[ControlPanel.tsx](./深入浅出%20React%20和%20Redux/tutorial/src/components/ControlPanel/index.tsx) [Counter.tsx](深入浅出%20React%20和%20Redux/tutorial/src/components/ControlPanel/Counter.tsx)|
|组件向外传递数据|[ControlPanel.tsx](./深入浅出%20React%20和%20Redux/tutorial/src/components/ControlPanel2/index.tsx) [Counter.tsx](./深入浅出%20React%20和%20Redux/tutorial/src/components/ControlPanel2/Counter.tsx)|
|基于 Redux 全局状态的组件|[动作类型](./深入浅出%20React%20和%20Redux/tutorial/src/components/ReduxControlPanel/ActionTypes.ts) [动作生成函数](./深入浅出%20React%20和%20Redux/tutorial/src/components/ReduxControlPanel/actions.ts) [reducer](./深入浅出%20React%20和%20Redux/tutorial/src/components/ReduxControlPanel/reducer.ts) [store](./深入浅出%20React%20和%20Redux/tutorial/src/components/ReduxControlPanel/store.ts) [Counter.tsx](./深入浅出%20React%20和%20Redux/tutorial/src/components/ReduxControlPanel/Counter.tsx) [ControlPanel.tsx](./深入浅出%20React%20和%20Redux/tutorial/src/components/ReduxControlPanel/index.tsx)|
|聪明组件与傻瓜组件(与教程不同)|[动作](./深入浅出%20React%20和%20Redux/tutorial/src/components/ReduxControlPanel2/actions.ts) [reducer](./深入浅出%20React%20和%20Redux/tutorial/src/components/ReduxControlPanel2/reducer.ts) [store](./深入浅出%20React%20和%20Redux/tutorial/src/components/ReduxControlPanel2/store.ts) [Counter.tsx](./深入浅出%20React%20和%20Redux/tutorial/src/components/ReduxControlPanel2/Counter.tsx) [ControlPanel.tsx](./深入浅出%20React%20和%20Redux/tutorial/src/components/ReduxControlPanel2/index.tsx)|
|聪明组件与傻瓜组件2(与教程不同)|[动作](./深入浅出%20React%20和%20Redux/tutorial/src/components/ReduxControlPanel3/actions.ts) [reducer](./深入浅出%20React%20和%20Redux/tutorial/src/components/ReduxControlPanel3/reducer.ts) [store](./深入浅出%20React%20和%20Redux/tutorial/src/components/ReduxControlPanel3/store.ts) [Counter.tsx](./深入浅出%20React%20和%20Redux/tutorial/src/components/ReduxControlPanel3/Counter.tsx) [ControlPanel.tsx](./深入浅出%20React%20和%20Redux/tutorial/src/components/ReduxControlPanel3/index.tsx)|
|Redux with Context|[StoreContext.tsx](./深入浅出%20React%20和%20Redux/tutorial/src/components/ReduxWithContext/StoreContext.tsx)|
|connect 方法的使用|[index.tsx](./深入浅出%20React%20和%20Redux/tutorial/src/components/ReduxConnect/index.tsx)|

* stackoverflow
  * [What does enctype='multipart/form-data' mean?](./stackoverflow/form-enctype.md)
  * [Typescript: How to export a variable](./stackoverflow/ts-export-a-variable.md)
  * [如何将 React 的 Component 作为参数进行传递 - TypeScript](./stackoverflow/如何将%20React%20的%20Component%20作为参数进行传递%20-%20TypeScript.md)
  * [react-dom-router - Route 中 component 和 render 的区别](./stackoverflow/react-dom-router%20-%20Route%20中%20component%20和%20render%20的区别.md)
  * [FontAwesomeIcon 的 spinner 不 spin](./stackoverflow/Font%20awesome%20spinner%20not%20spinning.md)
* work tips
  * [Windows 下搭建 React 项目 - TypeScript](./work%20tips/Windows%20下搭建%20React%20项目%20-%20TypeScript.md)
  * [Mac 下修改 docker 的源 - 客户端](./work%20tips/Mac%20下修改%20docker%20的源%20-%20客户端.md)
  * [Mac 下清理 DNS 缓存](./work%20tips/Mac%20下清理%20DNS%20缓存.md)
  * common templates
    * [.editorconfig](./work%20tips/通用模板/.editorconfig)
* Linux Commands
  * [watch](./Linux%20Commands/watch.md)
* useful packages
  * GO
    * github.com/felixge/httpsnoop `捕获与 http 相关的度量指标，如响应时间、状态码`
    * github.com/justinas/alice `链式处理中间件`
    * github.com/gorilla/handlers `预设的、常用的中间件`
    * github.com/gorilla/rpc `rpc 包`
    * github.com/emicklei/go-restful `快速开发面向资源的 restful api`
    * github.com/mattn/go-sqlite3 `sqlite3 驱动库`
    * github.com/gin-gonic/gin `正在用的框架`