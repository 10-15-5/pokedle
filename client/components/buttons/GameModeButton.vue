<template>
    <div
        class="card relative flex w-[180px] flex-row items-center justify-between gap-2 px-2 pt-2 pb-1 !font-pkmEmerald text-2xl transition-all duration-100 ease-linear hover:scale-110 sm:w-[150px] sm:text-xl"
    >
        <span class="relative">{{ gameMode }}</span>
        <SvgIcon
            v-if="gameMode === GameModes.Classic"
            type="mdi"
            class="mt-[-4px] h-11 w-11 text-teal-500 sm:h-9 sm:w-9"
            :path="mdiHelpCircle"
        />
        <SvgIcon
            v-else-if="gameMode === GameModes.Flavortext"
            type="mdi"
            class="mt-[-4px] h-11 w-11 text-teal-500 sm:h-9 sm:w-9"
            :path="mdiMessageQuestion"
        />
        <SvgIcon
            v-else-if="gameMode === GameModes.Silhouette"
            type="mdi"
            class="mt-[-4px] h-11 w-11 text-teal-500 sm:h-9 sm:w-9"
            :path="mdiAccountQuestion"
        />

        <div v-if="isAttention" class="absolute right-0 top-0 mt-[-7px] mr-[-8px]">
            <span class="relative flex h-4 w-4">
                <span class="absolute inline-flex h-full w-full animate-ping rounded-full bg-sky-400 opacity-75"></span>
                <span class="relative inline-flex h-4 w-4 rounded-full bg-teal-300"></span>
            </span>
        </div>
    </div>
</template>

<script setup>
import SvgIcon from '@jamescoyle/vue-icon';
import { mdiHelpCircle, mdiMessageQuestion, mdiAccountQuestion } from '@mdi/js';
import { GameModes } from '../../constants';
const props = defineProps({
    gameMode: String,
    isAttention: Boolean,
});
</script>
