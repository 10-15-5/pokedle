<template>
    <button class="toggle-button" @click="store.toggleTheme()">
        <SvgIcon v-if="store.isDark" type="mdi" class="h-9 w-9" :path="mdiWeatherSunny" />
        <SvgIcon v-else type="mdi" class="h-9 w-9" :path="mdiWeatherNight" />
    </button>
</template>

<script setup>
import SvgIcon from '@jamescoyle/vue-icon';
import { mdiWeatherNight, mdiWeatherSunny } from '@mdi/js';
import { useStore } from '../../stores/store';
const store = useStore();
</script>
