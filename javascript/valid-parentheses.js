/**
 * @param {string} s
 * @return {boolean}
 */
 var isValid = function(s) {
  const openingBraces = new Set(['(', '[', '{']);
  const match = {
    ')': '(',
    '}': '{',
    ']': '['
  }
  
  const openingBracesOrder = [];
  
  if (openingBraces.has(s[0]) === false) {
      // Starts with a closing bracket so it's invalid
      return false;
  }
  
  if (s.length % 2 !== 0) {
    // Doesn't contain even number of characters so its invalid
    return false;
  }
  
  for (let i = 0; i < s.length; i++) {
    const char = s[i];
    
    if (openingBraces.has(char)) {
      openingBracesOrder.push(char);
    } else {
      const lastOpening = openingBracesOrder[openingBracesOrder.length - 1];
      if (match[char] === lastOpening) {
        openingBracesOrder.pop()
      } else {
        return false
      }
    }
  }
  
  if (openingBracesOrder.length === 0) {
    return true;
  }
  return false;
};