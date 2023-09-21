function flatlandSpaceStations(n, c) {
    c.sort((a, b) => a - b);
    let result = c[0];
    
    for (let i = 1; i < c.length; i++) {
        result = Math.max(result, Math.floor((c[i] - c[i - 1]) / 2));
    }
    
    result = Math.max(result, n - 1 - c[c.length - 1]);
    return result;
}
