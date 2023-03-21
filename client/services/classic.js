import * as apiService from './api/apiService.js'

const setNewSecretPokemonClassic = async () => {
    const response = await apiService.newSecretPokemon();
    //Object.assign(secretPokemon, response.data);
    localStorage.classicSecretPokemon = JSON.stringify(response.data);
};

const setSecretPokemonClassic = async () => {
    const response = await apiService.getSecretPokemon();
    //Object.assign(secretPokemon, response.data);
    localStorage.classicSecretPokemon = JSON.stringify(response.data);
};

export {
    setNewSecretPokemonClassic,
    setSecretPokemonClassic,   
}