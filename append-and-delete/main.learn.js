function appendAndDelete(from, to, steps) {
    let count = 0
    let fl = from.length
    let tl = to.length
  
    if (from == to || fl + tl <= steps) return 'Yes'
  
    for (let i in to) {
      if (from[i] == to[i]) {
        count++
      } else break
    }
  
    let check = (fl - count) + (tl - count) - steps
    if (check <= 0 && check % 2 == 0) {
      return 'Yes'
    }
  
    return 'No'
  }