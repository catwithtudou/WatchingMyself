# 设计模式-单例模式

> 对于系统中的某些类来说，只有一个实例很重要，例如，一个系统中可以存在多个打印任务，但是只能有一个正在工作的任务；一个系统只能有一个窗口管理器或文件系统；一个系统只能有一个计时工具或ID（序号）生成器。
>
> 如何保证一个类只有一个实例并且这个实例易于被访问呢？定义一个全局变量可以确保对象随时都可以被访问，但不能防止我们实例化多个对象。
>
> 一个更好的解决办法是让类自身负责保存它的唯一实例。这个类可以保证没有其他实例被创建，并且它可以提供一个访问该实例的方法。这就是单例模式的模式动机。

## 模式定义

**单例模式**:

**单例模式**，也叫**单子模式**，是一种常用的[软件设计模式](https://zh.wikipedia.org/wiki/软件设计模式)  。在应用这个模式时，单例对象的[类](https://zh.wikipedia.org/wiki/类_(计算机科学) ) 必须保证只有一个实例存在。许多时候整个系统只需要拥有一个的全局[对象](https://zh.wikipedia.org/wiki/对象) ，这样有利于我们协调系统整体的行为。比如在某个[服务器](https://zh.wikipedia.org/wiki/服务器) 程序中，该服务器的配置信息存放在一个[文件](https://zh.wikipedia.org/wiki/文件) 中，这些配置数据由一个单例对象统一读取，然后服务[进程](https://zh.wikipedia.org/wiki/进程) 中的其他对象再通过这个单例对象获取这些配置信息。这种方式简化了在复杂环境下的配置管理。

## 模式结构

![](https://raw.githubusercontent.com/catwithtudou/photo/master/20190813203946.png) 

- 单例模式就只有一个角色:

  Singleton:单例

  ![](https://raw.githubusercontent.com/catwithtudou/photo/master/20190813204024.png) 





## 代码实现

在单例模式实现的方法有非常多,我就列举出常用的几种

### 饿汉式

> 饿汉式单例在单例类被加载时候，就实例化一个对象交给自己的引用

```java
public class Singleton
{
	private static Singleton instance = new Singleton() ;
	
	public static Singleton getInstance() 
	{
		return instance ;
	}
}
```

### 懒汉式

> 懒汉式在调用取得实例方法的时候才会实例化对象

```java
public class Singleton{
    private static Singleton singleton;
    
    private Singleton() {}
    
    public static synchronized Singleton getInstance() {
        if (singleton==null) {
            singleton = new Singleton() ;
        }
        return singleton;
    }
}
/**
/这两种是一样的,都是采取双重锁来保证线程安全,只是写法不一样
public class Singleton
{
	private static Singleton instance;
	public static Singleton getInstance() 
	{
		if (instance == null) 
		{
			synchronized (Singleton.class) 
			{
				if (instance == null) 
				{
					instance = new Singleton() ;
				}
			}
		}
		return instance;
	}
}
**/
```

### 持有类

> 为了不在初始化的时候加载

```java
public class Singleton
{
	private static final class InstanceHolder
	{
		private static Singleton INSTANCE = new Singleton() ;
	}
 
	public static Singleton getInstance() 
	{
		return InstanceHolder.INSTANCE;
	}
}
```

### *枚举

这个方法我也没用过,感兴趣的小伙伴可以去搜索一下

```java
public enum Singleton
{
	INSTANCE;
}
```

### *go实现

go的实现方式有些特别,其实用的地方也有点少

```go
var _instance *service
var once sync.Once
type service struct {
	Name string
}

func Service()  *service {
	once.Do(func()  {
           _instance = &service{}
        }) 
	return _instance
}

func (this *service)  Startup()  {
}

func (this *service)  Stop()  {

}

func (this *service)  Restart()  {

}
```

## 模式分析

在单例模式的实现过程中,它有以下几个要素：

- 私有的构造方法
- 指向自己实例的私有静态引用
- 以自己实例为返回值的静态的公有的方法

## 模式优缺点

- 优点
    - 提供了对唯一实例的受控访问。因为单例类封装了它的唯一实例，所以它可以严格控制客户怎样以及何时访问它，并为设计及开发团队提供了共享的概念。
    - 由于在系统内存中只存在一个对象，因此可以节约系统资源，对于一些需要频繁创建和销毁的对象，单例模式无疑可以提高系统的性能。
    - 允许可变数目的实例。我们可以基于单例模式进行扩展，使用与单例控制相似的方法来获得指定个数的对象实例。
    - 可以全局访问

- 缺点
    - 由于单例模式中没有抽象层，因此单例类的扩展有很大的困难。
    - 单例类的职责过重，在一定程度上违背了“单一职责原则”。因为单例类既充当了工厂角色，提供了工厂方法，同时又充当了产品角色，包含一些业务方法，将产品的创建和产品的本身的功能融合到一起。
    - 滥用单例将带来一些负面问题，如为了节省资源将数据库连接池对象设计为单例类，可能会导致共享连接池对象的程序过多而出现连接池溢出；现在很多面向对象语言(如Java、C#) 的运行环境都提供了自动垃圾回收的技术，因此，如果实例化的对象长时间不被利用，系统会认为它是垃圾，会自动销毁并回收资源，下次利用时又将重新实例化，这将导致对象状态的丢失。

## 模式适用场景

- 需要频繁实例化然后要销毁的对象
- 创建对象时损耗的资源或时间过多,但又经常用到的对象
- 需要频繁访问数据库或文件的对象
- 有状态的工具类对象

> 在java中，饿汉式单例要优于懒汉式单例。C++中则一般使用懒汉式单例。

> 下面列举一个常用的实用场景
>
> 一个具有自动编号主键的表可以有多个用户同时使用，但数据库中只能有一个地方分配下一个主键编号，否则会出现主键重复，因此该主键编号生成器必须具备唯一性，可以通过单例模式来实现。

## 模式注意事项

- 只能使用单例类提供的方法得到单例对象，不要使用反射，否则将会实例化一个新对象。
- 不要做断开单例类对象与类中静态引用的危险操作。
- 多线程使用单例使用共享资源时，注意线程安全问题。

其中线程安全的问题,一般上面的方法都是线程安全的,除非有特殊情况,比如用反射产生新的单例对象就非常危险

> 在分布式系统、多个类加载器、以及序列化的的情况下，会产生多个单例，这一点是无庸置疑的。那么在同一个jvm中，会不会产生单例呢？使用单例提供的getInstance() 方法只能得到同一个单例，除非是使用反射方式，将会得到新的单例。代码如下
>
> ```java
> Class c = Class.forName(Singleton.class.getName() ) ;
> Constructor ct = c.getDeclaredConstructor() ;
> ct.setAccessible(true) ;
> Singleton singleton = (Singleton) ct.newInstance() ;
> ```
>
> 即每次运行都会产生新的单例对象。所以运用单例模式时，一定注意不要使用反射产生新的单例对象。

## 参考

https://blog.csdn.net/zhengzhb/article/details/7331369

https://design-patterns.readthedocs.io/zh_CN/latest/creational_patterns/singleton.html