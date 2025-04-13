function testFunctionDeclaration() {
    console.log(FunctionExpression)
    var FunctionExpression = function () { }
    console.log(FunctionExpression)

    console.log(F)
    function F() { }
    console.log(F)
}


function testVarDeclaration() {
    console.log(v)
    var v = 3
    console.log(v)
}

function testLetDeclaration() {
    // console.log(v) // ReferenceError
    let v = 4
}

function testClassDeclaration() {
    // console.log(Cl) // ReferenceError
    class Cl { }
    console.log(Cl)
}

testFunctionDeclaration()
testVarDeclaration()
testLetDeclaration()
testClassDeclaration()