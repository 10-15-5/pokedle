import * as apiService from './api/apiService.js';
import {
    GameModes,
    DateOfFirstPokeldeGameClassic,
    DateOfFirstPokeldeGameFlavortext,
    DateOfFirstPokeldeGameSilhouette,
} from '../constants';
import moment from 'moment-timezone';
import { useStore } from '../stores/store.js';
import { lowerCaseAndCapitalizeWord } from './helpers.js';
import * as localStorageService from './localStorage.js';
import { playWinnerSound } from '../services/sound.js';
import { launchConfetti } from '../services/confetti.js';

const getAndSetNewSecretPokemon = async (gameMode) => {
    let response;

    switch (gameMode) {
        case GameModes.Classic:
            response = await apiService.newClassicSecretPokemon();
            localStorageService.setSecretPokemonToLocalStorage(gameMode, response.data);
            break;
        case GameModes.Flavortext:
            response = await apiService.newFlavortextSecretPokemon();
            localStorageService.setSecretPokemonToLocalStorage(gameMode, response.data);
            break;
        case GameModes.Silhouette:
            response = await apiService.newSilhouetteSecretPokemon();
            localStorageService.setSecretPokemonToLocalStorage(gameMode, response.data);
            break;
        default:
            throw new Error('Gamemode Required');
    }
};

const getCurrentSecretPokemon = async (gameMode) => {
    let response;

    switch (gameMode) {
        case GameModes.Classic:
            response = await apiService.getClassicSecretPokemon();
            return response.data;
        case GameModes.Flavortext:
            response = await apiService.getFlavortextSecretPokemon();
            return response.data;
        case GameModes.Silhouette:
            response = await apiService.getSilhouetteSecretPokemon();
            return response.data;
        default:
            throw new Error('Gamemode Required');
    }
};
const getAndSetCurrentSecretPokemon = async (gameMode) => {
    let pokemon;
    switch (gameMode) {
        case GameModes.Classic:
            pokemon = await getCurrentSecretPokemon(gameMode);
            localStorageService.setSecretPokemonToLocalStorage(gameMode, pokemon);
            break;
        case GameModes.Flavortext:
            pokemon = await getCurrentSecretPokemon(gameMode);
            localStorageService.setSecretPokemonToLocalStorage(gameMode, pokemon);
            break;
        case GameModes.Silhouette:
            pokemon = await getCurrentSecretPokemon(gameMode);
            localStorageService.setSecretPokemonToLocalStorage(gameMode, pokemon);
            break;
        default:
            throw new Error('Gamemode Required');
    }
};

const setDailyGamesWonCount = async (gameMode) => {
    const store = useStore();

    const date = moment().format('YYYY-MM-DD');
    const res = await apiService.getDailyStats(date);

    switch (gameMode) {
        case GameModes.Classic:
            return store.setDailyClassicGamesWon(res.data.classicGamesWon);
        case GameModes.Flavortext:
            return store.setDailyFlavortextGamesWon(res.data.flavortextGamesWon);
        case GameModes.Silhouette:
            return store.setDailySilhouetteGamesWon(res.data.Silhouette);
        default:
            throw new Error('Gamemode Required');
    }
};

const updateUserWithGameWon = async (gameMode, numberOfGuesses) => {
    const userId = localStorage.userId;
    let response;

    switch (gameMode) {
        case GameModes.Classic:
            response = await apiService.updateUserClassicWins(userId, numberOfGuesses);
            return response.data.user;
        case GameModes.Flavortext:
            response = await apiService.updateUserFlavortextWins(userId, numberOfGuesses);
            return response.data.user;
        case GameModes.Silhouette:
            response = await apiService.updateUserSilhouetteWins(userId, numberOfGuesses);
            return response.data.user;
        default:
            throw new Error('Gamemode Required');
    }
};

