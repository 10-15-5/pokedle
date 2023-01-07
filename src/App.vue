<template>
  <v-app id="app">
    <v-main>
      <div class="container">
        <div class="header-container">
          <v-img src="/src/assets/title.png"
                 class="title-img"></v-img>
        </div>
        <div>
          <h1 class="title">Guess todays pokemon!</h1>
        </div>
        <div>
          <the-search-drop-down class="search-field"
          :pokemonNames="state.pokemonNames"
                                @some-event="submitGuess" />
        </div>
        <div>
            <v-list class="bg-red">
              <v-list-subheader>Your guesses</v-list-subheader>
              <v-list-item v-for="(guess, i) in state.guesses"
                           :key="i"
                           :value="guess"> {{ guess }} </v-list-item>
            </v-list>
            <!-- <square-container></square-container> -->
          </div>
      </div>
    </v-main>
    <!-- <HomeView /> -->
  </v-app>
</template>

<script setup>
import SquareContainer from './views/SquareContainer.vue';
import TheSearchDropDown from './views/TheSearchDropDown.vue';
import pokemonData from './assets/pokemon.json';
import { ref } from 'vue';

const state = ref({
  pokemonNames: pokemonData.map((pokemonInfo) => pokemonInfo.Name),
  guesses: [],
})
const pokemonToGuess = "Charizard";
const submitGuess = (guess) => {
  if (guess === pokemonToGuess) {
    console.log("🥳🎉🎊 Congrats! You guessed the secret pokemon: " + pokemonToGuess);
  } else {
    console.log("❌❌❌ Wrong Guess. The secret pokemon was not " + guess + " ❌❌❌");
  }
  state.value.pokemonNames = state.value.pokemonNames.filter(e => {
    if (e === guess) {
      state.value.guesses.push(guess);
      console.log(state.value.guesses)
    }
    else return true;
  });
}
</script>

<style>
.header-container {
  width: 80%;
  max-width: 250px;
}

.search-field {
  height: 5%;
  width: 25%;
  min-width: 200px;
}

.title {
  color: white;
}

.title-img {
  transition: transform 0.2s;
}

.title-img:hover {
  transform: scale(1.1);
}

.container {
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
  text-align: center;
  gap: 5%;
  width: 100%;
  height: 100%;
}

main {
  background-image: url("./assets/background.jpg");
  background-repeat: no-repeat;
  background-size: cover;
  background-color: black;
}
</style>
