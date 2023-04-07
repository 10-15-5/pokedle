import * as test from '../language/test.json';
import * as english from '../language/english.json';
import * as german from '../language/german.json';
import types from '../language/types.json';
import colors from '../language/colors.json';

const setLanguage = (lang) => {
    localStorage.language = lang;
};

const getLanguage = () => localStorage.language || 'english';

const text = () => {
    localStorage.getItem('language') == null ? setLanguage('english') : false;
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

const translateColor = (color) => colors.find((e) => e.color === color).colors.find((c) => c.language === getLanguage()).color;

const getGuessFieldTooltipsFromIndex = (index) => {
    const values = Object.values(text().guessFieldTooltips);
    return values.at(index);
};

export {
    setLanguage,
    getLanguage,
    text,
    getGuessFieldTooltipsFromIndex,
    getFlavortextLanguageFrom,
    translateType,
    translateYesOrNo,
    translateColor
};
