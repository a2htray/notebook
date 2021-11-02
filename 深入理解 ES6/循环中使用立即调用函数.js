var funcs = []

for (var i = 0; i < 10; i++) {
    funcs.push(function (value) {
        return function () {
            console.log(value)
        }
    }(i))
}

funcs.forEach(func => func()) // 输出 10 个 10