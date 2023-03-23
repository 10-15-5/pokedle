<template>
    <div
        class="card flex w-[180px] flex-row items-center justify-between gap-2 px-2 pt-2 pb-1 text-2xl transition-all duration-100 ease-linear hover:scale-110 !font-pkmEmerald"
    >
        {{ gameMode }}
        <SvgIcon v-if="gameMode === GameModes.Classic" type="mdi" class="mt-[-4px] h-10 w-10" :path="mdiHelpCircle" />
        <SvgIcon
            v-else-if="gameMode === GameModes.Flavortext"
            type="mdi"
            class="mt-[-4px] h-10 w-10"
            :path="mdiChatQuestion"
        />
    </div>
</template>

<script setup>
import SvgIcon from '@jamescoyle/vue-icon';
import { mdiHelpCircle, mdiChatQuestion } from '@mdi/js';
import { GameModes } from '../../constants';
const props = defineProps({
    gameMode: String,
});
</script>
