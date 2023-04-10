<template>
    <div
        class="card relative flex w-[180px] flex-row items-center justify-between gap-2 px-2 pt-2 pb-1 !font-pkmEmerald text-2xl transition-all duration-100 ease-linear hover:scale-110 sm:w-[150px] sm:text-xl"
    >
        <span class="relative">{{ text().gameModes[gameMode] }}</span>
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
        <CornerSpriteAnimation v-if="isAttention" />
    </div>
</template>

<script setup>
import SvgIcon from '@jamescoyle/vue-icon';
import { mdiHelpCircle, mdiMessageQuestion, mdiAccountQuestion } from '@mdi/js';
import { GameModes } from '../../constants';
import CornerSpriteAnimation from '../animations/CornerSpriteAnimation.vue';
import { text } from '../../services/language';
const props = defineProps({
    gameMode: String,
    isAttention: Boolean,
});
</script>
