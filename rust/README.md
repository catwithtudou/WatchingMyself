[TOC]

# Rust Begin

## Rust语言版本说明

- Rust语言的版本包括以下三个相互正交的概念：
  - 语义化版本（Sem Ver，Semantic Versioning）
  - 发行版本
  - Edition版次

1. **语义化版本（Sem Ver, Semantic versioning）**

其格式为：`主版本号.次版本号.修订号`

- 语义版本号递增规则：
  - 主版本号：当做了不兼容的API修改；
  - 次版本号：当做了向下兼容的功能性新增；
  - 修订号：当做了向下兼容的问题修正；

![image-20210225201734434](http://img.zhengyua.cn/20210225201734.png)

![image-20210225201750210](http://img.zhengyua.cn/20210225201750.png)

![image-20210225201828443](http://img.zhengyua.cn/20210225201828.png)

## Rust词法结构

### Rust编译过程

![image-20210225202037721](http://img.zhengyua.cn/20210225202037.png)

### Rust词法结构

包含其六大部分：

- 关键字（Keywords）
- 标识符（Identifier）
- 注释（Comment）
- 空白（Whitespace）
- 词条（Tokens）
- 路径（Path）

1. **关键字**

- 严格关键字（Strict）
  - as/ break/ const/ continue/ crate/ if/ else/ struct/ enum/ true/ false/ fn/ for/ in/ let/ loop/ impl/ mod/ match/ move / mut /pub/ ref/ return/ self/ Self/ static/ super/ trait/ type/ unsafe/ use/ where/ while /async/ await/ dyn / main
- 保留字（Reserved）
  - abstract/ become/ box/ do/ final/ macro/ override/ priv/ typeof/ unsized/ virtual/ yield / try
  - 被保留的关键字不代表将来一定会使用
- 弱关键字（Weak）
  - 2018 Edition: union, ‘static
  - 2015 Edition: dyn
  - 被保留的关键字不代表将来一定会使用

2. **标识符**

![image-20210225202442423](http://img.zhengyua.cn/20210225202442.png)

3. **注释**

![image-20210225202508603](http://img.zhengyua.cn/20210225202508.png)

![image-20210225202534392](http://img.zhengyua.cn/20210225202534.png)

4. **空白**

- Rust中空白字符包括：`\n`、`\t`、`tab`等；
- 任何形式的空白字符在 Rust 中只用于分隔标记，无语义意义；

5. **词条**

1. 语言项（item）
2. 块（block）
3. 语句（Stmt）
4. 表达式 （Expr）
5. 模式（Pattern）
6. 关键字（Keyword）
7. 标识符（Ident）
8. 字面量（Literal）
9. 生命周期（Lifetime）

10. 可见性（Vis）
11. 标点符号（Punctuation）
12. 分隔符（delimiter）
13. 词条树（Token Tree）
14. 属性（Attribute）

6. **路径**

![image-20210225203003899](http://img.zhengyua.cn/20210225203003.png)

## Rust语法面面观：面向表达式

### 表达式和语句

表达式和语句中的广义视角：**每行代码都可以看作是一个语句**。

语句的四种类型：

- 声明语句
- 流程控制语句
- 表达式语句
- 宏语句

Rust 语法的“骨架”可以**缩为一行**，并最终可以简化为三个关键元素：

- 属性，类似于`#![...]`；
- 分号`;`即行分隔符；
- 花括号`{...}`即块分隔符；

Rust 为面向表达式的语言，借鉴了函数式语言，面向表达式。

#### 面向表达式

1. **分号和块**是 Rust 语言的两个基本表达式。

- 分号表达式

  - 单元类型（Unit Type）

    `; -> ()`

- 块表达式

  - 块中最后一个表达式的值

2. Rust 中的**求值规则**：

- 分号表达式返回值永远为自身的单元（Unit）类型：`()`；
- 分号表达式只有在块表达式最后一行才会进行求值，其他时候只作为*连接符*存在；
- 块表达式只对其最后一行表达式进行求值；

3. **验证求值规则**

```Rust
fn main() {
    ;
    ;
    {
        ()
    }
    {
        ();
        use std::vec::Vec;
    }
    ();
    &{;}; // -> &()
    ;     // -> ()
    ;
}
```

4. **另一种划分方式**

- 基本语句
  - 声明语句
  - 表达式语句
- 表达式
  - 块中最后一行不加分号的表达式

流程控制也是表达式。

**除了基本的声明语句，其他皆为表达式**。


