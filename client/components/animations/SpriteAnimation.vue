<template>
    <span class="relative flex h-10 w-10 sm:h-8 sm:w-8 items-center justify-center">
            <span class="absolute w-5 h-5 bg-red-500 opacity-0 rounded-full cursor-pointer" @click="disable"></span>
            <img :src="getRandomPokemonAnimation()" />
        </span>
</template>

<script setup>
import { ref } from 'vue';
import {getRandomPokemonAnimation} from '../../services/assets.js'

const props = defineProps({
    isShowNotification: {
        type: Boolean,
        default: true,
    },
});

const isClicked = ref(false);

const disable = () => {
    isClicked.value = true;
};

</script>
