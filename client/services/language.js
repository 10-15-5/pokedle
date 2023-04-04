import * as german from '../language/german.json';
import * as english from '../language/english.json';

const setLanguage = (lang) => {
    localStorage.language = lang;
}

const getLanguage = () =>  {
    localStorage.getItem('language') == null ? setLanguage('english') : false;
    // /client/language/' + localStorage.getItem('language') + '.json',
    switch(localStorage.language) {
        case "english":
            return english;
        case "german":
            return german;
        default:
            throw new Error("This should never happen")
    }
}

export {
    setLanguage,
    getLanguage,
}