const getYesterdaysPokemon = async (gameMode) => {
    let response;

    switch (gameMode) {
        case GameModes.Classic:
            response = await apiService.getClassicPreviousSecretPokemon();
            return response.data;
        case GameModes.Flavortext:
            response = await apiService.getFlavortextPreviousSecretPokemon();
            return response.data;
        case GameModes.Silhouette:
            response = await apiService.getSilhouettePreviousSecretPokemon();
            return response.data;
        default:
            throw new Error('Gamemode Required');
    }
};

const updateCurrentUserStreakDisplay = async (gameMode) => {
    const store = useStore();
    switch (gameMode) {
        case GameModes.Classic:
            return store.setCurrentStreak(store.user.classicCurrentStreak);
        case GameModes.Flavortext:
            return store.setCurrentStreak(store.user.flavortextCurrentStreak);
        case GameModes.Silhouette:
            return store.setCurrentStreak(store.user.silhouetteCurrentStreak);
        default:
            throw new Error('Gamemode Required');
    }
};

const incrementGamesWonCount = async (gameMode, numberOfGuesses) => {
    const store = useStore();

    let response;

    switch (gameMode) {
        case GameModes.Classic:
            response = await apiService.updateStatsClassicWins(numberOfGuesses);
            return store.setDailyClassicGamesWon(response.data.dailyStats.classicGamesWon);
        case GameModes.Flavortext:
            response = await apiService.updateStatsFlavortextWins(numberOfGuesses);
            return store.setDailyFlavortextGamesWon(response.data.dailyStats.flavortextGamesWon);
        case GameModes.Silhouette:
            response = await apiService.updateStatsSilhouetteWins(numberOfGuesses);
            return store.setDailySilhouetteGamesWon(response.data.dailyStats.silhouetteGamesWon);
        default:
            throw new Error('Gamemode Required');
    }
};

const removePokemonNameFromArray = (name, names) => {
    let guessRemovedFromList = false;
    let removedName = '';

    const updatedNames = names.filter((e) => {
        if (!guessRemovedFromList && e.startsWith(name.toLowerCase())) {
            removedName = e;
            guessRemovedFromList = true;
        } else return {};
    });

    return {
        removedName,
        updatedNames,
    };
};

const getRandomColor = () => (localStorage.isShiny === 'true' ? 'shiny' : Math.random() * 100 < 5 ? 'shiny' : 'normal');

const submitGuess = (gameMode, guess, componentStore, colors) => {
    if (!guess) return;

    const { removedName, updatedNames } = removePokemonNameFromArray(guess, componentStore.pokemonNames);

    if (updatedNames.length >= componentStore.pokemonNames.length) return;

    colors.push(getRandomColor());
    componentStore.guesses.unshift(removedName);
    componentStore.pokemonNames = updatedNames;
    localStorageService.addGuessesToLocalStorage(gameMode, componentStore.guesses);
    localStorageService.addColorsToLocalStorage(gameMode, colors);

    return removedName;
};

const decideGame = async (
    gameMode,
    guess,
    secretPokemonName,
    color,
    numberOfGuesses,
    soundDelayMS = 0,
    confettiDelayMS = 0,
    winFunction = () => {}
) => {
    if (guess === secretPokemonName) {
        const store = useStore();

        winFunction();

        setTimeout(() => {
            playWinnerSound();
        }, soundDelayMS);

        setTimeout(async () => {
            launchConfetti(color === 'shiny', numberOfGuesses === 1);

            setIsGameWon(gameMode, true);

            const user = await updateUserWithGameWon(gameMode, numberOfGuesses);
            await incrementGamesWonCount(gameMode, numberOfGuesses);

            store.setUser(user);
            updateCurrentUserStreakDisplay(gameMode);
        }, confettiDelayMS);
    }
};

const shouldLoadExistingGameData = (dayOfLastUpdate, localStorageSecretPokemon, databaseSecretPokemon) =>
    parseInt(dayOfLastUpdate) == moment().date() &&
    localStorageSecretPokemon &&
    localStorageSecretPokemon?.name === databaseSecretPokemon?.name;

