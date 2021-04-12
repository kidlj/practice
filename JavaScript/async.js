async function foo() {
    try {
        console.log(1)
        await Promise.reject(3)
        console.log(4) // never prints
    } catch (e) {
        console.log(e)
    }
}

foo()
// Output:
// 1
// 3

console.log("end")