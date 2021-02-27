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