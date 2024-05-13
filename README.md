# GO-DISTRIBUTED

## 项目结构

    $ go mod tidy
    D:\Distributed-Simple
    ├─cmd  //项目启动入口
    ├─grades //demo学生成绩服务模块
    ├─log  //log日志模块
    ├─registry  //服务注册模块
    ├─service  //服务启动模块
    └service  //服务启动模块

## 思维导图

```mermaid
graph LR;
    A(Client) <==> B[WEB Protal Service]
    B <==> C[Registry Service];
    B --> D[Log Service];
    B --> E[Grading Service];
    C <==> D;
    C <==> E;
    E --> D;

    style A fill:#f9f,stroke:#333,stroke-width:4px,font-size:32px,color:#000,font-weight:bold,left:100px,top:100px;
    style B fill:#bbf,stroke:#333,stroke-width:4px,font-size:24px,color:#000,font-weight:bold,left:300px,top:100px;
    style C fill:#bfb,stroke:#333,stroke-width:4px,font-size:16px,color:#000,font-weight:bold,left:500px,top:100px;
    style D fill:#fbf,stroke:#333,stroke-width:4px,font-size:100px,color:#000,font-weight:bold,left:700px,top:300px
    style E fill:#bff,stroke:#333,stroke-width:4px,font-size:12px,color:#000,font-weight:bold,left:700px,top:300px
```