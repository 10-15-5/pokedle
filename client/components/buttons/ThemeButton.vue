<template>
    <button
        class="toggle-button"
        @click="toggleDark()"
    >
        <SvgIcon v-if="isDark" type="mdi" class="w-9 h-9" :path="mdiWeatherSunny"/>
        <SvgIcon v-else type="mdi" class="w-9 h-9" :path="mdiWeatherNight"/>
    </button>
</template>

<script setup>
import { useDark, useToggle } from '@vueuse/core';
import SvgIcon from '@jamescoyle/vue-icon'
import { mdiWeatherNight, mdiWeatherSunny } from '@mdi/js'

const isDark = useDark();
const toggleDark = useToggle(isDark);
</script>
