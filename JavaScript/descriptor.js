function testDescriptor() {
    let dest, src, result;

    dest = {
        set a(val) {
            console.log(`Invoked dest setter with param ${val}`);
        }
    }

    src = {
        get a() {
            console.log('Invoked src getter');
            return "foo";
        }
    }

    Object.assign(dest, src);
    console.log(dest);
}

testDescriptor();

