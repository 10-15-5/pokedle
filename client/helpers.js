const removeSpecialCharactersExceptDashFromString = (text) => {
    if(!text) return "";
  return text.replace(/[^-\w\s]/gi, '');
}

export {
  removeSpecialCharactersExceptDashFromString
}