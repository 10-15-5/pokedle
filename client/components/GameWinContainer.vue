<template>
    <div class="card flex flex-col items-center justify-center p-4">
        <p class="font-pkmEmerald text-[32px] !text-light-green dark:!text-dark-green">VICTORY!</p>
        <p class="mb-3 font-pkmEmerald text-[18px]">You Guessed: {{ pokemon }}</p>
        <SquareContent
            :pokemon="removeSpecialCharactersExceptDashFromString(pokemon)"
            :color="props.color"
        />
        <p class="mt-2 font-pkmEmerald text-[14px]">
            Next Pokemon Will Appear In:
            {{ ' ' + hoursRemaining + ':' + minRemaining + ':' + secRemaining }}
        </p>
    </div>
</template>

<script setup>
import { onBeforeMount, onMounted, ref } from 'vue';
import SquareContent from './SquareContent.vue';
import { removeSpecialCharactersExceptDashFromString } from '../helpers';

const props = defineProps({
    pokemon: String,
    color: String,
});

const secRemaining = ref('');
const minRemaining = ref('');
const hoursRemaining = ref('');
const secondsRemaining = ref(0);

const setSecRemaining = () => {
    const countDownToDate = new Date(new Date().setHours(0, 0, 0));
    countDownToDate.setDate(countDownToDate.getDate() + 1);
    secondsRemaining.value = Math.floor((countDownToDate.getTime() - Date.now()) / 1000);
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

<style scoped>
</style>
