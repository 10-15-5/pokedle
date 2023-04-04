<template>
    <BaseDialog :isVisible="isVisible" :onClose="onClose">
        <h1 class="text-center font-pkmEmerald text-[32px] uppercase">{{ getLanguage().settingsDialog.settings }}</h1>
        <div class="flex flex-col items-center gap-y-4 sm:gap-x-0">
            <div class="flex w-full flex-row justify-between px-20 sm:px-0">
                <p class="font-pkmEmerald text-[24px]">{{ getLanguage().settingsDialog.darkMode }}:</p>
                <BaseToggle :isEnabled="store.isDark" :toggle="store.toggleTheme">
                    <template #leftIcon>
                        <SvgIcon type="mdi" :path="mdiWeatherNight" class="h-[30px] w-[30px]" />
                    </template>
                    <template #rightIcon>
                        <SvgIcon type="mdi" :path="mdiWeatherSunny" class="h-[30px] w-[30px]" /> </template
                ></BaseToggle>
            </div>
            <div class="flex w-full flex-row justify-between px-20 sm:px-0">
                <p class="font-pkmEmerald text-[24px]">{{ getLanguage().settingsDialog.shinyOnly }}:</p>
                <BaseToggle :isEnabled="store.isShiny" :toggle="store.toggleShiny">
                    <template #leftIcon>
                        <SvgIcon type="mdi" :path="mdiStarOutline" class="h-[30px] w-[30px]" />
                    </template>
                    <template #rightIcon>
                        <SvgIcon type="mdi" :path="mdiStarOffOutline" class="h-[30px] w-[30px]" /> </template
                ></BaseToggle>
            </div>
            <div class="flex w-full flex-row justify-between px-20 sm:px-0">
                <p class="font-pkmEmerald text-[24px]">{{ getLanguage().settingsDialog.hintMode }}:</p>
                <BaseToggle :isEnabled="store.isHintMode" :toggle="store.toggleHintMode">
                    <template #leftIcon>
                        <SvgIcon type="mdi" :path="mdiFlowerOutline" class="h-[30px] w-[30px]" />
                    </template>
                    <template #rightIcon>
                        <SvgIcon type="mdi" :path="mdiSwordCross" class="h-[30px] w-[30px]" /> </template
                ></BaseToggle>
            </div>
            <div class="flex w-full flex-row justify-between px-20 sm:px-0">
                <p class="font-pkmEmerald text-[24px]">{{ getLanguage().settingsDialog.sound }}:</p>
                <BaseToggle :isEnabled="store.isSound" :toggle="store.toggleSound">
                    <template #leftIcon>
                        <SvgIcon type="mdi" :path="mdiVolumeHigh" class="h-[30px] w-[30px]" />
                    </template>
                    <template #rightIcon>
                        <SvgIcon type="mdi" :path="mdiVolumeOff" class="h-[30px] w-[30px]" /> </template
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
    mdiFlowerOutline,
    mdiSwordCross,
    mdiVolumeHigh,
    mdiVolumeOff,
} from '@mdi/js';
import { useStore } from '../../stores/store.js';
import { getLanguage } from '../../services/language';

const store = useStore();

const props = defineProps({
    isVisible: Boolean,
    onClose: Function,
});
</script>

<style></style>
