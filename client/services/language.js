import * as test from '../language/test.json';
import * as english from '../language/english.json';
import * as german from '../language/german.json';
import types from '../language/types.json';
import colors from '../language/colors.json';

const languages = ['german', 'english'];

const setLanguage = (lang) => {

    if (lang && typeof lang === 'string' && languages.includes(lang.toLowerCase())) {
        localStorage.language = lang;
        return true;
    }
};

function setLanguageFromLanguageCode() {

    if (localStorage.language) return;

    const preferredLanguages = navigator.languages || [navigator.language];
    // Use the first preferred language as the user's language
    const userLanguage = preferredLanguages[0];

    const languageMap = {
        de: languages[0],
        en: languages[1],
    };

    const mappedLanguage = languageMap[userLanguage.substr(0, 2)];

    const lang = languages.includes(mappedLanguage) ? mappedLanguage : 'english'

    setLanguage(lang)
}

const getLanguage = () => localStorage.language || 'english';

const text = () => {

    !(localStorage.language) ? setLanguage('english') : false;

    switch (localStorage.language) {
        case 'english':
            return english;
        case 'german':
            return german;
        case 'test':
            return test;
        default:
            throw new Error('This should never happen');
    }
};

const translateYesOrNo = (x) => (x === 'yes' ? text().resultCards.yes : text().resultCards.no);

const getFlavortextLanguageFrom = (secretPokemon) =>
    secretPokemon.flavortexts.find((flavortext) => flavortext.language === getLanguage()).flavortext;

const translateType = (type) => types.find((e) => e.type === type).types.find((t) => t.language === getLanguage()).type;

const translateColor = (color) =>
    colors.find((e) => e.color === color).colors.find((c) => c.language === getLanguage()).color;

const getGuessFieldTooltipsFromIndex = (index) => {
    const values = Object.values(text().guessFieldTooltips);
    return values.at(index);
};

const getTranslatedName = (names) => names.find((e) => e.language === getLanguage()).name;

export {
    setLanguage,
    getLanguage,
    text,
    getGuessFieldTooltipsFromIndex,
    getFlavortextLanguageFrom,
    translateType,
    translateYesOrNo,
    translateColor,
    languages,
    getTranslatedName,
    setLanguageFromLanguageCode,
};
