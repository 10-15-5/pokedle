<template>
    <!-- TODO: UNUSED!! -->
    <a :href="baseTweetURL + encodeURIComponent(twitterText)">
        <div
            class="card cursor-pointer gap-1 !bg-twitterblue p-1 px-2 text-white transition-all duration-100 ease-linear hover:scale-110 sm:py-0"
        >
            <SvgIcon type="mdi" :path="mdiTwitter" class="w-[24px] sm:w-[20px]" />
            <p class="pt-1 sm:pt-[5px] font-pkmEmerald text-sm sm:text-xs sm:font-bold">SHARE RESULTS!</p>
        </div>
    </a>
</template>

<script setup>
import SvgIcon from '@jamescoyle/vue-icon';
import { mdiTwitter } from '@mdi/js';
var baseTweetURL = 'https://twitter.com/intent/tweet?text=';

const props = defineProps({
    twitterText: String,
});
</script>
