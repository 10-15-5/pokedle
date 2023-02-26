const removeSpecialCharactersExceptDashFromString = (text) => {
    if(!text) return "";
  return text.replace(/[^-\w\s]/gi, '');
}

const getCookie = (document) => Object.fromEntries(document.cookie.split('; ')
    .map(v=>v.split(/=(.*)/s).map(decodeURIComponent)));


export {
  removeSpecialCharactersExceptDashFromString,
  getCookie
}