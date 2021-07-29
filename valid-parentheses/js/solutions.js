/**
 * @param {string} s
 * @return {boolean}
 */
 var isValid = function(s) {
  const openingBraces = new Set(['(', '[', '{']);
  const closingBraces = new Set([')', ']', '}']);
  const match = {
    ')': '(',
    '}': '{',
    ']': '['
  }
  
  const openingBracesOrder = [];
  
  if (closingBraces.has(s[0]) || openingBraces.has(s[s.length -1])) {
      // Starts with a closing bracket or ends with an opening bracket
      // so it's invalid
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
