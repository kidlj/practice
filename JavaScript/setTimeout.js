function double(value) {
    setTimeout(() => setTimeout(console.log, 0, value * 2), 1000)
    setTimeout(console.log, 2000, value * 3)
}

double(3)