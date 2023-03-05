<template>
    <BaseDialog :isVisible="isVisible" :onClose="onClose">
        <h1 class="font-pkmEmerald text-[32px] uppercase">About</h1>
        <p class="font-pkmEmerald text-[20px]">
            Welcome to the world of POKEMéN! <br />
            Every day a wild secret POKéMON will be hiding and its your job as a POKéMON TRAINER to
            catch (guess) it!
        </p>
        <hr class="divider" />
        <h1 class="font-pkmEmerald text-[32px] uppercase">How To Play</h1>
        <div class="font-pkmEmerald text-[20px]">
            <p>
                <b>1. Guess:</b> Guess any random POKEMéN <b>from gen 1-3</b>, since we currently
                have no clues. We will guess <b>SQUIRTLE</b> here.
            </p>
            <p>
                <b>2. Guess:</b> We now know that we are looking for a <b>BLUE</b> POKEMéN.
                Furthermore the POKEMéN has to be in its <b>final evolution stage</b>. Lastly, the
                POKEMéN is <b>not</b> from gen 1. <br />
                Let's guess <b> HERACROSS </b> since it matches these characteristics.
            </p>
            <p>
                <b>3. Guess:</b> Now we know the hidden POKEMéN <b>must</b> be from gen 3. Since the
                hidden POKEMéN cannot be a <b>WATER TYPE</b>, this leaves only a couple of possible
                options. <br />
                Let's guess
                <b>SALAMANCE.</b>
            </p>
            <p>
                <b>4. Guess:</b> We now know that the POKEMéN lives in <b>ROUGH TERRAIN</b> and is
                in <b>evolution stage 3</b> <br />
                This leaves only one POKEMéN; <b>METAGROSS.</b>
            </p>
            <v-img class="card mx-auto w-[600px]" :src="'./client/assets/game-example.png'"></v-img>
        </div>
        <hr class="divider" />
        <h1 class="font-pkmEmerald text-[32px] uppercase">Credit</h1>
        <p class="font-pkmEmerald text-[20px]">
            Pokédlé is heavily inspired by loldle.net and wordle. Original artist/creator of images,
            please DM for removal or credit.
        </p>
        <hr class="divider" />
        <h1 class="font-pkmEmerald text-[32px] uppercase">Contact</h1>
        <p class="font-pkmEmerald text-[20px]">
            If you have any questions/feedback/bug reports feel free to contact me on the email
            below. <br />
            <a
                class="text-[24px] font-bold text-blue-500 hover:text-blue-400 dark:text-blue-300 dark:hover:text-blue-400"
                href="mailto: pokedle.contact@gmail.com"
            >
                pokedle.contact@gmail.com
            </a>
        </p>
    </BaseDialog>
</template>

<script setup>
import BaseDialog from './BaseDialog.vue';

const props = defineProps({
    isVisible: Boolean,
    onClose: Function,
});
</script>
