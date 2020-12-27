function SuperType() {
    this.superProperty = "hello"
    this.colors = ["red", "blue"]
}

function SubType() { }

SubType.prototype = new SuperType()

instance1 = new SubType()
instance1.superProperty = "world" // add property, basic type
instance1.colors.push("black")
console.log(instance1.superProperty)
console.log(instance1.colors)
console.log(instance1.hasOwnProperty("superProperty"))
console.log(instance1.hasOwnProperty("colors")) // false

instance2 = new SubType()
console.log(instance2.superProperty)
console.log(instance2.colors)
console.log(instance2.hasOwnProperty("superProperty"))
console.log(instance2.hasOwnProperty("colors"))

instance1.colors = instance1.colors // add property, reference type
instance1.colors.push("white")
console.log(instance1.colors)
console.log(instance2.colors)
console.log(instance1.hasOwnProperty("colors")) // true
console.log(instance2.hasOwnProperty("colors")) // false
// Output:
// world
// [ 'red', 'blue', 'black' ]
// true
// false

// hello
// [ 'red', 'blue', 'black' ]
// false
// false

// [ 'red', 'blue', 'black', 'white' ]
// [ 'red', 'blue', 'black', 'white' ]
// true
// false