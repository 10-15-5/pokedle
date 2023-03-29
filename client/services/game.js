import * as apiService from './api/apiService.js';
import { GameModes } from '../constants';
import moment from 'moment-timezone';
import { useStore } from '../stores/store.js';
import { lowerCaseAndCapitalizeWord } from '../helpers.js';
import { addColorsToLocalStorage, addGuessesToLocalStorage } from './localStorage.js';
import {playWinnerSound} from '../services/sound.js';
import {launchConfetti} from '../services/confetti.js';

const setNewSecretPokemon = async (gameMode) => {
    let response;

    switch (gameMode) {
        case GameModes.Classic:
            response = await apiService.newClassicSecretPokemon();
            localStorage.classicSecretPokemon = JSON.stringify(response.data);
            break;

        case GameModes.Flavortext:
            response = await apiService.newFlavortextSecretPokemon();
            localStorage.flavortextSecretPokemon = JSON.stringify(response.data);
            break;
        case GameModes.Silhouette:
            response = await apiService.newSilhouetteSecretPokemon();
            localStorage.silhouetteSecretPokemon = JSON.stringify(response.data);
            break;
        default:
            throw new Error('Gamemode Required');
    }
};

const getAndSetSecretPokemon = async (gameMode) => {
    let response;

    switch (gameMode) {
        case GameModes.Classic:
            response = await apiService.getClassicSecretPokemon();
            localStorage.classicSecretPokemon = JSON.stringify(response.data);
            break;
        case GameModes.Flavortext:
            response = await apiService.getFlavortextSecretPokemon();
            localStorage.flavortextSecretPokemon = JSON.stringify(response.data);
            break;
        case GameModes.Silhouette:
            response = await apiService.getSilhouetteSecretPokemon();
            localStorage.silhouetteSecretPokemon = JSON.stringify(response.data);
            break;
        default:
            throw new Error('Gamemode Required');
    }
};

const setSecretPokemon = (gameMode, secretPokemon) => {
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

const setDailyGamesWonCount = async (gameMode) => {
    const store = useStore();

    const date = moment().format('YYYY-MM-DD');
    const res = await apiService.getDailyStats(date);

    switch (gameMode) {
        case GameModes.Classic:
            return store.setDailyClassicGamesWon(res.data.classicGamesWon);
        case GameModes.Flavortext:
            return store.setDailyFlavortextGamesWon(res.data.classicGamesWon);
        case GameModes.Silhouette:
            return store.setDailySilhouetteGamesWon(res.data.classicGamesWon);
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

const getRandomColor = () =>
    localStorage.isShiny === 'true' ? 'shiny' : Math.random() * 100 < 5 ? 'shiny' : 'normal';

const submitGuess = (gameMode, guess, componentStore, colors) => {
    if (!guess) return;

    const { removedName, updatedNames } = removePokemonNameFromArray(guess, componentStore.pokemonNames);

    if (updatedNames.length >= componentStore.pokemonNames.length) return;

    colors.push(getRandomColor());
    componentStore.guesses.unshift(removedName);
    componentStore.pokemonNames = updatedNames;
    addGuessesToLocalStorage(gameMode, componentStore.guesses);
    addColorsToLocalStorage(gameMode, colors);

    return removedName;
};

const decideGame = async (gameMode, guess, secretPokemonName, color, numberOfGuesses) => {
    if (guess === secretPokemonName) {
        const store = useStore();

        playWinnerSound();

        launchConfetti(color === 'shiny', numberOfGuesses === 1);
        store.setIsFlavortextGameWon(true); //TODO: refactor this

        const user = await updateUserWithGameWon(gameMode, numberOfGuesses);
        await incrementGamesWonCount(gameMode, numberOfGuesses);

        store.setUser(user);
    }
};

export {
    setNewSecretPokemon,
    getAndSetSecretPokemon,
    setSecretPokemon,
    setDailyGamesWonCount,
    updateUserWithGameWon,
    updateCurrentUserStreakDisplay,
    submitGuess,
    decideGame,
    getRandomColor,
};
