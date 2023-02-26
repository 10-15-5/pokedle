import {apiClient} from './httpClient.js';

const getSecretPokemon = () => apiClient().get(`secret-pokemons`);

const newSecretPokemon = () => apiClient().post(`secret-pokemons`);

const updateDailyGamesWonCount = () => apiClient().post(`games-won`);

const getDailyStats = (date) => apiClient().get(`daily-stats/${date}`);

const createUser = () => apiClient().post(`user`);

export {
    getSecretPokemon,
    newSecretPokemon,
    updateDailyGamesWonCount,
    getDailyStats,
    createUser
}