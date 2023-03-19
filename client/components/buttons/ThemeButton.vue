<template>
    <button
        class="toggle-button"
        @click="store.toggleTheme()"
    >
        <SvgIcon v-if="store.isDark" type="mdi" class="w-9 h-9" :path="mdiWeatherSunny"/>
        <SvgIcon v-else type="mdi" class="w-9 h-9" :path="mdiWeatherNight"/>
    </button>
</template>

<script setup>
import SvgIcon from '@jamescoyle/vue-icon'
import { mdiWeatherNight, mdiWeatherSunny } from '@mdi/js'
import { useStore } from '../../stores/store';
const store = useStore();
</script>