const loadAndSetIsGameWon = (gameMode) => {
    const { capitalizedWord } = lowerCaseAndCapitalizeWord(gameMode);

    const loadedIsGameWon = localStorage.getItem(`is${capitalizedWord}GameWon`);

    if (loadedIsGameWon && loadedIsGameWon === 'true') {
        setIsGameWon(gameMode, true);
    } else {
        setIsGameWon(gameMode, false);
    }
};

const setIsGameWon = (gameMode, isWon) => {
    const store = useStore();

    switch (gameMode) {
        case GameModes.Classic:
            return store.setIsClassicGameWon(isWon);
        case GameModes.Flavortext:
            return store.setIsFlavortextGameWon(isWon);
        case GameModes.Silhouette:
            return store.setIsSilhouetteGameWon(isWon);
        default:
            throw new Error('Gamemode Required');
    }
};

const removeAllGuessedPokemonsFromGuessPool = (componentStore) => {
    componentStore.pokemonNames = componentStore.pokemonNames.filter((pokemon) => {
        for (const guessedPokemon of componentStore.guesses) {
            if (guessedPokemon === pokemon) return false;
        }
        return true;
    });
};

const loadExistingGameData = (gameMode, componentStore) => {
    componentStore.guesses = localStorageService.getGuessesFromLocalStorage(gameMode);
    loadAndSetIsGameWon(gameMode);
    removeAllGuessedPokemonsFromGuessPool(componentStore);
};

const startNewGame = (gameMode, currSecretPokemon, secretPokemon) => {
    localStorageService.clearLocalStorageGameMode(gameMode);
    setIsGameWon(gameMode, false);
    localStorageService.setSecretPokemonToLocalStorage(gameMode, currSecretPokemon);
    localStorageService.setSecretPokemonFromLocalStorage(gameMode, secretPokemon);
    localStorageService.setNewDate();
};

const loadGameData = async (gameMode, setHints, secretPokemon, componentStore) => {
    const dayOfLastUpdate = localStorage.dayOfLastUpdate;
    if (!dayOfLastUpdate) localStorageService.setNewDate();

    localStorageService.setSecretPokemonFromLocalStorage(gameMode, secretPokemon);
    const currSecretPokemon = await getCurrentSecretPokemon(gameMode);

    if (shouldLoadExistingGameData(dayOfLastUpdate, secretPokemon, currSecretPokemon)) {
        loadExistingGameData(gameMode, componentStore);
    } else {
        startNewGame(gameMode, currSecretPokemon, secretPokemon);
    }
    setHints();
    updateCurrentUserStreakDisplay(gameMode);
};

const getCurrentPokemonNumber = (gameMode, todayInMS) => {
    const oneDayInMS = 1000 * 60 * 60 * 24;
    let dateOfFirstPokedlePokemon;

    switch (gameMode) {
        case GameModes.Classic:
            dateOfFirstPokedlePokemon = DateOfFirstPokeldeGameClassic;
            break;
        case GameModes.Flavortext:
            dateOfFirstPokedlePokemon = DateOfFirstPokeldeGameFlavortext;
            break;
        case GameModes.Silhouette:
            dateOfFirstPokedlePokemon = DateOfFirstPokeldeGameSilhouette;
            break;
        default:
            throw new Error('Gamemode Required');
    }
    console.log((todayInMS - dateOfFirstPokedlePokemon.valueOf()) / oneDayInMS);
    console.log(Math.floor((todayInMS - dateOfFirstPokedlePokemon.valueOf()) / oneDayInMS));
    return Math.floor((todayInMS - dateOfFirstPokedlePokemon.valueOf()) / oneDayInMS);

};

export {
    getAndSetNewSecretPokemon,
    getAndSetCurrentSecretPokemon,
    setDailyGamesWonCount,
    updateUserWithGameWon,
    updateCurrentUserStreakDisplay,
    submitGuess,
    decideGame,
    getRandomColor,
    getCurrentSecretPokemon,
    shouldLoadExistingGameData,
    loadAndSetIsGameWon,
    removeAllGuessedPokemonsFromGuessPool,
    setIsGameWon,
    getYesterdaysPokemon,
    loadGameData,
    getCurrentPokemonNumber
};
