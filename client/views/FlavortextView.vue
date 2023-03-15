<template>
    <!-- <SearchField
        v-if="!store.isGameWon"
        :pokemonNames="componentStore.pokemonNames"
        @submit-guess="submitGuess"
    />
    <GameWinContainer v-else :pokemon="componentStore.guesses[0]" :color="colors.at(-1)" /> -->
</template>

<script setup>
import SearchField from '../components/SearchField.vue';
import GameWinContainer from '../components/GameWinContainer.vue';
</script>
