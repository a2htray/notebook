// 变量 value 的声明被提升至函数顶部，而初始化操作依旧留在原处执行，
// 这就意味着在 else 子句中也可以访问到该变量，且由于此时该变量尚未
// 初始化，所以其值为 undefined。

function getValue(condition) {
    if (condition) {
        var value = "blue";
        // 其他代码
        return value;
    } else {
        // 此处 value 为 undefined
        return null
    }
    // 此处 value 为 undefined
}