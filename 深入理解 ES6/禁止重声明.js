// 假设作用域中已经存在某个标识符，此时再使用 let 关键字声明它就会抛出错误

var count = 1

// 抛出语法错误
let count = 2


// 如果当前作用域内嵌另一个作用域，便可在内嵌的作用域中用let声明同名变量

var value = 2

if (value === 2) {
    let value = 3
}


