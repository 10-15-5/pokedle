<template>
    <div id="app">
        <div
            class="min-h-screen w-full"
            :class="{
                'background-white': !store.isDark && !store.isHintMode,
                'background-black': store.isDark && !store.isHintMode,
                'background-white-pokemon-collage': !store.isDark && store.isHintMode,
                'background-black-pokemon-collage': store.isDark && store.isHintMode,
            }"
        >
            <ThemeButton class="absolute right-0" />
            <div class="flex flex-col items-center justify-center">
                <RouterLink to="/">
                    <img
                        src="/client/assets/logo/pokedle-logo.png"
                        class="title-img mt-5 mb-3 w-[280px] transition-all duration-100 ease-linear hover:scale-110 sm:w-[220px]"
                    />
                </RouterLink>
                <HeaderContainer />
            </div>
            <RouterView class="mt-6" />
        </div>
        <div v-if="isDevelopment" class="flex items-center justify-center">
        <button class="card p-2 text-xs hover:!bg-green-400" @click="revealPokemon">Reveal</button>
        <button class="card p-2 text-xs hover:!bg-purple-400" @click="getNewGame">
            Get New Game
        </button>
        <button class="card p-2 text-xs hover:!bg-blue-400" @click="setNewGame">
            Set New Game
        </button>
        <button class="card p-2 text-xs hover:!bg-pink-400" @click="launchConfetti(false, false)">
            Confetti
        </button>
        <button class="card p-2 text-xs hover:!bg-pink-400" @click="playWinnerSound">
            WIN SOUND
        </button>
    </div>
    </div>
</template>

<script setup>
import HeaderContainer from './components/headerIcons/HeaderIconContainer.vue';
import ThemeButton from './components/buttons/ThemeButton.vue';
import { reactive, ref, onBeforeMount, computed } from 'vue';
import * as apiService from './services/api/apiService.js';
import { useStore } from './stores/store.js';
import moment from 'moment-timezone';
import { launchConfetti } from './services/confetti';
import { setNewSecretPokemon, setSecretPokemon } from './services/game.js';
import { clearLocalStorageGameMode } from './services/localStorage';
import { playWinnerSound } from './services/sound';
import {GameModes} from './constants'

const store = useStore();

const isDevelopment = computed(() => ENVIRONMENT === 'development');

onBeforeMount(async () => {
    store.loadTheme();
    loadIsHintMode();
    loadIsShiny();

    const user = await getOrCreateUser();

    console.log('Loaded at: ' + moment().toString());
    console.log('ENVIRONMENT: ' + ENVIRONMENT);
    console.log(user);

    if (user) {
        store.setUser(user);
    }
});

const getOrCreateUser = async () => {
    const userId = localStorage.userId;

    console.log(userId);
    if (userId) {
        await apiService.updateUserStreak(userId);
        const response = await apiService.getUser(userId);
        return response.data.user;
    }

    const response = await apiService.createUser();
    const { user } = response.data;
    localStorage.userId = user._id;
    return user;
};

const loadIsShiny = () => {
    if (localStorage.isShiny && localStorage.isShiny === 'true') {
        store.setShiny(true);
        return;
    }
};

const loadIsHintMode = () => {
    if (localStorage.isHintMode === 'true') {
        store.setHintMode(true);
    } else {
        store.setHintMode(false);
    }
};

const getNewGame = async () => {

    //Classic
    clearLocalStorageGameMode(GameModes.Classic);
    await setSecretPokemon(GameModes.Classic)

    //Flavortext
    clearLocalStorageGameMode(GameModes.Flavortext)
    await setSecretPokemon(GameModes.Flavortext)
    //TODO: implement

    location.reload();
};

// Force update pokemon in DB
const setNewGame = async () => {

    //Classic
    clearLocalStorageGameMode(GameModes.Classic);
    await setNewSecretPokemon(GameModes.Classic);

    //Flavortext
    clearLocalStorageGameMode(GameModes.Flavortext)
    await setNewSecretPokemon(GameModes.Flavortext);
    //TODO: implement

    location.reload();  
};

const revealPokemon = async () => {
    console.log(localStorage.classicSecretPokemon);
    const backendResponse = await apiService.getSecretPokemon();
    console.log(backendResponse.data);
};
</script>

<style scoped>
.background-white-pokemon-collage {
    background-image: url('./client/assets/backgrounds/background-white-pokemon-collage.png');
    background-size: cover;
    background-repeat: no-repeat;
    background-attachment: fixed;
}

.background-black-pokemon-collage {
    background: linear-gradient(rgba(0, 0, 0, 0.25), rgba(0, 0, 0, 0.25)),
        url('./client/assets/backgrounds/background-black-pokemon-collage.jpg');
    background-size: cover;
    background-repeat: no-repeat;
    background-attachment: fixed;
}

.background-white {
    background-image: url('./client/assets/backgrounds/background-white.jpeg');
    background-size: cover;
    background-repeat: no-repeat;
    background-attachment: fixed;
}

.background-black {
    background-image: url('./client/assets/backgrounds/background-black.jpeg');
    background-size: cover;
    background-repeat: no-repeat;
    background-attachment: fixed;
}
</style>
