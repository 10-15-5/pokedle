const removeSpecialCharactersExceptDashFromString = (text) => {
  return text.replace(/[^-\w\s]/gi, '');
}

export {
  removeSpecialCharactersExceptDashFromString
}