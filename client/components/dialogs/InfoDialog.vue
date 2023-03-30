<template>
    <BaseDialog :isVisible="isVisible" :onClose="onClose">
        <div class="relative max-w-[700px] font-pkmEmerald text-[20px]">
            <h1 class="mb-2 text-center text-[32px] uppercase">About</h1>
            <div class="relative flex flex-row items-end">
                <span class="relatives w-fit text-center font-pkmEmerald text-[20px]">
                    Welcome to the world of POKéMON! </span
                ><SpriteAnimation />
            </div>
            <p>
                Every day a wild POKéMON from gen 1-3 will be hiding and its your job as a POKéMON TRAINER to catch
                (guess) it!
            </p>
            <hr class="divider" />
            <h1 class="mb-2 text-center font-pkmEmerald text-[32px] uppercase">How To Play</h1>
            <p class="my-2">
                <b>1. Guess:</b> Guess any random POKéMON <b>from gen 1-3</b>, since we currently have no clues. We will
                guess <b>SQUIRTLE</b> here.
            </p>
            <p class="my-2">
                <b>2. Guess:</b> We now know that we are looking for a <b>BLUE</b> POKéMON. Furthermore the POKéMON has
                to be in its <b>final evolution stage</b>. Lastly, the POKéMON is <b>not</b> from gen 1. <br />
                Let's guess <b> HERACROSS </b> since it matches these characteristics.
            </p>
            <p class="my-2">
                <b>3. Guess:</b> Now we know the hidden POKéMON <b>must</b> be from gen 3. Since the hidden POKéMON
                cannot be a <b>WATER TYPE</b>, this leaves only a couple of possible options. <br />
                Let's guess
                <b>SALAMANCE.</b>
            </p>
            <p class="my-2">
                <b>4. Guess:</b> We now know that the POKéMON lives in <b>ROUGH TERRAIN</b> and is in
                <b>evolution stage 3</b> <br />
                This leaves only one POKéMON; <b>METAGROSS.</b>
            </p>
            <img class="card mx-auto w-[600px]" :src="'./client/assets/game-examples/game-example.png'" />

            <hr class="divider" />
            <h1 class="mb-2 text-center font-pkmEmerald text-[32px] uppercase">Habitats</h1>
            <p>
                The Habitats used are all taken from
                <a class="link" href="https://bulbapedia.bulbagarden.net/wiki/List_of_Pok%C3%A9mon_by_habitat">
                    Bulbapedia
                </a>
                and include:
            </p>
            <div class="grid grid-flow-col gap-1 p-2 sm:grid-flow-col sm:grid-rows-3">
                <img :src="'./client/assets/habitats/cave.png'" title="cave" />
                <img :src="'./client/assets/habitats/forest.png'" title="forest" />
                <img :src="'./client/assets/habitats/grassland.png'" title="grassland" />
                <img :src="'./client/assets/habitats/mountain.png'" title="mountain" />
                <img :src="'./client/assets/habitats/rare.png'" title="rare" />
                <img :src="'./client/assets/habitats/rough-terrain.png'" title="rough-terrain" />
                <img :src="'./client/assets/habitats/sea.png'" title="sea" />
                <img :src="'./client/assets/habitats/urban.png'" title="urban" />
                <img :src="'./client/assets/habitats/waters-edge.png'" title="waters-edge" />
            </div>

            <hr class="divider" />
            <h1 class="mb-2 text-center font-pkmEmerald text-[32px] uppercase">Contribute</h1>
            <p class="leading-8">
                Willing to contribute? Thats great! 🥳 You can find info on how to reach me at the bottom of this card.
                <br />
                At the time of this writing, some of the things I could use help with are:
            </p>
            <ul class="ml-5 list-disc">
                <li>Animations 🎥</li>
                <li>Translation 🇯🇵 🇪🇸 🇮🇳 🇨🇳 🇩🇪 🇫🇷</li>
                <li>Design / Styling 🧑‍🎨 🎨</li>
                <li>Feature Development 👷‍♂️</li>
            </ul>
            <hr class="divider" />
            <h1 class="mb-2 text-center font-pkmEmerald text-[32px] uppercase">Credit</h1>
            <div class="w-full leading-8">
                <p class="mb-2">Pokédle is created by Gabriel Haugbøl. I hope you enjoy playing along 😊🎮!</p>
                <p class="mb-2">
                    Pokédle is heavily inspired by
                    <a class="link" href="https://loldle.net/">loldle.net</a> and
                    <a href="https://www.nytimes.com/games/wordle/index.html" class="link">wordle</a>.
                </p>
                <p>
                    <span class="font-extrabold">Background artists</span>
                    <br />
                    <b>- Gen 1-3 Collage black/white bg:</b>
                    u/ZaryXYZ (reddit)
                    <br />
                    <b> - Cubone Background:</b> u/blizard72 (reddit)
                    <br />
                    <b>- Mt. Silver Battle Background: </b>
                    <a class="link" href="https://www.deviantart.com/serain">Serian</a>
                </p>
                <p>Original artist/creator of image(s), please DM for removal or credit.</p>
            </div>
            <hr class="divider" />
            <h1 class="mb-2 text-center text-[32px] uppercase">Contact</h1>
            <p>
                If you have any questions/feedback/bug reports feel free to contact me on one of the following below:
                <br />
                &#8594;
                <a class="link ml-2 text-[24px] font-bold" href="mailto: pokedle.contact@gmail.com">
                    pokedle.contact@gmail.com
                </a>
                <br />
                &#8594;
                <a class="link ml-2 text-[24px] font-bold" href="https://twitter.com/pokedle"> Twitter </a>
            </p>
        </div>
    </BaseDialog>
</template>

<script setup>
import BaseDialog from './BaseDialog.vue';
import SpriteAnimation from '../animations/SpriteAnimation.vue';
const props = defineProps({
    isVisible: Boolean,
    onClose: Function,
});
</script>
