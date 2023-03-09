import allPokemonData from '../../server/data/pokemonData-v5-flavorText.json';
import { guessState, guessType } from '../constants.js';

const getGuessResults = (pokemonName, secretPokemon) => {

    const pokemonData = allPokemonData.find(e => e.name === pokemonName);

    if (!pokemonData) throw Error("No pokemon data found");

    const result = {
        name: {
            name: pokemonName,
            guessState: guessState.None,
            type: guessType.Pokemon
        },
        type1: {
            text: pokemonData.type1,
            guessState: pokemonData.type1 === secretPokemon.type1 ? guessState.CorrectGuess : guessState.WrongGuess,
            type: guessType.Text
        },
        type2: {
            text: !pokemonData.type2 ? "None" : pokemonData.type2,
            guessState: pokemonData.type2 === secretPokemon.type2 ? guessState.CorrectGuess : guessState.WrongGuess
            , type: guessType.Text
        },
        evolutionLevel: {
            text: '' + pokemonData.evolutionState,
            guessState: pokemonData.evolutionState === secretPokemon.evolutionState ? guessState.CorrectGuess : guessState.WrongGuess
            , type: guessType.Text
        },
        isFullyEvolved: {
            text: '' + pokemonData.isFullyEvolved,
            guessState: pokemonData.isFullyEvolved === secretPokemon.isFullyEvolved ? guessState.CorrectGuess : guessState.WrongGuess
            , type: guessType.Text
        },
        color: {
            text: pokemonData.color,
            guessState: pokemonData.color === secretPokemon.color ? guessState.CorrectGuess : guessState.WrongGuess
            , type: guessType.Text
        },
        habitat: {
            habitat: pokemonData.habitat,
            guessState: pokemonData.habitat === secretPokemon.habitat ? guessState.CorrectGuess : guessState.WrongGuess
            , type: guessType.Habitat
        },
        generation: {
            text: "Gen " + pokemonData.generation,
            guessState: pokemonData.generation === secretPokemon.generation ? guessState.CorrectGuess : guessState.WrongGuess,
            type: guessType.Text
        },
    }

    if (
        result.type1.guessState !== guessState.CorrectGuess &&
        result.type2.guessState !== guessState.CorrectGuess
    ) {
        if (pokemonData.type1 === secretPokemon.type2) result.type1.guessState = guessState.PartlyCorrectGuess;
        if (pokemonData.type2 === secretPokemon.type1) result.type2.guessState = guessState.PartlyCorrectGuess;
    }

    return result;
}

export {
    getGuessResults
}