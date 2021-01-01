[TOC]

> 以下为[左耳听风](https://time.geekbang.org/column/intro/48) 记录的笔记。

# 分布式系统架构



## 为什么需要分布式系统

放弃传统的单体架构，选择分布式系统，主要有两方面原因：

- 增大系统容量。

  当业务量越来越大，一台机器的性能无法满足，需要多台机器才能应对大规模应用场景，所以需要垂直或者水平拆分业务系统，使其变成分布式架构。

- 加强系统可用。

  业务越来越关键，需要提高整个系统架构的可用性，意味着架构中不能存在单点故障，为避免一台机器故障导致整体不可用，所以需要通过分布式架构来冗余系统以消除单点故障。

> 还有其他优势：
>
> - 模块化，系统模块重用度高；
> - 软件服务模块被拆分，开发和发布速度可以并行而变得更快；
> - 系统扩展性更高；
> - 团队协作流程得到改善等。

![image-20201231201827935](http://img.zhengyua.cn/img/20201231201835.png)

## 分布式系统的发展

开发、维护和使用 SOA 要遵循以下几条基本原则：

- 可重用，粒度合适，模块化，可组合，构件化以及有互操作性；
- 符合开放标准（通用的或行业的）；
- 服务的识别和分类，提供和发布，监控和跟踪。

SOA架构的演化图：

![image-20201231202201323](http://img.zhengyua.cn/img/20201231202201.png)

- 单体架构，软件模块高度耦合；
- 通过标准的协议或是中间件来联动其它相关联的服务来实现松耦合，利用IOC（控制反转）和DIP（依赖倒置原则）设计思想进行实践；
- 业务拆分成多个微服务，微服务独立完整运行，服务间的整合需要服务编排或是服务整合的引擎，可以使工作流引擎，也可以是网关，也需要辅助于像容器化调度的技术方式来进行管理。

微服务的出现使得开发速度变得更快，部署快，隔离性高，系统的扩展度也很好，但是在集成测试、运维和服务管理等方面难度提升，所以需要一套较好的微服务PaaS平台。

## 分布式系统中需要注意的问题

### 异构系统的不标准问题

主要表现在：

- 软件和应用不标准；
- 通讯协议不标准；
- 数据格式不标准；
- 开发和运维的过程和方法不标准；

### 系统架构中的服务依赖性问题

服务依赖会导致的问题主要有以下场景：

- 如果非关键业务被关键业务所依赖，会导致非关键业务变成一个关键业务。
- 服务依赖链中，出现“木桶短板效应”——整个 SLA 由最差的那个服务所决定。

这就涉及到服务治理的内容。

> 不但要拆分服务，还要为每个服务拆分相应的数据库。

### 故障发生的概率更大

- 出现故障不可怕，故障恢复时间过长才可怕。
- 出现故障不可怕，故障影响面过大才可怕。

所谓 Design for Failure。在设计时就要考虑如何减轻故障。如果无法避免，也要使用自动化的方式恢复故障，减少故障影响面。

### 多层架构的运维复杂度更大

通常来说可以把系统分成四层：基础层、平台层、应用层和接入层。

- 基础层：机器、网络和存储设备等。
- 平台层：中间件层，Tomcat、MySQL、Redis、Kafka 之类的软件。
- 应用层：业务软件，比如，各种功能的服务。
- 接入层：接入用户请求的网关、负载均衡或是 CDN、DNS 类似。

需要注意：

- 任何一层的问题都会导致整体的问题；
- 没有统一的视图和管理，导致运维被割裂开来，造成更大的复杂度。