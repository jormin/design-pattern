Design Pattern
=====

[![Build Status](https://github.com/jormin/design-pattern/workflows/test/badge.svg?branch=master)](https://github.com/jormin/design-pattern/actions?query=workflow%3Atest)
[![Codecov](https://codecov.io/gh/jormin/design-pattern/branch/master/graph/badge.svg?)](https://codecov.io/gh/jormin/design-pattern)
[![Go Report Card](https://goreportcard.com/badge/github.com/jormin/design-pattern)](https://goreportcard.com/report/github.com/jormin/design-pattern)
[![](https://img.shields.io/badge/version-v1.0.0-success.svg)](https://github.com/jormin/design-pattern)

Go 实现设计模式，集合简单示例及实际例子进行说明验证。

类型
-----

#### 创建型

创建型模式关注点是**如何创建对象**，其核心思想是要把对象的创建和使用相分离，这样使得两者能相对独立地变换。

#### 结构型

结构型模式主要涉及**如何组合各种对象**以便获得更好、更灵活的结构。虽然面向对象的继承机制提供了最基本的子类扩展父类的功能，但结构型模式不仅仅简单地使用继承，而更多地通过组合与运行期的动态组合来实现更灵活的功能。

#### 行为型

行为型模式主要涉及**算法和对象间的职责分配**。通过使用对象组合，行为型模式可以描述一组对象应该如何协作来完成一个整体任务。

常用设计模式
-----

| 序号 | 模式         | 英文                     | 文档视频                                                | Demo | 实例 |
| ---- | ------------ | ------------------------ | ------------------------------------------------------------ | ---- | ---- |
| 创建型 |  |  |  |  |  |
| 01   | 工厂模式     | Factory Pattern          | [廖雪峰](https://www.liaoxuefeng.com/wiki/1252599548343744/1281319170474017) [菜鸟](https://www.runoob.com/design-pattern/factory-pattern.html) [知乎](https://zhuanlan.zhihu.com/p/388067512) [BiliBili](https://www.bilibili.com/video/BV1G4411c7N4?p=39) | ✅    |      |
| 02   | 抽象工厂模式 | Abstract Factory Pattern | [廖雪峰](https://www.liaoxuefeng.com/wiki/1252599548343744/1281319134822433) [菜鸟](https://www.runoob.com/design-pattern/abstract-factory-pattern.html) [知乎](https://zhuanlan.zhihu.com/p/388067512) [BiliBili](https://www.bilibili.com/video/BV1G4411c7N4?p=45) | ✅    |      |
| 03   | 单例模式     | Singleton Pattern        | [廖雪峰](https://www.liaoxuefeng.com/wiki/1252599548343744/1281319214514210) [菜鸟](https://www.runoob.com/design-pattern/singleton-pattern.html) [知乎](https://zhuanlan.zhihu.com/p/387357546) [BiliBili](https://www.bilibili.com/video/BV1G4411c7N4?p=29) | ✅    |      |
| 04   | 建造者模式   | Builder Pattern          | [廖雪峰](https://www.liaoxuefeng.com/wiki/1252599548343744/1281319155793953) [菜鸟](https://www.runoob.com/design-pattern/builder-pattern.html) [知乎](https://zhuanlan.zhihu.com/p/388815746) [BiliBili](https://www.bilibili.com/video/BV1G4411c7N4?p=55) | ✅ |      |
| 05   | 原型模式     | Prototype Pattern        | [廖雪峰](https://www.liaoxuefeng.com/wiki/1252599548343744/1281319195639841) [菜鸟](https://www.runoob.com/design-pattern/prototype-pattern.html) [知乎](https://zhuanlan.zhihu.com/p/387707620) [BiliBili](https://www.bilibili.com/video/BV1G4411c7N4?p=49) | ✅ |      |
| 结构型 |  |  |  | | |
| 01   | 适配器模式 | Adapter Pattern   | [廖雪峰](https://www.liaoxuefeng.com/wiki/1252599548343744/1281319245971489) [菜鸟](https://www.runoob.com/design-pattern/adapter-pattern.html) [知乎](https://zhuanlan.zhihu.com/p/389509948) [BiliBili](https://www.bilibili.com/video/BV1G4411c7N4?p=60) | ✅ |      |
| 02   | 桥接模式   | Bridge Pattern     | [廖雪峰](https://www.liaoxuefeng.com/wiki/1252599548343744/1281319266943009) [菜鸟](https://www.runoob.com/design-pattern/bridge-pattern.html) |      |      |
| 03   | 组合模式   | Composite Pattern | [廖雪峰](https://www.liaoxuefeng.com/wiki/1252599548343744/1281319283720226) [菜鸟](https://www.runoob.com/design-pattern/composite-pattern.html) |      |      |
| 04   | 装饰器模式 | Decorator Pattern | [廖雪峰](https://www.liaoxuefeng.com/wiki/1252599548343744/1281319302594594) [菜鸟](https://www.runoob.com/design-pattern/decorator-pattern.html) |      |      |
| 05   | 外观模式   | Facade Pattern    | [廖雪峰](https://www.liaoxuefeng.com/wiki/1252599548343744/1281319346634785) [菜鸟](https://www.runoob.com/design-pattern/facade-pattern.html) |      |      |
| 06   | 享元模式   | Flyweight Pattern | [廖雪峰](https://www.liaoxuefeng.com/wiki/1252599548343744/1281319417937953) [菜鸟](https://www.runoob.com/design-pattern/flyweight-pattern.html) |      |      |
| 07   | 代理模式   | Proxy Pattern     | [廖雪峰](https://www.liaoxuefeng.com/wiki/1252599548343744/1281319432618017) [菜鸟](https://www.runoob.com/design-pattern/proxy-pattern.html) |      |      |
| 行为型 |  |  |  | | |
| 01   | 责任链模式 | Chain of Responsibility Pattern | [廖雪峰](https://www.liaoxuefeng.com/wiki/1252599548343744/1281319474561057) [菜鸟](https://www.runoob.com/design-pattern/chain-of-responsibility-pattern.html) |      |      |
| 02   | 命令模式   | Command Pattern                 | [廖雪峰](https://www.liaoxuefeng.com/wiki/1252599548343744/1281319491338273) [菜鸟](https://www.runoob.com/design-pattern/command-pattern.html) |      |      |
| 03   | 解释器模式 | Interpreter Pattern             | [廖雪峰](https://www.liaoxuefeng.com/wiki/1252599548343744/1281319508115489) [菜鸟](https://www.runoob.com/design-pattern/interpreter-pattern.html) |      |      |
| 04   | 迭代器模式 | Iterator Pattern                | [廖雪峰](https://www.liaoxuefeng.com/wiki/1252599548343744/1281319524892705) [菜鸟](https://www.runoob.com/design-pattern/iterator-pattern.html) |      |      |
| 05   | 中介者模式 | Mediator Pattern                | [廖雪峰](https://www.liaoxuefeng.com/wiki/1252599548343744/1281319541669922) [菜鸟](https://www.runoob.com/design-pattern/mediator-pattern.html) |      |      |
| 06   | 备忘录模式 | Memento Pattern                 | [廖雪峰](https://www.liaoxuefeng.com/wiki/1252599548343744/1281319562641441) [菜鸟](https://www.runoob.com/design-pattern/memento-pattern.html) |      |      |
| 07   | 观察者模式 | Observer Pattern                | [廖雪峰](https://www.liaoxuefeng.com/wiki/1252599548343744/1281319577321505) [菜鸟](https://www.runoob.com/design-pattern/observer-pattern.html) |      |      |
| 08   | 状态模式   | State Pattern                   | [廖雪峰](https://www.liaoxuefeng.com/wiki/1252599548343744/1281319592001569) [菜鸟](https://www.runoob.com/design-pattern/state-pattern.html) |      |      |
| 09   | 策略模式   | Strategy Pattern                | [廖雪峰](https://www.liaoxuefeng.com/wiki/1252599548343744/1281319606681634) [菜鸟](https://www.runoob.com/design-pattern/strategy-pattern.html) |      |      |
| 10   | 模版模式   | Template Pattern                | [廖雪峰](https://www.liaoxuefeng.com/wiki/1252599548343744/1281319636041762) [菜鸟](https://www.runoob.com/design-pattern/template-pattern.html) |      |      |
| 11   | 访问者模式 | Visitor Pattern                 | [廖雪峰](https://www.liaoxuefeng.com/wiki/1252599548343744/1281319659110433) [菜鸟](https://www.runoob.com/design-pattern/visitor-pattern.html) |      |      |

### 参考

1. [廖雪峰-设计模式](https://www.liaoxuefeng.com/wiki/1252599548343744/1264742167474528)
2. [菜鸟-设计模式](https://www.runoob.com/design-pattern/design-pattern-tutorial.html)
3. [知乎-Golang设计模式](https://www.zhihu.com/column/c_1393206420800598016)
4. [BiliBili-尚硅谷设计模式](https://www.bilibili.com/video/BV1G4411c7N4)

License
-------

Under the [MIT](./LICENSE) License