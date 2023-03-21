<template>
    <a :href="baseTweetURL+encodeURIComponent(twitterText)">
    <div
        class="card cursor-pointer transition-all duration-100 ease-linear hover:scale-110 gap-1 p-1 px-2 sm:py-0 !bg-twitterblue text-white"
        >
        <SvgIcon type="mdi" :path="mdiTwitter" class="w-[24px] sm:w-[20px]" />
        <p class="font-pkmEmerald pt-1 text-sm sm:text-xs sm:font-bold">SHARE RESULTS!</p>
    </div>
</a>
</template>

<script setup>
import SvgIcon from '@jamescoyle/vue-icon';
import { mdiTwitter } from '@mdi/js';
var baseTweetURL = 'https://twitter.com/intent/tweet?text=';

const props = defineProps({
    twitterText: String
})
</script>
