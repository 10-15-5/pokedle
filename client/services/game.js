import * as apiService from './api/apiService.js';
import { GameModes } from '../constants';
import moment from 'moment-timezone';
import { useStore } from '../stores/store.js';

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
}

const getDailyGamesWonCount = async (gameMode) => {
    const date = moment().format('YYYY-MM-DD');

    const res = await apiService.getDailyStats(date);

    switch (gameMode) {
        case GameModes.Classic:
            return res.data.classicGamesWon;
        case GameModes.Flavortext:
            return res.data.flavortextGamesWon;
        case GameModes.Silhouette:
            return res.data.silhouetteGamesWon;
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

export {
    setNewSecretPokemon,
    getAndSetSecretPokemon,
    setSecretPokemon,
    getDailyGamesWonCount,
    updateUserWithGameWon,
    updateCurrentUserStreakDisplay,
};
