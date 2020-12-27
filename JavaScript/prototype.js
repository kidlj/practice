function Person() { }

Person.prototype.name = "Nicholas";
Person.prototype.age = 29;
Person.prototype.job = "Software Engineer";
Person.prototype.sayName = function () {
    console.log(this.name);
}

let p1 = new Person();
let p2 = new Person();

console.log(p1.constructor === Person);
console.log(hasPrototypeProperty(p1, "constructor")); // true

console.log(typeof Person.prototype);
console.log(Person.prototype.constructor === Person);

console.log(p1.__proto__ === Person.prototype);
console.log(p1.__proto__.constructor === Person);

console.log(Person.prototype.__proto__ === Object.prototype);
console.log(Person.prototype.__proto__.constructor === Object);
console.log(Person.prototype.__proto__.__proto__ === null);

console.log(p1.__proto__ === p2.__proto__);

console.log(p1 instanceof Person);
console.log(p1 instanceof Object);
console.log(Person.prototype instanceof Object);

console.log(Person.prototype.isPrototypeOf(p1));
console.log(Person.prototype.__proto__.isPrototypeOf(p1));

console.log(Object.getPrototypeOf(p1) === Person.prototype);
console.log(Object.getPrototypeOf(p1).name === "Nicholas");

Person.prototype = {}
let p3 = new Person();
console.log(Person.prototype.constructor !== Person);
console.log(Person.prototype.constructor === Object);
console.log(p3.constructor === Object);
// Output:
// true
// true
// object
// true
// true
// true
// true
// true
// true
// true
// true
// true
// true
// true
// true
// true
// true
// true
// true
// true


function hasPrototypeProperty(object, name) {
    return !object.hasOwnProperty(name) && (name in object)
}