var funcs = []

for (var i = 0; i < 10; i++) {
    funcs.push(function () {
        console.log(i)
    })
}

funcs.forEach(func => func()) // 输出 10 个 10