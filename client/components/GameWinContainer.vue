<template>
    <div class="card flex flex-col p-4 justify-center items-center">
        <p class="text-[32px] font-pkmEmerald  text-green-600 dark:!text-green-300">VICTORY!</p>
        <p class="text-[18px] font-pkmEmerald mb-3"> You Guessed: {{ pokemon }}</p>
        <SquareContent :pokemon="removeSpecialCharactersExceptDashFromString(pokemon)"
                       :color="props.color" />
        <p class="text-[14px] font-pkmEmerald mt-2">Next Pokemon Will Appear In: {{ " " + hoursRemaining + ":" + minRemaining + ":" + secRemaining }}</p>
    </div>
</template>

<script setup>
import { onBeforeMount, onMounted, ref } from 'vue';
import SquareContent from './SquareContent.vue';
import { removeSpecialCharactersExceptDashFromString } from '../helpers';

const props = defineProps({
    pokemon: String,
    color: String
})

const secRemaining = ref("")
const minRemaining = ref("")
const hoursRemaining = ref("")
const secondsRemaining = ref(0)


const setSecRemaining = () => {
    const countDownToDate = new Date(new Date().setHours(0, 0, 0));
    countDownToDate.setDate(countDownToDate.getDate() + 1);
    secondsRemaining.value = Math.floor((countDownToDate.getTime() - Date.now()) / 1000);
}

onBeforeMount(() => {
    setSecRemaining();
    updateTimeRemaining();
});

const updateTimeRemaining = () => {
    secondsRemaining.value--;

    if (secondsRemaining.value <= 0) {
        secRemaining.value = "00";
        minRemaining.value = "00";
        hoursRemaining.value = "00";

        return clearInterval(updateTimeRemaining);
    }

    const sec = Math.floor(secondsRemaining.value % 60);
    const min = Math.floor((secondsRemaining.value / 60) % 60);
    const hours = Math.floor((secondsRemaining.value / (60 * 60)) % 60);
    secRemaining.value = sec < 10 ? "0" + sec : "" + sec;
    minRemaining.value = min < 10 ? "0" + min : "" + min;
    hoursRemaining.value = hours < 10 ? "0" + hours : "" + hours;

}
setInterval(() => {
    updateTimeRemaining();
}, 1000)

</script>

<style scoped>
.card-text,
.title,
.card-smaller-text {
    font-family: pkmEmerald;
    text-transform: capitalize;
}

.card-text {
    font-size: 18px;
}

.card-smaller-text {
    font-size: 14px;
}

</style>