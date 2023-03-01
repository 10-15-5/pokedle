import {apiClient} from './httpClient.js';

const getSecretPokemon = () => apiClient().get(`secret-pokemons`);

const newSecretPokemon = () => apiClient().post(`secret-pokemons`);

const updateDailyGamesWonCount = (numberOfGuesses) => apiClient().post(`games-won`, {numberOfGuesses});

const getDailyStats = (date) => apiClient().get(`daily-stats/${date}`);

const getUser = (userId) => apiClient().get(`users/${userId}`);

const createUser = () => apiClient().post(`users`);

const updateUserWithGameWon = (userId,numberOfGuesses) => apiClient().post(`users/${userId}/gamesWon`, {numberOfGuesses});

export {
    getSecretPokemon,
    newSecretPokemon,
    updateDailyGamesWonCount,
    getDailyStats,
    createUser,
    getUser,
    updateUserWithGameWon
}