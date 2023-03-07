<template>
<div class="card font-pkmEmerald p-1 mt-[-8px] text-xs capitalize">
    <p>Yesterdays Pokemon Was <b class=" text-green-600 dark:!text-green-300">{{ pokemonName }}</b> </p>
</div>
</template>

<script setup>
import { onMounted, ref } from 'vue';
import {getPreviousSecretPokemon} from '../services/apiService.js';
const pokemonName = ref('')
onMounted(async () => {
    const yesterdaysPokemon = (await getPreviousSecretPokemon()).data;
    pokemonName.value = yesterdaysPokemon.name;
})
</script>