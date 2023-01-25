<template>
  <v-app id="app">
    <v-main>
      <div class="container">
        <div>
          <v-img src="/src/assets/title.png"
                 class="title-img mt-8" />
        </div>
        <!-- <div>
          <h1 class="title">Guess todays pokemon!</h1>
        </div> -->
        <div>
          <div class="search-field mt-8">
            <search-field :pokemonNames="state.pokemonNames"
                          @submit-guess="submitGuess" />
          </div>
          <div v-if="state.guesses.length">
            <v-list class="guess-container">
              <SquareContentHeader />
              <v-list-item v-for="(guess, i) in state.guesses"
                           :key="guess"
                           :value="guess">
                <square-container :pokemonName="guess"
                                  :guessResults="getGuessResults(guess)" />
              </v-list-item>
            </v-list>
          </div>
        </div>
        <!-- <square-container></square-container> -->
      </div>
    </v-main>
    <div class="game-button-container">
      <v-btn @click="revealPokemon">reveal secret pokemon</v-btn>
      <v-btn @click="resetGame">reset game</v-btn>
    </div>
    <!-- <HomeView /> -->
  </v-app>
</template>

<script setup>
import SquareContainer from './components/SquareContainer.vue';
import SquareContentHeader from './components/SquareContentHeader.vue';
import SearchField from './components/SearchField.vue';
import pokemonData from '../dev-server/database/pokemonData-v2.json';
import { onMounted, ref, computed, reactive } from 'vue';
import { guessState } from './constants.js';
import { getSecretPokemon } from './services/service';

const state = reactive({
  pokemonNames: pokemonData.map((pokemonInfo) => pokemonInfo.name).sort(),
  guesses: [],
});

let secretPokemon;
let pokemonToGuess;

//TODO: Should eventually be handled on the backend
const getGuessResults = (pokemonName) => {

  console.log("PokemonName " + pokemonName);
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
  if (guess === pokemonToGuess) {
    console.log("🥳🎉🎊 Congrats! You guessed the secret pokemon: " + guess);
  } else {
    console.log("❌❌❌ Wrong Guess. The secret pokemon was not " + guess + " ❌❌❌");
  }

}

const addGuessesToLocalStorage = (guesses) => {
  localStorage.setItem('guesses', JSON.stringify(guesses));
}

onMounted(() => {
  const loadedSecretPokemon = localStorage.getItem('secretPokemon');

  if (loadedSecretPokemon) {
    secretPokemon = JSON.parse(loadedSecretPokemon);
    pokemonToGuess = secretPokemon.name;
  } else {
    secretPokemon = pokemonData[Math.floor(Math.random() * pokemonData.length)];
    localStorage.setItem('secretPokemon', JSON.stringify(secretPokemon));
    pokemonToGuess=secretPokemon.name;
  }

  const loadedGuesses = localStorage.getItem('guesses');
  if (loadedGuesses) {
    state.guesses = JSON.parse(loadedGuesses);
    state.pokemonNames = state.pokemonNames.filter(pokemon => {
      for (const guessedPokemon of state.guesses) {
        if(guessedPokemon === pokemon) return false;
      }
      return true;
    })
  }
});

const revealPokemon = async () => {
  console.log(pokemonToGuess);
  const backendResponse = await getSecretPokemon();
  console.log(`Backend Response: ${backendResponse}`)
}

const resetGame = () => {
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
}

.search-field {
  display: flex;
  justify-content: center;
}

.title-img {
  transition: transform 0.2s;
  min-width: 300px;
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
}

main {
  background-image: url("./assets/background-white.jpg");
  background-repeat: repeat;
  background-size: 100%;
  background-color: black;
}

/* .v-application__wrap {
    background-color: black;
    background-image: url("./src/assets/background.jpg");
    background-repeat: no-repeat;
    height: 100%;
  } */
</style>
