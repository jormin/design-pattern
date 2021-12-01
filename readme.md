Design Pattern
=====

[![Build Status](https://github.com/jormin/design-pattern/workflows/test/badge.svg?branch=master)](https://github.com/jormin/design-pattern/actions?query=workflow%3Atest)
[![Codecov](https://codecov.io/gh/jormin/design-pattern/branch/master/graph/badge.svg?)](https://codecov.io/gh/jormin/design-pattern)
[![Go Report Card](https://goreportcard.com/badge/github.com/jormin/design-pattern)](https://goreportcard.com/report/github.com/jormin/design-pattern)
[![](https://img.shields.io/badge/version-v1.0.0-success.svg)](https://github.com/jormin/design-pattern)

Go 实现设计模式，集合简单示例进行说明验证。

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

| 序号 | 类型 | 模式         | 英文                     | 文档视频                                                | 完成 |
| ---- | ------------ | ------------------------ | ------------------------------------------------------------ | ---- | ---- |
| 01   | 创建型 | 工厂模式     | Factory Pattern          | [深入设计模式](https://refactoringguru.cn/design-patterns/factory-method) | ✅    |
| 02   | 创建型 | 抽象工厂模式 | Abstract Factory Pattern | [深入设计模式](https://refactoringguru.cn/design-patterns/abstract-factory) | ✅    |
| 03   | 创建型 | 单例模式     | Singleton Pattern        | [深入设计模式](https://refactoringguru.cn/design-patterns/singleton) | ✅    |
| 04   | 创建型 | 建造者模式   | Builder Pattern          | [深入设计模式](https://refactoringguru.cn/design-patterns/builder) | ✅ |
| 05 | 创建型 | 原型模式     | Prototype Pattern        | [深入设计模式](https://refactoringguru.cn/design-patterns/prototype) | ✅ |
| 06  | 结构型 | 适配器模式 | Adapter Pattern   | [深入设计模式](https://refactoringguru.cn/design-patterns/adapter) | ✅ |
| 07 | 结构型 | 桥接模式   | Bridge Pattern     | [深入设计模式](https://refactoringguru.cn/design-patterns/bridge) | ✅ |
| 08  | 结构型 | 组合模式   | Composite Pattern | [深入设计模式](https://refactoringguru.cn/design-patterns/composite) |      |
| 09  | 结构型 | 装饰器模式 | Decorator Pattern | [深入设计模式](https://refactoringguru.cn/design-patterns/decorator) |      |
| 10 | 结构型 | 外观模式   | Facade Pattern    | [深入设计模式](https://refactoringguru.cn/design-patterns/facade) |      |
| 11 | 结构型 | 享元模式   | Flyweight Pattern | [深入设计模式](https://refactoringguru.cn/design-patterns/flyweight) |      |
| 12 | 结构型 | 代理模式   | Proxy Pattern     | [深入设计模式](https://refactoringguru.cn/design-patterns/proxy) |      |
| 13 | 行为型 | 责任链模式 | Chain of Responsibility Pattern | [深入设计模式](https://refactoringguru.cn/design-patterns/chain-of-responsibility) |      |
| 14 | 行为型 | 命令模式   | Command Pattern                 | [深入设计模式](https://refactoringguru.cn/design-patterns/command) |      |
| 15 | 行为型 | 迭代器模式 | Iterator Pattern                | [深入设计模式](https://refactoringguru.cn/design-patterns/iterator) |      |
| 16 | 行为型 | 中介者模式 | Mediator Pattern                | [深入设计模式](https://refactoringguru.cn/design-patterns/mediator) |      |
| 17 | 行为型 | 备忘录模式 | Memento Pattern                 | [深入设计模式](https://refactoringguru.cn/design-patterns/memento) |      |
| 18 | 行为型 | 观察者模式 | Observer Pattern                | [深入设计模式](https://refactoringguru.cn/design-patterns/observer) |      |
| 19 | 行为型 | 状态模式   | State Pattern                   | [深入设计模式](https://refactoringguru.cn/design-patterns/state) |      |
| 20 | 行为型 | 策略模式   | Strategy Pattern                | [深入设计模式](https://refactoringguru.cn/design-patterns/strategy) |      |
| 21 | 行为型 | 模版模式   | Template Pattern                | [深入设计模式](https://refactoringguru.cn/design-patterns/template-method) |      |
| 22 | 行为型 | 访问者模式 | Visitor Pattern                 | [深入设计模式](https://refactoringguru.cn/design-patterns/visitor) |      |

### 参考
1. [深入设计模式](https://refactoringguru.cn/design-patterns)

License
-------

Under the [MIT](./LICENSE) License