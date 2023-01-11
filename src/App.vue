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
            <SquareContentHeader/>
            <v-list-item v-for="(guess, i) in state.guesses"
                         :key="i"
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
    <v-btn @click="revealPokemon">reveal secret pokemon</v-btn>
    <!-- <HomeView /> -->
  </v-app>
</template>

<script setup>
import SquareContainer from './components/SquareContainer.vue';
import SquareContentHeader from './components/SquareContentHeader.vue';
import SearchField from './components/SearchField.vue';
import pokemonData from './assets/pokemon.json';
import { ref } from 'vue';
import { guessState } from './constants.js';

const state = ref({
  pokemonNames: pokemonData.map((pokemonInfo) => pokemonInfo.name).sort(),
  guesses: [],
});
const correctFields = pokemonData[Math.floor(Math.random()*pokemonData.length)];
const pokemonToGuess = correctFields.name;

//TODO: Should eventually be handled on the backend
const getGuessResults = (pokemonName) => {
  const data = pokemonData.find(e => e.name === pokemonName);
  const result = {
    name: {
      name: pokemonName,
      guessState: guessState.None
    },
    type1: {
      text: data.type1,
      guessState: data.type1 === correctFields.type1 ? guessState.CorrectGuess : guessState.WrongGuess,
    },
    type2: {
      text: !data.type2 ? "None" : data.type2 ,
      guessState: data.type2 === correctFields.type2 ? guessState.CorrectGuess : guessState.WrongGuess
    },
    evolutions: {
      text: '' + data.evolutions,
      guessState: data.evolutions === correctFields.evolutions ? guessState.CorrectGuess : guessState.WrongGuess
    },
    evolutionLevel: {
      text: '' + data.evolutionState,
      guessState: data.evolutionState === correctFields.evolutionState ? guessState.CorrectGuess : guessState.WrongGuess
    },
    generation: {
      text: "Gen " + data.generation,
      guessState: data.generation === correctFields.generation ? guessState.CorrectGuess : guessState.WrongGuess
    },
  }

  if (
    result.type1.guessState !== guessState.CorrectGuess &&
    result.type2.guessState !== guessState.CorrectGuess
  ) {
    if (data.type1 === correctFields.type2) result.type1.guessState = guessState.PartlyCorrectGuess;
    if (data.type2 === correctFields.type1) result.type2.guessState = guessState.PartlyCorrectGuess;
  }

  return result;
}

//TODO: Should eventually be handled on the backend
const submitGuess = (guess) => {
  if(!guess) return;

  let guessRemovedFromList = false;
  state.value.pokemonNames = state.value.pokemonNames.filter(e => {
    if (!guessRemovedFromList && e.startsWith(guess.toLowerCase())) {
      state.value.guesses.unshift(e);
      guessRemovedFromList = true;
    }
    else return true;
  });

  if(!guessRemovedFromList) return;

  guess = state.value.guesses[0];
  if (guess === pokemonToGuess) {
    console.log("🥳🎉🎊 Congrats! You guessed the secret pokemon: " + guess);
  } else {
    console.log("❌❌❌ Wrong Guess. The secret pokemon was not " + guess + " ❌❌❌");
  }

}

const revealPokemon = () => {
  console.log(pokemonToGuess);
}

</script>

<style scoped>
.guess-container {
  background-color: transparent;
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
