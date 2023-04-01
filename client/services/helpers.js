const removeSpecialCharactersExceptDashFromString = (text) => {
    if (!text) return '';
    return text.replace(/[^-\w\s]/gi, '');
};

const getCookie = (document) =>
Object.fromEntries(document.cookie.split('; ').map((v) => v.split(/=(.*)/s).map(decodeURIComponent)));

const lowerCaseAndCapitalizeWord = (word) => {
    if (typeof word !== 'string') throw new Error('word must be string');

    const lowerCaseWord = word.toLowerCase();
    const capitalizedWord = lowerCaseWord.charAt(0).toUpperCase() + lowerCaseWord.slice(1);

    return {
        lowerCaseWord,
        capitalizedWord,
    };
};

export {
    removeSpecialCharactersExceptDashFromString,
    getCookie,
    lowerCaseAndCapitalizeWord,
};
