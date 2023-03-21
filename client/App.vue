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
import moment from 'moment-timezone';

const store = useStore();

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

const loadIsHintMode = () => {
    if (localStorage.isHintMode === 'true') {
        store.setHintMode(true);
    } else {
        store.setHintMode(false);
    }
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
