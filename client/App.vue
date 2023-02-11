<template>
    <v-app id="app">
        <v-main>
            <div class="container">
                <div>
                    <v-img src="/client/assets/title.png"
                           class="title-img mt-4" />
                    <HeaderContainer />
                </div>
                <SearchField v-if="!isGameWon"
                             :pokemonNames="state.pokemonNames"
                             
                             @submit-guess="submitGuess" />
                <GameWinContainer v-else
                                  :pokemon="state.guesses[0]"
                                  :color="colors.at(-1)" />
                <div v-if="state.guesses.length"
                     class="guess-container">
                    <SquareContentHeader class="mb-n1" />
                    <v-card v-for="(guess, i) in state.guesses"
                            :key="guess"
                            :value="guess">
                        <SquareContainer :pokemonName="guess"
                                         :guessResult="getGuessResults(guess, secretPokemon)"
                                         :color="colors.at(state.guesses.length - 1 - i)" />
                    </v-card>
                </div>
            </div>
            <!-- <square-container></square-container> -->
        </v-main>
        <div class="game-button-container">
            <v-btn @click="revealPokemon">reveal secret pokemon</v-btn>
            <v-btn @click="newGame">new game</v-btn>
        </div>
        <!-- <HomeView /> -->
    </v-app>
</template>

<script setup>
import SquareContainer from './components/SquareContainer.vue';
import SquareContentHeader from './components/SquareContentHeader.vue';
import GameWinContainer from './components/GameWinContainer.vue';
import SearchField from './components/SearchField.vue';
import pokemonData from '../server/data/pokemonData-v4.json';
import HeaderContainer from './components/HeaderContainer.vue';
import { onMounted, reactive, ref } from 'vue';
import { getGuessResults } from './services/guess';
import { getSecretPokemon, newSecretPokemon } from './services/service';

//Use ref here? https://github.com/vuejs/docs/issues/801#issuecomment-757587022
const state = reactive({
    pokemonNames: pokemonData.map((pokemonInfo) => pokemonInfo.name).sort(),
    guesses: [],
});
const isGameWon = ref(false);
const isSearchFieldDisabled = ref(false);

let colors = [];

let secretPokemon;

const getRandomColor = () => {
    const random = Math.random() * 100;
    return random < 5 ? 'shiny' : 'normal';
}

const removePokemonFromGuessPool = (guess) => {
    let guessRemovedFromList = false;
    let pokemonName = "";
    const updatedPokemonNames = state.pokemonNames.filter(e => {
        if (!guessRemovedFromList && e.startsWith(guess.toLowerCase())) {
            state.guesses.unshift(e);
            pokemonName = e;
            guessRemovedFromList = true;
        }
        else return true;
    });

    return {
        pokemonName,
        updatedPokemonNames,
    }
}

const decideGame = (guess) => {
    if (guess === secretPokemon.name) {
        isSearchFieldDisabled.value=true;

        //Wait for all cards to flip
        setTimeout(() => {
            isGameWon.value = true;
            localStorage.setItem('isGameWon', 'true')
            console.log("🥳🎉🎊 Congrats! You guessed the secret pokemon: " + guess);
        }, 2750); 
    } else {
        console.log("❌❌❌ Wrong Guess. The secret pokemon was not " + guess + " ❌❌❌");
    }
}

const addGuessesToLocalStorage = () => {
    localStorage.setItem('guesses', JSON.stringify(state.guesses));
}

const addColorsToLocalStorage = () => {
    localStorage.setItem('colors', JSON.stringify(colors));
}

const submitGuess = (guess) => {
    if (!guess || isSearchFieldDisabled.value) return;

    const {
        updatedPokemonNames,
        pokemonName
    } = removePokemonFromGuessPool(guess);

    if (updatedPokemonNames.length >= state.pokemonNames.length) return;

    colors.push(getRandomColor());
    state.pokemonNames = updatedPokemonNames;
    addGuessesToLocalStorage();
    addColorsToLocalStorage();

    decideGame(pokemonName);
}


const loadSecretPokemon = () => {
    const loadedSecretPokemon = localStorage.getItem('secretPokemon');
    secretPokemon = JSON.parse(loadedSecretPokemon);
}

const loadGuesses = () => {
    const guesses = localStorage.getItem('guesses');
    if (guesses) state.guesses = JSON.parse(guesses);
}

const loadColors = () => {
    const loadedColors = localStorage.getItem('colors');
    if (loadedColors) colors = JSON.parse(loadedColors);
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

const removePokemonsFromGuessPool = () => {
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
    //TODO: currently we set new secret pokemon every day at 00:00. 
    // Users will see a random pokemon each, since we rely on this instead of job/CDN
    const response = await newSecretPokemon();
    secretPokemon = response.data;
    localStorage.setItem('secretPokemon', JSON.stringify(secretPokemon));
}

onMounted(async () => {
    const dayOfLastUpdate = localStorage.getItem('dayOfLastUpdate');

    if (dayOfLastUpdate && parseInt(dayOfLastUpdate) == new Date().getUTCDate()) {
        loadSecretPokemon();
        loadColors();
        loadGuesses();
        loadIsGameWon();
        removePokemonsFromGuessPool();
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
