function SuperType() {
    this.property = true
}

SuperType.prototype.getSuperValue = function () {
    return this.property
}

function SubType() {
    this.subProperty = false
}

SubType.prototype = new SuperType()
SubType.prototype.getSubValue = function () {
    return this.subProperty
}

instance = new SubType()

console.log(instance.getSuperValue())
console.log(instance.constructor === SuperType)
console.log(hasPrototypeProperty(instance, "constructor"))
// Output:
// true
// true
// true

function hasPrototypeProperty(object, name) {
    return !object.hasOwnProperty(name) && (name in object)
}
