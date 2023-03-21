import moment from "moment-timezone";
import { lowerCaseAndCapitalizeWord } from "../helpers";

const setNewDate = () => (localStorage.dayOfLastUpdate = moment().date());

const clearLocalStorageGameMode = (gameMode) => { //TODO: make generic to take params
    const { lowerCaseWord, capitalizedWord} = lowerCaseAndCapitalizeWord(gameMode);

    localStorage.removeItem(`${lowerCaseWord}SecretPokemon`);
    localStorage.removeItem(`${lowerCaseWord}Guesses`);
    localStorage.removeItem(`${lowerCaseWord}Colors`);
    localStorage.removeItem(`is${capitalizedWord}GameWon`);
}

const addGuessesToLocalStorage = (gameMode, guesses) => {
    const { lowerCaseWord } = lowerCaseAndCapitalizeWord(gameMode);
    localStorage.setItem(`${lowerCaseWord}Guesses`, JSON.stringify(guesses));
};

const addColorsToLocalStorage = (gameMode, colors) => {
    const { lowerCaseWord } = lowerCaseAndCapitalizeWord(gameMode);
    localStorage.setItem(`${lowerCaseWord}Colors`, JSON.stringify(colors));
};

export{ setNewDate, clearLocalStorageGameMode, addGuessesToLocalStorage, addColorsToLocalStorage}