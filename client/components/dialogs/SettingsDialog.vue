<template>
    <BaseDialog :isVisible="isVisible" :onClose="onClose">
        <h1 class="text-center font-pkmEmerald text-[32px] uppercase">SETTINGS</h1>
        <div class="flex flex-col items-center gap-y-4 sm:gap-x-4">
            <div class="flex flex-row w-full justify-between px-20 sm:px-6">
                <p class="font-pkmEmerald text-[24px]">DarkMode:</p>
                <BaseToggle :isEnabled="isDark" :toggle="toggleDark">
                    <template #leftIcon>
                        <SvgIcon type="mdi" :path="mdiWeatherNight" class="h-[30px] w-[30px]" />
                    </template>
                    <template #rightIcon>
                        <SvgIcon
                            type="mdi"
                            :path="mdiWeatherSunny"
                            class="h-[30px] w-[30px]"
                        /> </template
                ></BaseToggle>
            </div>
            <div class="flex flex-row w-full justify-between px-20 sm:px-6">
                <p class="font-pkmEmerald text-[24px]">Shiny Only:</p>
                <BaseToggle :isEnabled="store.isShiny" :toggle="store.toggleShiny">
                    <template #leftIcon>
                        <SvgIcon type="mdi" :path="mdiStarOutline" class="h-[30px] w-[30px]" />
                    </template>
                    <template #rightIcon>
                        <SvgIcon
                            type="mdi"
                            :path="mdiStarOffOutline"
                            class="h-[30px] w-[30px]"
                        /> </template
                ></BaseToggle>
            </div>
            <div class="flex flex-row w-full justify-between px-20 sm:px-6">
                <p class="font-pkmEmerald text-[24px]">Champion Difficulty:</p>
                <BaseToggle
                    :isEnabled="store.isDifficultyInsane"
                    :toggle="store.toggleDifficultyInsane"
                >
                    <template #leftIcon>
                        <SvgIcon type="mdi" :path="mdiCoffin" class="h-[30px] w-[30px]" />
                    </template>
                    <template #rightIcon>
                        <SvgIcon
                            type="mdi"
                            :path="mdiEmoticonExcitedOutline"
                            class="h-[30px] w-[30px]"
                        /> </template
                ></BaseToggle>
            </div>
        </div>
    </BaseDialog>
</template>

<script setup>
import BaseDialog from './BaseDialog.vue';
import BaseToggle from '../buttons/BaseToggle.vue';
import SvgIcon from '@jamescoyle/vue-icon';
import {
    mdiStarOffOutline,
    mdiStarOutline,
    mdiWeatherNight,
    mdiWeatherSunny,
    mdiCoffin,
    mdiEmoticonExcitedOutline,
} from '@mdi/js';
import { useStore } from '../../stores/store.js';
import { useDark, useToggle } from '@vueuse/core';
const store = useStore();
const isDark = useDark();
const toggleDark = useToggle(isDark);

const props = defineProps({
    isVisible: Boolean,
    onClose: Function,
});
</script>

<style></style>
