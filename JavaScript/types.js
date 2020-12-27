function testType() {
    let s = "hello"
    console.log(s instanceof String) // false

    s = new Object("hello")
    console.log(s instanceof String) // true

    s = new String("hello")
    console.log(s instanceof String) // true
}

testType()

function testEqual() {
    let a = 3
    let b = a
    let c = 3
    console.log(a === b)
    console.log(a === c)

    let s = new String("s")
    let e = new String("s")
    console.log(s === e) // false
}

testEqual()