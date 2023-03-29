import moment from 'moment-timezone';
import { lowerCaseAndCapitalizeWord } from '../helpers';
import { GameModes } from '../constants';
const setNewDate = () => (localStorage.dayOfLastUpdate = moment().date());

const clearLocalStorageGameMode = (gameMode) => {
    //TODO: make generic to take params
    if (!gameMode) throw new Error('Gamemode Required');

    const { lowerCaseWord, capitalizedWord } = lowerCaseAndCapitalizeWord(gameMode);

    localStorage.removeItem(`${lowerCaseWord}SecretPokemon`);
    localStorage.removeItem(`${lowerCaseWord}Guesses`);
    localStorage.removeItem(`${lowerCaseWord}Colors`);
    localStorage.removeItem(`is${capitalizedWord}GameWon`);
};

const addGuessesToLocalStorage = (gameMode, guesses) => {
    if (!gameMode) throw new Error('Gamemode Required');
    const { lowerCaseWord } = lowerCaseAndCapitalizeWord(gameMode);
    localStorage.setItem(`${lowerCaseWord}Guesses`, JSON.stringify(guesses));
};

const addColorsToLocalStorage = (gameMode, colors) => {
    if (!gameMode) throw new Error('Gamemode Required');
    const { lowerCaseWord } = lowerCaseAndCapitalizeWord(gameMode);
    localStorage.setItem(`${lowerCaseWord}Colors`, JSON.stringify(colors));
};

const setSecretPokemonToLocalStorage = (gameMode, secretPokemon) => {
    switch (gameMode) {
        case GameModes.Classic:
            localStorage.classicSecretPokemon = JSON.stringify(secretPokemon);
            break;
        case GameModes.Flavortext:
            localStorage.flavortextSecretPokemon = JSON.stringify(secretPokemon);
            break;
        case GameModes.Silhouette:
            localStorage.silhouetteSecretPokemon = JSON.stringify(secretPokemon);
            break;
        default:
            throw new Error('Gamemode Required');
    }
};

const setSecretPokemonFromLocalStorage = (gameMode, secretPokemon) => {
    switch (gameMode) {
        case GameModes.Classic:
            if (localStorage.classicSecretPokemon) {
                Object.assign(secretPokemon, JSON.parse(localStorage.classicSecretPokemon));
            }
            return;
        case GameModes.Flavortext:
            if (localStorage.flavortextSecretPokemon) {
                Object.assign(secretPokemon, JSON.parse(localStorage.flavortextSecretPokemon));
            }
            return;
        case GameModes.Silhouette:
            if (localStorage.silhouetteSecretPokemon) {
                Object.assign(secretPokemon, JSON.parse(localStorage.silhouetteSecretPokemon));
            }
            return;
        default:
            throw new Error('Gamemode Required');
    }
};

const getColorsFromLocalStorage = (gameMode) => {
    if (!gameMode) throw new Error('Gamemode Required');
    const { lowerCaseWord } = lowerCaseAndCapitalizeWord(gameMode);
    const loadedColors = localStorage.getItem(`${lowerCaseWord}Colors`);
    if (loadedColors) return JSON.parse(loadedColors);
    else return [];
};

const getGuessesFromLocalStorage = (gameMode) => {
    if (!gameMode) throw new Error('Gamemode Required');
    const { lowerCaseWord } = lowerCaseAndCapitalizeWord(gameMode);
    const guesses = localStorage.getItem(`${lowerCaseWord}Guesses`);
    if (guesses) return JSON.parse(guesses);
    else return [];
};

export {
    setNewDate,
    setSecretPokemonFromLocalStorage,
    setSecretPokemonToLocalStorage,
    clearLocalStorageGameMode,
    addGuessesToLocalStorage,
    addColorsToLocalStorage,
    getColorsFromLocalStorage,
    getGuessesFromLocalStorage,
};
