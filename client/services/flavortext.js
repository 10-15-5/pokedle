import * as apiService from './api/apiService.js'

const setNewSecretPokemonFlavortext = async () => {
    const response = await apiService.newSecretPokemon(); //TODO: implement
    //Object.assign(secretPokemon, response.data);
    localStorage.classicSecretPokemon = JSON.stringify(response.data);
};

const setSecretPokemonFlavortext = async () => {
    const response = await apiService.getSecretPokemon(); //TODO: implement
    //Object.assign(secretPokemon, response.data);
    localStorage.classicSecretPokemon = JSON.stringify(response.data);
};

export {
    setNewSecretPokemonFlavortext,
    setSecretPokemonFlavortext,   
}