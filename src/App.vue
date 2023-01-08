<template>
  <v-app id="app">
    <v-main>
      <div class="container">
        <div>
          <v-img src="/src/assets/title.png"
                 class="title-img" />
        </div>
        <div>
          <h1 class="title">Guess todays pokemon!</h1>
        </div>
        <div class="search-field">
          <search-field :pokemonNames="state.pokemonNames"
                        @some-event="submitGuess" />
        </div>
        <div>
          <v-list class="guess-container">
            <v-list-subheader>Your guesses</v-list-subheader>
            <SquareContentHeader/>
            <v-list-item v-for="(guess, i) in state.guesses"
                         :key="i"
                         :value="guess">
              <square-container :pokemonName="guess"
                                :guessResults="getGuessResults(guess)" />
            </v-list-item>
          </v-list>
        </div>
        <!-- <square-container></square-container> -->
      </div>
    </v-main>
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
  pokemonNames: pokemonData.map((pokemonInfo) => pokemonInfo.name),
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

const submitGuess = (guess) => {

  if (guess.toLowerCase() === pokemonToGuess) {
    console.log("🥳🎉🎊 Congrats! You guessed the secret pokemon: " + pokemonToGuess);
  } else {
    console.log("❌❌❌ Wrong Guess. The secret pokemon was not " + guess + " ❌❌❌");
  }

  state.value.pokemonNames = state.value.pokemonNames.filter(e => {
    if (e.toLowerCase() === guess.toLowerCase()) {
      state.value.guesses.push(e);
      console.log(state.value.guesses)
    }
    else return true;
  });

}

</script>

<style>
.guess-container {
  background-color: transparent;
}

.search-field {
  height: 5%;
  width: 25%;
  min-width: 200px;
  z-index: 1;
}

.title {
  color: white;
}

.title-img {
  transition: transform 0.2s;
  min-width: 250px;

}

.title-img:hover {
  transform: scale(1.1);
}

.container {
  display: flex;
  align-items: center;
  flex-direction: column;
  gap: 5%;
  width: auto;
  height: 100%;
}

main {
  background-image: url("./assets/background.jpg");
  background-repeat: no-repeat;
  background-size: cover;
  background-color: black;
}

/* .v-application__wrap {
    background-color: black;
    background-image: url("./src/assets/background.jpg");
    background-repeat: no-repeat;
    height: 100%;
  } */
</style>
