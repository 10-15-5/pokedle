import moment from 'moment-timezone';
import { DateOfFirstPokeldeGameClassic } from './constants.js';

const removeSpecialCharactersExceptDashFromString = (text) => {
    if (!text) return '';
    return text.replace(/[^-\w\s]/gi, '');
};

const getCookie = (document) =>
    Object.fromEntries(
        document.cookie.split('; ').map((v) => v.split(/=(.*)/s).map(decodeURIComponent))
    );

const getCurrentClassicPokemonNumber = () => {
    const oneDayInMS = 1000 * 60 * 60 * 24;

    const todayInMS = moment();

    return Math.round((todayInMS - DateOfFirstPokeldeGameClassic.valueOf()) / oneDayInMS);
};

export { removeSpecialCharactersExceptDashFromString, getCookie, getCurrentClassicPokemonNumber };
