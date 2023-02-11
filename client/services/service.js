import {apiClient} from './httpClient.js';

const getSecretPokemon = () => apiClient().get(`secret-pokemons`);

const newSecretPokemon = () => apiClient().post(`secret-pokemons`);

const updateDailyGamesWonCount = () => apiClient().post(`games-won`);

export {
    getSecretPokemon,
    newSecretPokemon,
    updateDailyGamesWonCount
}