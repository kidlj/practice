class Person {
    constructor() {
        this.name = new String("Jack")
        this.sayName = () => console.log(this.name)
        this.nicknames = ["Jake", "J-Dog"]
    }

    // prototype method
    locate() {
        console.log("prototype", this)
    }

    // static method
    static locate() {
        console.log("class", this)
    }

    static anotherLocate() {
        console.log("class [anotherLocate]", this)
    }
}

console.log(Person.prototype.constructor === Person) // true

let p1 = new Person()
let p2 = new Person()

console.log(p1.name === p2.name) // false
console.log(p1.sayName === p2.sayName) // false
console.log(p1.nicknames === p2.nicknames) // false

p1.nicknames.push("Jackie")
console.log(p1.nicknames)
console.log(p2.nicknames) // not same

console.log(p1.locate === p2.locate) // true

console.log(Person.prototype.locate())
console.log(Person.locate())
console.log(Person.anotherLocate())