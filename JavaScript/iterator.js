class Counter {
    constructor(limit) {
        this.limit = limit
    }

    [Symbol.iterator]() {
        let count = 1
        let limit = this.limit

        return {
            [Symbol.iterator]() {
                return this
            },
            next() {
                if (count <= limit) {
                    return { done: false, value: count++ }
                } else {
                    return { done: true, value: undefined }
                }
            }
        }
    }
}

let counter = new Counter(3)
let ct = counter[Symbol.iterator]()

console.log("counter:")
for (let i of counter) {
    console.log(i)
}

console.log("ct:")
for (let i of ct) {
    console.log(i)
}

let iter1 = counter[Symbol.iterator]()
let iter2 = iter1[Symbol.iterator]()
console.log(iter1 === iter2) // true

console.log("counter:")
for (let item of counter) {
    console.log(item)
}

console.log("iter1:")
for (let item of iter1) {
    console.log(item)
}

// Equals iter2; 已迭代，无输出
console.log("iter2:")
for (let item of iter2) {
    console.log(item)
}

// class CounterIterator {
//     constructor(limit) {
//         this.count = 1
//         this.limit = limit
//     }

//     next() {
//         if (this.count <= this.limit) {
//             return { done: false, value: count++ }
//         } else {
//             return { done: true, value: undefined }
//         }
//     }
// }

// let ct = new CounterIterator(3) // TypeError: ct is not iterable
// for (let i of ct) {
//     console.log("here")
//     console.log(i)
// }