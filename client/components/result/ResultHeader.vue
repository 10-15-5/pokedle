<template>
    <div class="card gap-x-2 p-2 py-1 sm:gap-x-[2px] sm:space-x-[2px] sm:!px-1 sm:!py-[2px]">
        <div
            v-for="(fieldTitle, _, index) in guessFieldTitles"
            :key="fieldTitle"
            :value="fieldTitle"
            class="group flex h-[40px] w-[60px] items-center justify-center border-b-2 border-light-border text-center font-pkmEmerald text-sm font-bold leading-4 dark:border-dark-border sm:h-[30px] sm:w-[43px] sm:text-[10px] !m-0 sm:leading-3"
        >
            {{ fieldTitle }}
            <div
                v-if="getGuessFieldTooltipsFromIndex(index)"
                class="card z-1 absolute mb-[100px] origin-top scale-0 p-2 transition-all delay-500 duration-150 group-hover:scale-100 sm:mb-[66px] sm:p-1"
            >
                {{ getGuessFieldTooltipsFromIndex(index) }}
            </div>
        </div>
    </div>
</template>

<script setup>
import { guessFieldTitles, getGuessFieldTooltipsFromIndex } from '../../constants.js';
</script>

<style scoped></style>
