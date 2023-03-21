import * as apiService from './api/apiService.js'

const setNewSecretPokemon = async () => {
    const response = await apiService.newSecretPokemon();
    //Object.assign(secretPokemon, response.data);
    localStorage.classicSecretPokemon = JSON.stringify(response.data);
};

const setSecretPokemon = async () => {
    const response = await apiService.getSecretPokemon();
    //Object.assign(secretPokemon, response.data);
    localStorage.classicSecretPokemon = JSON.stringify(response.data);
};

const clearLocalstorageClassic = () => {
    localStorage.removeItem('classicSecretPokemon');
    localStorage.removeItem('classicGuesses');
    localStorage.removeItem('classicColors');
    localStorage.removeItem('isClassicGameWon');
}


export {
    setNewSecretPokemon,
    setSecretPokemon,
    clearLocalstorageClassic
}