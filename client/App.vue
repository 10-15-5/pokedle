<template>
    <div id="app">
        <div
            class="min-h-screen w-full"
            :class="{
                'background-white': !(isDark && store.isDifficultyInsane),
                'background-black': isDark && !store.isDifficultyInsane,
                'background-white-difficulty-insane': !isDark && store.isDifficultyInsane,
                'background-black-difficulty-insane': isDark && store.isDifficultyInsane,
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
            <RouterView />
        </div>
    </div>
</template>

<script setup>
import HeaderContainer from './components/headerIcons/HeaderIconContainer.vue';
import ThemeButton from './components/buttons/ThemeButton.vue';
import { reactive, ref, onBeforeMount } from 'vue';
import * as apiService from './services/api/apiService.js';
import { useStore } from './stores/store.js';
import { useDark } from '@vueuse/core';

const isDark = useDark();
const store = useStore();

onBeforeMount(async () => {
    const user = await getOrCreateUser();
    console.log(user);

    if (user) {
        store.setUser(user);
    }

    loadIsShiny();
    loadIsDifficultyInsane();
});

const getOrCreateUser = async () => {
    const userId = localStorage.getItem('userId');

    console.log(userId);
    if (userId) {
        await apiService.updateUserStreak(userId);
        const response = await apiService.getUser(userId);
        return response.data.user;
    }

    const response = await apiService.createUser();
    const { user } = response.data;
    localStorage.setItem('userId', user._id);
    return user;
};

const loadIsShiny = () => {
    const isShinyString = localStorage.getItem('isShiny');
    if (isShinyString && isShinyString === 'true') {
        store.setShiny(true);
        return;
    }
};

const loadIsDifficultyInsane = () => {
    const isDifficultyInsaneString = localStorage.getItem('isDifficultyInsane');
    console.log(isDifficultyInsaneString);
    if (isDifficultyInsaneString && isDifficultyInsaneString === 'true') {
        store.setDifficultyInsane(true);
        return;
    }
};
</script>

<style scoped>
html.dark {
    color-scheme: dark;
}

.background-white {
    background-image: url('./client/assets/backgrounds/background-white.png');
    background-size: cover;
    background-repeat: no-repeat;
    background-attachment: fixed;
}

.background-black {
    background: linear-gradient(rgba(0, 0, 0, 0.25), rgba(0, 0, 0, 0.25)),
        url('./client/assets/backgrounds/background-black.jpg');
    background-size: cover;
    background-repeat: no-repeat;
    background-attachment: fixed;
}

.background-white-difficulty-insane {
    background-image: url('./client/assets/backgrounds/background-white-trainer-red.jpeg');
    background-size: cover;
    background-repeat: no-repeat;
    background-attachment: fixed;
}

.background-black-difficulty-insane {
    background-image: url('./client/assets/backgrounds/background-black-cubone.jpeg');
    background-size: cover;
    background-repeat: no-repeat;
    background-attachment: fixed;
}
</style>
