<template>
    <div class="card sm:pl-0 card gap-x-2 sm:gap-x-[2px] py-1 static">
        <p v-for="(fieldTitle, _, index) in guessFieldTitles"
           :key="fieldTitle"
           :value="fieldTitle"
           class="w-[60px] sm:w-[45px] h-[40px] sm:h-[30px] flex justify-center items-center border-b-2
                            border-light-border dark:border-dark-border
                             font-pkmEmerald font-bold text-center group text-sm sm:text-[12px] leading-4">{{ fieldTitle }}
        <div v-if="getGuessFieldTooltipsFromIndex(index)"
             class="card transition-all duration-150 scale-0 origin-top absolute
                        group-hover:scale-100 z-1 mb-[100px] p-2 delay-500
                "> {{ getGuessFieldTooltipsFromIndex(index) }} </div>
        </p>
    </div>
</template>

<script setup>
import { guessFieldTitles, getGuessFieldTooltipsFromIndex } from '../constants.js';


</script>

<style scoped>
</style>