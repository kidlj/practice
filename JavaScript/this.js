let image = {
    name: 'Jian Li',
    transform() {
        console.log(this.name)
    }
}

function test(func) {
    // 此处调用与 image 对象无关
    func()
}

image.transform() // print 'Jian Li'

test(image.transform) // print undefined

let f = image.transform
f() // print undefined

// 箭头函数
let o = {
    color: "blue"
}

let sayColor = () => {
    console.log(this.color)
}
sayColor() // undefined
o.sayColor = sayColor
o.sayColor() // undefined

// 函数表达式
let b = {
    num: 3
}

let sayNum = function () {
    console.log(this.num)
}
sayNum() // undefined
b.sayNum = sayNum
b.sayNum() // 3

// 闭包
let object = {
    identity: "My Object",
    getIdentityFunc() {
        // 每个函数被调用时会自动创建 this 变量
        // this 不作为闭包变量被内部函数 access
        return function () {
            return this.identity
        }
    },
    getIdentityArrow() {
        // this 作为闭包变量被内部箭头函数 access
        return () => {
            return this.identity
        }
    }
}
console.log(object.getIdentityFunc()()) // undefined
console.log(object.getIdentityArrow()()) // "My Object"