import allPokemonData from '../../server/data/pokemonData-v6-translations.json';
import { GuessState, GuessType } from '../constants.js';
import { getRandomColor } from './game';

const getGuessResults = (pokemonName, secretPokemon, color) => {

    const pokemonData = allPokemonData.find((e) => e.name === pokemonName);

    if (!pokemonData) throw Error('No pokemon data found');

    const result = {
        name: {
            name: pokemonName,
            guessState: GuessState.None,
            type: GuessType.Pokemon,
        },
        type1: {
            text: pokemonData.type1,
            guessState:
                pokemonData.type1 === secretPokemon.type1
                    ? GuessState.CorrectGuess
                    : GuessState.WrongGuess,
            type: GuessType.Text,
        },
        type2: {
            text: !pokemonData.type2 ? 'None' : pokemonData.type2,
            guessState:
                pokemonData.type2 === secretPokemon.type2
                    ? GuessState.CorrectGuess
                    : GuessState.WrongGuess,
            type: GuessType.Text,
        },
        evolutionLevel: {
            text: '' + pokemonData.evolutionState,
            guessState:
                pokemonData.evolutionState === secretPokemon.evolutionState
                    ? GuessState.CorrectGuess
                    : GuessState.WrongGuess,
            type: GuessType.Text,
        },
        isFullyEvolved: {
            text: '' + pokemonData.isFullyEvolved,
            guessState:
                pokemonData.isFullyEvolved === secretPokemon.isFullyEvolved
                    ? GuessState.CorrectGuess
                    : GuessState.WrongGuess,
            type: GuessType.Text,
        },
        color: {
            text: pokemonData.color,
            guessState:
                pokemonData.color === secretPokemon.color
                    ? GuessState.CorrectGuess
                    : GuessState.WrongGuess,
            type: GuessType.Text,
        },
        habitat: {
            habitat: pokemonData.habitat,
            guessState:
                pokemonData.habitat === secretPokemon.habitat
                    ? GuessState.CorrectGuess
                    : GuessState.WrongGuess,
            type: GuessType.Habitat,
        },
        generation: {
            text: 'Gen ' + pokemonData.generation,
            guessState:
                pokemonData.generation === secretPokemon.generation
                    ? GuessState.CorrectGuess
                    : GuessState.WrongGuess,
            type: GuessType.Text,
        },
    };

    if (
        result.type1.guessState !== GuessState.CorrectGuess &&
        result.type2.guessState !== GuessState.CorrectGuess
    ) {
        if (pokemonData.type1 === secretPokemon.type2)
            result.type1.guessState = GuessState.PartlyCorrectGuess;
        if (pokemonData.type2 === secretPokemon.type1)
            result.type2.guessState = GuessState.PartlyCorrectGuess;
    }

    return {
        isCorrectGuess: pokemonName === secretPokemon.name,
        color,
        fields: result,
    };
};

const isCorrectGuess = (pokemonName, secretPokemonName) => pokemonName === secretPokemonName;

export { getGuessResults,isCorrectGuess };
