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
            <v-list-item v-for="(guess, i) in state.guesses"
                         :key="i"
                         :value="guess">
              <square-container :pokemon-data="getPokemonDataFromName(guess)" />
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
import SquareContainer from './views/SquareContainer.vue';
import SearchField from './views/SearchField.vue';
import pokemonData from './assets/pokemon.json';
import { ref } from 'vue';

const state = ref({
  pokemonNames: pokemonData.map((pokemonInfo) => pokemonInfo.Name),
  guesses: [],
})
const pokemonToGuess = "Charizard";
const submitGuess = (guess) => {

  if (guess.toLowerCase() === pokemonToGuess.toLowerCase()) {
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

const getPokemonDataFromName = (guess) => pokemonData.find(e => {
  return e.Name === guess
});

</script>

<style>
.guess-container {
  background-color: aqua;
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
