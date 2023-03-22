import * as apiService from './api/apiService.js';
import { GameModes } from '../constants';

const setNewSecretPokemon = async (gameMode) => {
    let response

    switch (gameMode) {
        case GameModes.Classic:
            response = await apiService.newClassicSecretPokemon();
            localStorage.classicSecretPokemon = JSON.stringify(response.data);
            break;

        case GameModes.Flavortext:
            response = await apiService.newFlavortextSecretPokemon();
            localStorage.flavortextSecretPokemon = JSON.stringify(response.data);            
            break;
        default:
            break;
    }
    //Object.assign(secretPokemon, response.data);
};

const setSecretPokemon = async (gameMode) => {
    let response

    switch (gameMode) {
        case GameModes.Classic:
            response = await apiService.getClassicSecretPokemon();
            localStorage.classicSecretPokemon = JSON.stringify(response.data);
            break;

        case GameModes.Flavortext:
            response = await apiService.getFlavortextSecretPokemon();
            localStorage.flavortextSecretPokemon = JSON.stringify(response.data);            
            break;
        default:
            break;
    }

};

export { setNewSecretPokemon, setSecretPokemon };
