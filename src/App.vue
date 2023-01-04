<template>
    <v-app id="app">
        <v-main>
            <div class="container">
                <div class="header-container">
                    <v-img
                        src="/src/assets/title.png"
                        class="title-img"
                    ></v-img>
                </div>
                <div>
                    <h1 class="title">Guess todays pokemon!</h1>
                </div>
                <div class="guess-field">
                    <input
                        label="Type pokemon name"
                        v-model="searchTerm"
                        class="p-3 mb-0.5 w-full border border-gray-300 rounded"
                    >
                    <ul
                        v-if="searchPokemonNames.length"
                        class="w-full rounded bg-white border border-gray-300 px-4 py-2 space-y-1 absolute z-10"
                    >
                        <li
                            v-for="pokemonName in searchPokemonNames"
                            :key="pokemonName"
                            @click="selectPokemonName(pokemonName)"
                            class="cursor-pointer hover:bg-gray-100 p-1"
                        > {{ pokemonName }} </li>
                    </ul>
                </div>
                <div>
                    <square-container></square-container>
                </div>
            </div>
        </v-main>
        <!-- <HomeView /> -->
    </v-app>
</template>

<script>
import HomeView from './views/HomeView.vue';
import SquareContainer from './views/SquareContainer.vue';
import pokemonData from './assets/pokemon.json';
import { computed, ref } from 'vue';

export default {
    setup() {
        let searchTerm = ref('');
        const pokemonNames = pokemonData.map((pokemonInfo) => pokemonInfo.Name);
        let selectedPokemonName = ref('');

        const searchPokemonNames = computed(() => {
            if (searchTerm.value === '') {
                return [];
            }

            console.log("call to searchPokemonNames")

            let matches = 0

            return pokemonNames.filter(name => {
                if (
                    name.toLowerCase().startsWith(searchTerm.value.toLowerCase())
                    && matches < 5
                ) {
                    matches++;
                    return name
                }
            });
        });

        const selectPokemonName = (pokemonName) => {
            selectedPokemonName.value = pokemonName;
            searchTerm.value = '';
            console.log("select: " + selectedPokemonName.value)
        }

        return {
            searchTerm,
            pokemonNames,
            searchPokemonNames,
            selectPokemonName,
            selectedPokemonName,
        }
    }
}

</script>

<style>
.header-container {
    width: 80%;
    max-width: 250px;
}

.guess-field {
    background: white;
    width: 300px;
    height: 25px;
    margin-top: 20px;
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
    gap: 10px;
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
