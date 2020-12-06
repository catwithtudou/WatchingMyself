# 设计模式-原型模式

## 模式定义

**原型模式**是[创建型模式](https://zh.wikipedia.org/wiki/創建型模式) 的一种，其特点在于通过“复制”一个已经存在的[实例](https://zh.wikipedia.org/wiki/实例) 来返回新的实例,而不是新建实例。被复制的实例就是我们所称的“原型”，这个原型是可定制的。

原型模式多用于创建复杂的或者耗时的实例，因为这种情况下，复制一个已经存在的实例使程序运行更高效；或者创建值相等，只是命名不一样的同类数据。

> 以上摘自维基百科

## 模式结构

![](https://raw.githubusercontent.com/catwithtudou/photo/master/600px-Prototype_UML.svg.png)

- 原型模式主要用于对象复制

  其核心类就是类图中的Prototype,它需要具备以下两个条件

    - `实现Cloneable接口`:我们知道在java中有一个Cloneable接口,此接口的功能就是运行时通知虚拟机可以安全地在实现了此接口的类上使用clone方法。

      > 在java虚拟机中，只有实现了这个接口的类才可以被拷贝，否则在运行时会抛出CloneNotSupportedException异常。

    - `重写Object类中的clone方法`:在java中所有类的父类即是Object类,而Object类中的clone方法就是返回对象的一个拷贝,但是其作用域protected类型的,一般类无法调用,因此,ProtoType类中需要将clone方法的作用域修改为public类型

## 代码实现

```java
 /** Prototype Class **/
 public class Cookie implements Cloneable {
   
    public Object clone() throws CloneNotSupportedException
    {
        //In an actual implementation of this pattern you would now attach references to
        //the expensive to produce parts from the copies that are held inside the prototype.
        return (Cookie) super.clone();
    }
 }
 
 /** Concrete Prototypes to clone **/
 public class CoconutCookie extends Cookie { }
 
 /** Client Class**/
 public class CookieMachine
 {
 
   private Cookie cookie;//cookie必须是可复制的
 
     public CookieMachine(Cookie cookie) { 
         this.cookie = cookie; 
     } 

    public Cookie makeCookie()
    {
        try
        {
            return (Cookie) cookie.clone();
        } catch (CloneNotSupportedException e)
        {
            e.printStackTrace();
        }
        return null;
    } 

 
     public static void main(String args[]){ 
         Cookie tempCookie =  null; 
         Cookie prot = new CoconutCookie(); 
         CookieMachine cm = new CookieMachine(prot); //设置原型
         for(int i=0; i<100; i++) 
             tempCookie = cm.makeCookie();//通过复制原型返回多个cookie 
     } 
 }
```

- 我们可以看到, 原型模式是一种比较简单的模式，也非常容易理解，实现一个接口，重写一个方法即完成了原型模式。在实际应用中，原型模式很少单独出现。经常与其他模式混用，他的原型类Prototype也常用抽象类来替代。

## 模式优缺点

- 将创建对象简化高效

  使用原型模式创建对象是操作在Object类中的clone方法,而其方法是本地方法可以直接操作内存中的二进制流

  通过原型模式我们创建就像粘贴一样十分简单

## 模式适用场景

- 在需要重复地创建相似对象时可以考虑使用原型模式。比如需要在一个循环体内创建对象，假如对象创建过程比较复杂或者循环次数很多的话，使用原型模式不但可以简化创建过程，而且可以使系统的整体性能提高很多。

## 模式注意事项

- **使用原型模式复制对象不会调用类的构造方法**:我们上面提到我们实现的是Object类中的clone方法,它是直接在内存中复制数据,因此不会调用类的构造方法.而且构造方法中的代码不仅不会执行,而且连访问权限都在该模式下无效(protected作用域)

  > 还记得单例模式吗？单例模式中，只要将构造方法的访问权限设置为private型，就可以实现单例。但是clone方法直接无视构造方法的权限，所以，单例模式与原型模式是冲突的，在使用时要特别注意。



- **深拷贝与浅拷贝**:Object类的clone方法只会拷贝对象中的基本的数据类型，对于数组、容器对象、引用对象等都不会拷贝，这就是浅拷贝。如果要实现深拷贝，必须将原型模式中的数组、容器对象、引用对象等另行拷贝。例如接下来这个例子:

  ```java
  public class Prototype implements Cloneable {
  	private ArrayList list = new ArrayList();
  	public Prototype clone(){
  		Prototype prototype = null;
  		try{
  			prototype = (Prototype)super.clone();
  			prototype.list = (ArrayList) this.list.clone();
  		}catch(CloneNotSupportedException e){
  			e.printStackTrace();
  		}
  		return prototype; 
  	}
  }
  ```

  由于ArrayList不是基本类型，所以成员变量list，不会被拷贝，需要我们自己实现深拷贝，幸运的是java提供的大部分的容器类都实现了Cloneable接口。所以实现深拷贝并不是特别困难。

  > **深拷贝与浅拷贝问题中，会发生深拷贝的有java中的8中基本类型以及他们的封装类型，另外还有String类型。其余的都是浅拷贝。**

## 参考

https://blog.csdn.net/zhengzhb/article/details/7393528

[https://zh.wikipedia.org/wiki/%E5%8E%9F%E5%9E%8B%E6%A8%A1%E5%BC%8F](https://zh.wikipedia.org/wiki/原型模式)

