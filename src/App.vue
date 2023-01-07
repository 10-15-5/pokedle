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
        <div class="search-dropdown">
          <the-search-drop-down :pokemonNames="pokemonNames"
                                @some-event="submitGuess" />
        </div>
        <div>
          <square-container></square-container>
        </div>
      </div>
    </v-main>
    <!-- <HomeView /> -->
  </v-app>
</template>

<script setup>
import HomeView from './views/HomeView.vue';
import SquareContainer from './views/SquareContainer.vue';
import TheSearchDropDown from './views/TheSearchDropDown.vue';
import pokemonData from './assets/pokemon.json';

let pokemonNames = pokemonData.map((pokemonInfo) => pokemonInfo.Name);
const pokemonToGuess = "Charizard";
const submitGuess = (guess) => {
  if (guess === pokemonToGuess) {
    console.log("🥳🎉🎊 Congrats! You guessed the secret pokemon: " + pokemonToGuess);
  } else {
    //console.log(pokemonNames);
    console.log("❌❌❌ Wrong Guess. The secret pokemon was not " + guess + " ❌❌❌");
  }
}
</script>

<style>
.header-container {
  width: 80%;
  max-width: 250px;
}

.search-dropdown {
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
