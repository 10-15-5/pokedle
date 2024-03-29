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
            <RouterView v-if="isLoaded" class="mt-6" />
        </div>
        <div v-if="isDevelopment" class="flex items-center justify-center font-pkmEmerald">
            <button
                class="card overflow-hidden text-clip whitespace-nowrap p-2 text-xs hover:!bg-green-400 sm:text-[10px]"
                @click="revealPokemon"
            >
                Reveal
            </button>
            <button
                class="card overflow-hidden text-clip whitespace-nowrap p-2 text-xs hover:!bg-purple-400 sm:text-[10px]"
                @click="getCurrentGame"
            >
                Get New Game
            </button>
            <button
                class="card overflow-hidden text-clip whitespace-nowrap p-2 text-xs hover:!bg-blue-400 sm:text-[10px]"
                @click="setNewGame"
            >
                Set New Game
            </button>
            <button
                class="card overflow-hidden text-clip whitespace-nowrap p-2 text-xs hover:!bg-pink-400 sm:text-[10px]"
                @click="launchConfetti(false, false)"
            >
                Confetti
            </button>
            <button
                class="card overflow-hidden text-clip whitespace-nowrap p-2 text-xs hover:!bg-pink-400 sm:text-[10px]"
                @click="playWinnerSound"
            >
                Win Sound
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
import {
    getAndSetNewSecretPokemon,
    updateCurrentUserStreakDisplay,
    getAndSetCurrentSecretPokemon,
} from './services/game.js';
import { clearLocalStorageGameMode } from './services/localStorage';
import { playWinnerSound } from './services/sound';
import { GameModes, DateOfFirstPokeldeGameClassic } from './constants';
import { setLanguageFromLanguageCode } from './services/language.js';

const store = useStore();
const isLoaded = ref(false);
const isDevelopment = computed(() => ENVIRONMENT === 'development');

onBeforeMount(async () => {
    setLanguageFromLanguageCode();
    store.loadTheme();
    loadIsHintMode();
    loadIsShiny();
    loadIsSound();

    const user = await getOrCreateUser();

    console.log('Loaded at: ' + moment().toString());
    console.log('ENVIRONMENT: ' + ENVIRONMENT);
    console.log(user);
    console.log('DateOfFirstPokeldeGameClassic: ' + DateOfFirstPokeldeGameClassic.toString());
    console.log(moment().toString());

    if (user) {
        store.setUser(user);
    }
    if (!store.currentStreak) updateCurrentUserStreakDisplay(GameModes.Classic);

    isLoaded.value = true;
});

const getOrCreateUser = async () => {
    const userId = localStorage.userId;

    console.log('User From Storage: ' + userId);
    if (userId) {
        await apiService.updateUserStreaks(userId);
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

const loadIsSound = () => {
    if (localStorage.isSound === 'false') {
        store.setSound(false);
    } else {
        store.setSound(true);
    }
};

const clearLocalStorageAllGameModes = () => {
    clearLocalStorageGameMode(GameModes.Classic);
    clearLocalStorageGameMode(GameModes.Flavortext);
    clearLocalStorageGameMode(GameModes.Silhouette);
};

//Get most resent db entries
const getCurrentGame = async () => {
    clearLocalStorageAllGameModes();
    await Promise.all([
        getAndSetCurrentSecretPokemon(GameModes.Classic),
        getAndSetCurrentSecretPokemon(GameModes.Flavortext),
        getAndSetCurrentSecretPokemon(GameModes.Silhouette),
    ]);

    location.reload();
};

// Force update pokemon in DB
const setNewGame = async () => {
    clearLocalStorageAllGameModes();
    await Promise.all([
        getAndSetNewSecretPokemon(GameModes.Classic),
        getAndSetNewSecretPokemon(GameModes.Flavortext),
        getAndSetNewSecretPokemon(GameModes.Silhouette),
    ]);

    location.reload();
};

const revealPokemon = async () => {
    const [responseClassic, responseFlavortext, responseSilhouette] = await Promise.all([
        apiService.getClassicSecretPokemon(),
        apiService.getFlavortextSecretPokemon(),
        apiService.getSilhouetteSecretPokemon(),
    ]);

    const classicSecret = JSON.parse(localStorage.classicSecretPokemon);
    console.log(`Classic Secret: LS: ${classicSecret.name}, DB: ${responseClassic.data.name}`);

    const silhouetteSecret = JSON.parse(localStorage.silhouetteSecretPokemon);
    console.log(`Silhouette Secret: LS: ${silhouetteSecret.name}, DB: ${responseSilhouette.data.name}`);

    const flavortextSecret = JSON.parse(localStorage.flavortextSecretPokemon);
    console.log(`Flavortext Secret: LS: ${flavortextSecret.name}, DB: ${responseFlavortext.data.name}`);
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
