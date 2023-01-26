import {apiClient} from './httpClient.js';

const getSecretPokemon = () => apiClient().get(`secret-pokemons`)

const newSecretPokemon = () => apiClient().post(`secret-pokemons`)

export {
    getSecretPokemon,
    newSecretPokemon
}