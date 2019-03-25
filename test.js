print("init");
var foo = 0;
function test() {
    for (var i = 0; i < 1000; i++) {
        foo++;

        if (i % 200 === 0) {
            print("step: " + i);
        }
    }

    print("finished: " + foo);
}
print("init done");
