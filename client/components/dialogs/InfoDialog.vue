<template>
    <BaseDialog :isVisible="isVisible" :onClose="onClose">
        <div class="relative max-w-[700px] font-pkmEmerald text-[20px]">
            <h1 class="mb-2 text-center text-[32px] uppercase">About</h1>
            <div class="relative flex flex-row items-end">
                <span class="relatives w-fit text-center font-pkmEmerald text-[20px]">
                    {{ text().infoDialog.intro.welcome }} </span
                ><SpriteAnimation />
            </div>
            <p>
                {{ text().infoDialog.intro.text }}
            </p>
            <hr class="divider" />
            <h1 class="mb-2 text-center font-pkmEmerald text-[32px] uppercase">How To Play</h1>
            <p class="my-2">
                <b>{{ text().infoDialog.howToPlay.guesses[1].title }}</b
                >{{ text().infoDialog.howToPlay.guesses[1].text
                }}
            </p>
            <p class="my-2">
                <b>{{ text().infoDialog.howToPlay.guesses[2].title }}</b
                >{{ text().infoDialog.howToPlay.guesses[2].text
                }}
            </p>
            <p class="my-2">
                <b>{{ text().infoDialog.howToPlay.guesses[3].title }}</b
                >{{ text().infoDialog.howToPlay.guesses[3].text
                }}
            </p>
            <p class="my-2">
                <b>{{ text().infoDialog.howToPlay.guesses[4].title }}</b
                >{{ text().infoDialog.howToPlay.guesses[4].text
                }}
            </p>
            <img class="card mx-auto w-[600px]" :src="'./client/assets/game-examples/game-example.png'" />

            <hr class="divider" />
            <h1 class="mb-2 text-center font-pkmEmerald text-[32px] uppercase">
                {{ text().infoDialog.habitats.habitats }}
            </h1>
            <p>
                {{ text().infoDialog.habitats.sourcedFrom }}
                <a class="link" href="https://bulbapedia.bulbagarden.net/wiki/List_of_Pok%C3%A9mon_by_habitat">
                    Bulbapedia
                </a>
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
            <h1 class="mb-2 text-center font-pkmEmerald text-[32px] uppercase">
                {{ text().infoDialog.contribute.contribute }}
            </h1>
            <p class="leading-8">
                {{ text().infoDialog.contribute.wantToContribute }}
                <br />
                {{ text().infoDialog.contribute.helpNeeded }}
            </p>
            <ul class="ml-5 list-disc">
                <li>{{ text().infoDialog.contribute.animations }}</li>
                <li>{{ text().infoDialog.contribute.translation }}</li>
                <li>{{ text().infoDialog.contribute.styling }}</li>
                <li>{{ text().infoDialog.contribute.features }}</li>
            </ul>
            <hr class="divider" />
            <h1 class="mb-2 text-center font-pkmEmerald text-[32px] uppercase">
                {{ text().infoDialog.credit.credit }}
            </h1>
            <div class="w-full leading-8">
                <p class="mb-2">{{ text().infoDialog.credit.intro }}</p>
                <p class="mb-2">
                    {{ text().infoDialog.credit.inspiredBy }}
                    <a class="link" href="https://loldle.net/">loldle.net</a>,
                    <a href="https://www.nytimes.com/games/wordle/index.html" class="link">wordle</a
                    >.
                </p>
                <p>
                    <span class="font-extrabold">{{ text().infoDialog.credit.backgroundArtists }}</span>
                    <br />
                    <b>{{ text().infoDialog.credit.pokemonCollageBG }}</b>
                    {{ text().infoDialog.credit.pokemonCollageBGCreator }}
                    <br />
                    <b>{{ text().infoDialog.credit.cuboneBG }}</b
                    >{{ text().infoDialog.credit.cuboneBGCreator }}
                    <br />
                    <b>{{ text().infoDialog.credit.mtSilverBG }}</b>
                    <a class="link" href="https://www.deviantart.com/serain">{{
                        text().infoDialog.credit.mtSilverBGCreator
                    }}</a>
                </p>
                <p>{{ text().infoDialog.credit.contactForCreditOrRemoval }}</p>
            </div>
            <hr class="divider" />
            <h1 class="mb-2 text-center text-[32px] uppercase">{{ text().infoDialog.contact.contact }}</h1>
            <p>
                {{ text().infoDialog.contact.intro }}
                <br />
                &#8594;
                <a class="link ml-2 text-[24px] font-bold" href="mailto: pokedle.contact@gmail.com">
                    pokedle.contact@gmail.com
                </a>
                <br />
                &#8594;
                <a class="link ml-2 text-[24px] font-bold" href="https://twitter.com/pokedle">
                    {{ text().infoDialog.contact.twitter }}
                </a>
            </p>
        </div>
    </BaseDialog>
</template>

<script setup>
import BaseDialog from './BaseDialog.vue';
import SpriteAnimation from '../animations/SpriteAnimation.vue';
import { text } from '../../services/language';
const props = defineProps({
    isVisible: Boolean,
    onClose: Function,
});
</script>
