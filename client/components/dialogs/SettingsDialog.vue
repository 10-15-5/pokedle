<template>
    <BaseDialog :isVisible="isVisible" :onClose="onClose">
        <h1 class="text-center font-pkmEmerald text-[32px] uppercase">{{ text().settingsDialog.settings }}</h1>
        <div class="flex flex-col items-center gap-y-4 sm:gap-x-0">
            <!-- <div class="flex w-full flex-row justify-between px-20 sm:px-0">
                <p class="font-pkmEmerald text-[24px]">{{ text().settingsDialog.language }}:</p>
                <Dropdown :selectedItem="selectedLanguage" :items="languages" @onSelectItem="setNewLanguage" />
            </div> -->
            <div class="flex w-full flex-row justify-between px-20 sm:px-0">
                <p class="font-pkmEmerald text-[24px]">{{ text().settingsDialog.sound }}:</p>
                <BaseToggle :isEnabled="store.isSound" :toggle="store.toggleSound">
                    <template #leftIcon>
                        <SvgIcon type="mdi" :path="mdiVolumeHigh" class="h-[30px] w-[30px]" />
                    </template>
                    <template #rightIcon>
                        <SvgIcon type="mdi" :path="mdiVolumeOff" class="h-[30px] w-[30px]" /> </template
                ></BaseToggle>
            </div>
            <div class="flex w-full flex-row justify-between px-20 sm:px-0">
                <p class="font-pkmEmerald text-[24px]">{{ text().settingsDialog.darkMode }}:</p>
                <BaseToggle :isEnabled="store.isDark" :toggle="store.toggleTheme">
                    <template #leftIcon>
                        <SvgIcon type="mdi" :path="mdiWeatherNight" class="h-[30px] w-[30px]" />
                    </template>
                    <template #rightIcon>
                        <SvgIcon type="mdi" :path="mdiWeatherSunny" class="h-[30px] w-[30px]" /> </template
                ></BaseToggle>
            </div>
            <div class="flex w-full flex-row justify-between px-20 sm:px-0">
                <p class="font-pkmEmerald text-[24px]">{{ text().settingsDialog.shinyOnly }}:</p>
                <BaseToggle :isEnabled="store.isShiny" :toggle="store.toggleShiny">
                    <template #leftIcon>
                        <SvgIcon type="mdi" :path="mdiStarOutline" class="h-[30px] w-[30px]" />
                    </template>
                    <template #rightIcon>
                        <SvgIcon type="mdi" :path="mdiStarOffOutline" class="h-[30px] w-[30px]" /> </template
                ></BaseToggle>
            </div>
            <div class="flex w-full flex-row justify-between px-20 sm:px-0">
                <p class="font-pkmEmerald text-[24px]">{{ text().settingsDialog.hintMode }}:</p>
                <BaseToggle :isEnabled="store.isHintMode" :toggle="store.toggleHintMode">
                    <template #leftIcon>
                        <SvgIcon type="mdi" :path="mdiFlowerOutline" class="h-[30px] w-[30px]" />
                    </template>
                    <template #rightIcon>
                        <SvgIcon type="mdi" :path="mdiSwordCross" class="h-[30px] w-[30px]" /> </template
                ></BaseToggle>
            </div>
        </div>
    </BaseDialog>
</template>

<script setup>
import BaseDialog from './BaseDialog.vue';
import BaseToggle from '../buttons/BaseToggle.vue';
import Dropdown from '../dropdowns/Dropdown.vue';
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
import { text } from '../../services/language';
import { getLanguage, setLanguage, languages } from '../../services/language';
import { ref } from 'vue';

const setNewLanguage = (lang) => {
    if (setLanguage(lang)) {
        selectedLanguage.value = lang;
    }
};

const selectedLanguage = ref(getLanguage());
const store = useStore();
//'korean', 'french', 'german', 'spanish', 'italian', 'english', 'japanese'
const props = defineProps({
    isVisible: Boolean,
    onClose: Function,
});
</script>
