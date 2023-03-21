<template>
    <div class="card flex flex-col items-center justify-center p-5 py-2 sm:p-3">
        <p class="fontbo font-pkmEmerald text-[32px] !text-light-emerald dark:!text-dark-emerald">
            VICTORY!
        </p>
        <p class="mb-3 font-pkmEmerald text-[18px] capitalize">
            You Guessed: <b class="!text-light-emerald dark:!text-dark-emerald"> {{ pokemon }} </b>
        </p>
        <SquareContent
            :pokemon="removeSpecialCharactersExceptDashFromString(pokemon)"
            :color="props.color"
            :type="guessType.Pokemon"
        />
        <p class="mt-2 font-pkmEmerald text-[14px]">
            Next Pokemon Will Appear In :&nbsp;
            <b class="text-light-orange dark:!text-dark-orange">
                {{ ' ' + hoursRemaining + ' : ' + minRemaining + ' : ' + secRemaining }}</b
            >
        </p>
        <TweetButton class="mt-3" :twitterText="twitterText" />
    </div>
</template>

<script setup>
import { onBeforeMount, ref } from 'vue';
import SquareContent from './result/ResultSquare.vue';
import { removeSpecialCharactersExceptDashFromString } from '../helpers';
import { guessType } from '../constants';
import TweetButton from './buttons/TweetButton.vue';
import moment from 'moment-timezone';

const props = defineProps({
    pokemon: String,
    color: String,
    twitterText: String,
});

const secRemaining = ref('');
const minRemaining = ref('');
const hoursRemaining = ref('');
const secondsRemaining = ref(0);

const setSecRemaining = () => {
    const endOfToday = moment().endOf('day');
    const msLeft = endOfToday - moment();
    secondsRemaining.value = Math.floor(msLeft / 1000);
};

onBeforeMount(() => {
    setSecRemaining();
    updateTimeRemaining();
});

const updateTimeRemaining = () => {
    secondsRemaining.value--;

    if (secondsRemaining.value <= 0) {
        secRemaining.value = '00';
        minRemaining.value = '00';
        hoursRemaining.value = '00';

        return clearInterval(updateTimeRemaining);
    }

    const sec = Math.floor(secondsRemaining.value % 60);
    const min = Math.floor((secondsRemaining.value / 60) % 60);
    const hours = Math.floor((secondsRemaining.value / (60 * 60)) % 60);
    secRemaining.value = sec < 10 ? '0' + sec : '' + sec;
    minRemaining.value = min < 10 ? '0' + min : '' + min;
    hoursRemaining.value = hours < 10 ? '0' + hours : '' + hours;
};
setInterval(() => {
    updateTimeRemaining();
}, 1000);
</script>

<style scoped></style>
