<template>
<div class="card font-pkmEmerald p-1 mt-[-8px] text-xs sm:text-[10px] capitalize">
    <p>Yesterdays Pokemon Was <b class=" text-light-emerald dark:!text-dark-emerald">{{ pokemonName }}</b> </p>
</div>
</template>

<script setup>
import { onMounted, ref } from 'vue';
import {getPreviousSecretPokemon} from '../services/api/apiService.js';

const props = defineProps({
    pokemonName: String,
});
</script>