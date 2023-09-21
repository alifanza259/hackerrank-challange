function strangeCounter(t) {
	// Write your code here
   let counter = 3;
  
	while (t >= counter - 2) {
	  counter *= 2;
	}
  
	const startNum = counter / 2;
	const diff = t - (startNum - 2);
	return startNum - diff;
  }