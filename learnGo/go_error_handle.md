# GO 错误处理

## 一，error 处理的几种方式  

参考好文章  
>https://studygolang.com/articles/5188

- 普通用法，每次判断err == nil
- 封装checkError函数，有错误时 os.Exit()
- 提取出err, err!=nil时return,  defer函数处理err信息
- 将doStuff和error处理绑定, 最后判断err, 中间多个函数调用不判断err
- exported Error 使err可判断其内类或者值（os包export的错误码变量）

## 二，使用os.Exit(errCode int) 退出程序，和C的错误码表示一样，0为正常

## 三，使用recover和panic

参考好文章  
>https://blog.csdn.net/chenbaoke/article/details/41966827 

在defer函数中定义recover， A->B->C   在A函数defer中recover可以捕获到A和B和C的panic  
调用栈中只到最上层都没有defer：recover处理的话，程序异常终止，抛出异常。

