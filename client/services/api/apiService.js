import { apiClient } from '../api/httpClient.js';

const getClassicSecretPokemon = () => apiClient().get(`classic-secret-pokemons`);
const getClassicPreviousSecretPokemon = () => apiClient().get(`classic-previous-secret-pokemons`);
const newClassicSecretPokemon = () => apiClient().post(`classic-secret-pokemons`);
const updateUserClassicWins = (userId, numberOfGuesses) =>
    apiClient().post(`users/${userId}/classic-wins`, { numberOfGuesses });

const getFlavortextSecretPokemon = () => apiClient().get(`flavortext-secret-pokemons`);
const getFlavortextPreviousSecretPokemon = () => apiClient().get(`flavortext-previous-secret-pokemons`);
const newFlavortextSecretPokemon = () => apiClient().post(`flavortext-secret-pokemons`);
const updateUserFlavortextWins = (userId, numberOfGuesses) =>
    apiClient().post(`users/${userId}/flavortext-wins`, { numberOfGuesses });

const updateStatsClassicWins = (numberOfGuesses) => apiClient().post(`classic-games-won`, { numberOfGuesses });
const updateStatsFlavortextWins = (numberOfGuesses) => apiClient().post(`classic-games-won`, { numberOfGuesses });
const getDailyStats = (date) => apiClient().get(`daily-stats/${date}`);

const getUser = (userId) => apiClient().get(`users/${userId}`);

const createUser = () => apiClient().post(`users`);

const updateUserStreaks = (userId) => apiClient().patch(`users/${userId}`);

export {
    getClassicSecretPokemon,
    newClassicSecretPokemon,
    getClassicPreviousSecretPokemon,
    updateStatsClassicWins,
    updateUserClassicWins,
    getFlavortextSecretPokemon,
    getFlavortextPreviousSecretPokemon,
    newFlavortextSecretPokemon,
    updateStatsFlavortextWins,
    updateUserFlavortextWins,
    getDailyStats,
    createUser,
    getUser,
    updateUserStreaks,
};
