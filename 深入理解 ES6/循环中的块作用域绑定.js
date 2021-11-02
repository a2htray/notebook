for (var i = 0; i < 10; i++) {
    console.log(i)
}

// 因为使用 var 声明的变量会得到提升，所以这里也可以访问到 i
console.log(i)

for (let j = 0; j < 10; j++) {
    console.log(j)
}

// 这里就不能访问 j
console.log(j)