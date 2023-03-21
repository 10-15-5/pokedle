<template>
    <div
        class="card min-h-10 min-w-10 min flex max-w-[250px] flex-col items-center justify-center py-3 px-2 text-center font-pkmEmerald"
    >
        <p>{{ header }}</p>
        <p v-for="emojiResultRow in emojiResults" class="tracking-widest">
            {{ emojiResultRow }}
        </p>
        <p v-if="totalGuesses > 4">
            ＋{{ totalGuesses - 4 }} more {{ totalGuesses > 5 ? 'guesses' : 'guess' }}..
        </p>
    </div>
</template>

<script setup>
import { onMounted, computed } from 'vue';

const props = defineProps({
    emojiResults: Array,
    totalGuesses: Number,
});
onMounted(() => {
    console.log(props.emojiResults);
    console.log(props.totalGuesses);
});

const header = computed(
    () =>
        `I found the #1 hidden #Pokedle Pokémon ${
            props.totalGuesses === 1 ? 'FIRST TRY 🌟🥳🌠🏆' : `in ${props.totalGuesses} tries!🕵️🔎`
        }`
);

const footer = 'Play at pokedle.gg 🎮!';
</script>
