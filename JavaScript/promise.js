let p1 = Promise.resolve('foo')
let p2 = Promise.reject('bar')

// Promise.resolve()
setTimeout(console.log, 0, p1 === Promise.resolve(p1)) // true
setTimeout(console.log, 0, p1 === Promise.resolve(Promise.resolve(p1))) // true

setTimeout(console.log, 0, p2 === Promise.resolve(p2)) // true
setTimeout(console.log, 0, p2 === Promise.resolve(Promise.resolve(p2))) // true

// then()
let p8 = p1.then()
let p9 = p1.then(() => p1)

setTimeout(console.log, 0, p8); // resolved foo
setTimeout(console.log, 0, p9); // resolved foo
setTimeout(console.log, 0, p1 === p8); // false
setTimeout(console.log, 0, p1 === p9); // false

let p7 = p1.then(() => new Promise((resolve, reject) => setTimeout(() => resolve('baz'), 100)))
setTimeout(console.log, 0, 'p7:', p7) // pending
setTimeout(console.log, 200, 'p7:', p7) // resolved baz

// catch()
let p10 = p2.catch()
setTimeout(console.log, 0, p10); // rejected bar
setTimeout(console.log, 0, p2 === p10); // false

// finally()
let p20 = p1.finally()
setTimeout(console.log, 0, p20); // resolved foo
setTimeout(console.log, 0, p1 === p20); // false

let p21 = p1.finally(() => new Promise((resolve, reject) => setTimeout(() => resolve('qux'), 100)))
setTimeout(console.log, 0, 'p21', p21) // pending 
setTimeout(console.log, 200, 'p21', p21) // resolved foo

let p30 = Promise.resolve()
let p31 = p30.then(() => console.log('onResolved handler'))
console.log('then() returned')
console.log('p31:', p31) // pending!!!
setTimeout(console.log, 0, 'timeout p31:', p31) // resolved undefined

// chaining
let p40 = new Promise((resolve, reject) => {
    console.log('p40 executor')
    setTimeout(resolve, 5000)
})
p40.then(() => new Promise((resolve, reject) => {
    console.log('p41 executor')
    setTimeout(resolve, 5000)
})).then(() => new Promise((resolve, reject) => {
    console.log('p42 executor')
    setTimeout(resolve, 5000)
}))
console.log('p40:', p40) // pending
// Output:
// p40 executor
// p40: Promise { <pending> }
// p41 executor
// p42 executor