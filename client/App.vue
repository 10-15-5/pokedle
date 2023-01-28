<template>
    <v-app id="app">
        <v-main>
            <div class="container">
                <div>
                    <v-img src="/client/assets/title.png"
                           class="title-img mt-8" />
                </div>
                <div v-if="!isGameWon"
                     class="search-field mt-8">
                    <SearchField :pokemonNames="state.pokemonNames"
                                 @submit-guess="submitGuess" />
                </div>
                <GameWinContainer v-else
                                  :pokemon="state.guesses[0]" />
                <div v-if="state.guesses.length"
                     class="guess-container">
                    <SquareContentHeader class="mb-n1" />
                    <v-card v-for="(guess, i) in state.guesses"
                            :key="guess"
                            :value="guess">
                        <SquareContainer :pokemonName="guess"
                                         :guessResults="getGuessResults(guess)" />
                    </v-card>
                </div>
            </div>
            <!-- <square-container></square-container> -->
        </v-main>
        <div class="game-button-container">
            <v-btn @click="revealPokemon">reveal secret pokemon</v-btn>
            <v-btn @click="newGame">new game</v-btn>
            <v-btn @click="resetGame">resetGame game</v-btn>
        </div>
        <!-- <HomeView /> -->
    </v-app>
</template>

<script setup>
import SquareContainer from './components/SquareContainer.vue';
import SquareContentHeader from './components/SquareContentHeader.vue';
import GameWinContainer from './components/GameWinContainer.vue';
import SearchField from './components/SearchField.vue';
import pokemonData from '../server/database/pokemonData-v3.json';
import { onMounted, reactive, ref } from 'vue';
import { guessState } from './constants.js';
import { getSecretPokemon, newSecretPokemon } from './services/service';

//Use ref here? https://github.com/vuejs/docs/issues/801#issuecomment-757587022
const state = reactive({
    pokemonNames: pokemonData.map((pokemonInfo) => pokemonInfo.name).sort(),
    guesses: [],
});
const isGameWon = ref(false)

let secretPokemon;

const getGuessResults = (pokemonName) => {

    const data = pokemonData.find(e => e.name === pokemonName);

    if (!data) throw Error("No pokemon data found");

    const result = {
        name: {
            name: pokemonName,
            guessState: guessState.None
        },
        type1: {
            text: data.type1,
            guessState: data.type1 === secretPokemon.type1 ? guessState.CorrectGuess : guessState.WrongGuess,
        },
        type2: {
            text: !data.type2 ? "None" : data.type2,
            guessState: data.type2 === secretPokemon.type2 ? guessState.CorrectGuess : guessState.WrongGuess
        },
        evolutionLevel: {
            text: '' + data.evolutionState,
            guessState: data.evolutionState === secretPokemon.evolutionState ? guessState.CorrectGuess : guessState.WrongGuess
        },
        isFullyEvolved: {
            text: '' + data.isFullyEvolved,
            guessState: data.isFullyEvolved === secretPokemon.isFullyEvolved ? guessState.CorrectGuess : guessState.WrongGuess
        },
        generation: {
            text: "Gen " + data.generation,
            guessState: data.generation === secretPokemon.generation ? guessState.CorrectGuess : guessState.WrongGuess
        },
    }

    if (
        result.type1.guessState !== guessState.CorrectGuess &&
        result.type2.guessState !== guessState.CorrectGuess
    ) {
        if (data.type1 === secretPokemon.type2) result.type1.guessState = guessState.PartlyCorrectGuess;
        if (data.type2 === secretPokemon.type1) result.type2.guessState = guessState.PartlyCorrectGuess;
    }

    return result;
}

//TODO: Should eventually be handled on the backend
const submitGuess = (guess) => {
    if (!guess) return;

    let guessRemovedFromList = false;
    const updatedPokemonNames = state.pokemonNames.filter(e => {
        if (!guessRemovedFromList && e.startsWith(guess.toLowerCase())) {
            state.guesses.unshift(e);
            console.log("You guessed: " + e)
            addGuessesToLocalStorage(state.guesses);
            guessRemovedFromList = true;
        }
        else return true;
    });

    if (!guessRemovedFromList) return;
    state.pokemonNames = updatedPokemonNames;

    guess = state.guesses[0];
    if (guess === secretPokemon.name) {
        isGameWon.value = true
        localStorage.setItem('isGameWon', 'true')
        console.log("🥳🎉🎊 Congrats! You guessed the secret pokemon: " + guess);
    } else {
        console.log("❌❌❌ Wrong Guess. The secret pokemon was not " + guess + " ❌❌❌");
    }

}

const addGuessesToLocalStorage = (guesses) => {
    localStorage.setItem('guesses', JSON.stringify(guesses));
}

const loadSecretPokemon = () => {
    const loadedSecretPokemon = localStorage.getItem('secretPokemon');
    secretPokemon = JSON.parse(loadedSecretPokemon);
}

const loadGuesses = () => {
    const loadedGuesses = localStorage.getItem('guesses');
    if (loadedGuesses) state.guesses = JSON.parse(loadedGuesses);
}

const loadIsGameWon = () => {
    const loadedIsGameWon = localStorage.getItem('isGameWon');
    if (loadedIsGameWon === "true") {
        isGameWon.value = true;
    }
    else {
        localStorage.setItem('isGameWon', "false")
        isGameWon.value = false
    }
}

const removeGuessesFromSelection = () => {
    state.pokemonNames = state.pokemonNames.filter(pokemon => {
        for (const guessedPokemon of state.guesses) {
            if (guessedPokemon === pokemon) return false;
        }
        return true;
    });
}

const setNewDate = () => localStorage
    .setItem('dayOfLastUpdate', new Date().getUTCDate().toString());


const setNewSecretPokemon = async () => {
    const response = await getSecretPokemon();
    secretPokemon = response.data;
    localStorage.setItem('secretPokemon', JSON.stringify(secretPokemon));
}

const clearGuesses = () => localStorage.setItem('guesses', "");

onMounted(async () => {
    const dayOfLastUpdate = localStorage.getItem('dayOfLastUpdate');

    if (dayOfLastUpdate && parseInt(dayOfLastUpdate) == new Date().getUTCDate()) {
        loadSecretPokemon();
        loadGuesses();
        loadIsGameWon();
        removeGuessesFromSelection();
    } else {
        localStorage.clear();
        setNewDate();
        await setNewSecretPokemon();
    }
});

const revealPokemon = async () => {
    console.log(localStorage.getItem('secretPokemon'))
    const backendResponse = await getSecretPokemon();
    console.log(backendResponse.data)
}

const newGame = async () => {
    localStorage.clear();
    await newSecretPokemon();
    location.reload();
}

const resetGame = async () => {
    localStorage.clear();
    location.reload();
}

</script>

<style scoped>
.game-button-container {
    display: flex;
    justify-content: center;
}

.guess-container {
    background-color: transparent;
    margin-bottom: 50%;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    gap: 8px;
}

.search-field {
    display: flex;
    justify-content: center;
}

.title-img {
    transition: transform 0.2s;
    min-width: 230px;
    -webkit-filter: drop-shadow(0px 0px 20px rgb(255, 255, 255));
}

.title-img:hover {
    transform: scale(1.1);
}

.container {
    display: flex;
    align-items: center;
    flex-direction: column;
    width: auto;
    row-gap: 16px;
}

main {
    background-image: url("./client/assets/background-white.jpg");
    background-size: cover;
    background-repeat: no-repeat;
    background-attachment: fixed;
}
</style>
