<template>
    <button
        class="toggle-button hover:!text-cyan-400 dark:hover:!text-amber-400"
        @click="toggleDark()"
    >
        <SvgIcon v-if="isDark" type="mdi" class="w-9 h-9" :path="mdiWhiteBalanceSunny"/>
        <SvgIcon v-else type="mdi" class="w-9 h-9" :path="mdiWeatherNight"/>
    </button>
</template>

<script setup>
import { useDark, useToggle } from '@vueuse/core';
import SvgIcon from '@jamescoyle/vue-icon'
import { mdiWeatherNight, mdiWhiteBalanceSunny } from '@mdi/js'

const isDark = useDark();
const toggleDark = useToggle(isDark);
</script>
