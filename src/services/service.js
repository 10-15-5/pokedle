import {apiClient} from './httpClient.js';

const getSecretPokemon = () => apiClient().get(`secret-pokemons`)

export {
    getSecretPokemon
}