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