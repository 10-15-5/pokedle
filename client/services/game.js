import * as apiService from './api/apiService.js';
import { GameModes } from '../constants';

const setNewSecretPokemon = async (gameMode) => {
    switch (gameMode) {
        case GameModes.Classic:
            const response = await apiService.newSecretPokemon();
            localStorage.classicSecretPokemon = JSON.stringify(response.data);
            break;

        case GameModes.Flavortext:
            //TODO: implement
            break;
        default:
            break;
    }
    //Object.assign(secretPokemon, response.data);
};

const setSecretPokemon = async (gameMode) => {
    switch (gameMode) {
        case GameModes.Classic:
            const response = await apiService.getSecretPokemon();
            localStorage.classicSecretPokemon = JSON.stringify(response.data);
            break;

        case GameModes.Flavortext:
            //TODO: implement
            break;
        default:
            break;
    }

};

export { setNewSecretPokemon, setSecretPokemon };
