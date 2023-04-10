<template>
    <div class="card mt-[-8px] p-1 font-pkmEmerald text-xs capitalize sm:text-[10px]">
        <p>
            {{ text().previousPokemon.yesterdaysPokemonWas }}
            <b class="text-light-emerald dark:!text-dark-emerald">{{ getTranslatedName(pokemonName) }}</b>
        </p>
    </div>
</template>

<script setup>
import { text, getTranslatedName  } from '../../services/language';
const props = defineProps({
    pokemonName: String,
});
</script>
