<template>
    <BaseDialog :isVisible="isVisible" :onClose="onClose">
        <div class="relative max-w-[700px] font-pkmEmerald text-[20px]">
            <h1 class="mb-2 text-center text-[32px] uppercase">About</h1>
            <div class="relative flex flex-row items-end">
                <span class="relatives w-fit text-center font-pkmEmerald text-[20px]">
                    {{ getLanguage().infoDialog.intro.welcome }} </span
                ><SpriteAnimation />
            </div>
            <p>
                {{ getLanguage().infoDialog.intro.text }}
            </p>
            <hr class="divider" />
            <h1 class="mb-2 text-center font-pkmEmerald text-[32px] uppercase">How To Play</h1>
            <p class="my-2">
                <b>{{ getLanguage().infoDialog.howToPlay.guessNumbered[1][1] }}</b
                >{{ getLanguage().infoDialog.howToPlay.guessNumbered[1][2]
                }}<b>{{ getLanguage().infoDialog.howToPlay.guessNumbered[1][3] }}</b
                >{{ getLanguage().infoDialog.howToPlay.guessNumbered[1][4]
                }}<b>{{ getLanguage().infoDialog.howToPlay.guessNumbered[1][6] }}</b>
                {{ getLanguage().infoDialog.howToPlay.guessNumbered[1][6] }}
            </p>
            <p class="my-2">
                <b>{{ getLanguage().infoDialog.howToPlay.guessNumbered[2][1] }}</b
                >{{ getLanguage().infoDialog.howToPlay.guessNumbered[2][2]
                }}<b>{{ getLanguage().infoDialog.howToPlay.guessNumbered[2][3] }}</b
                >{{ getLanguage().infoDialog.howToPlay.guessNumbered[2][4]
                }}<b>{{ getLanguage().infoDialog.howToPlay.guessNumbered[2][5] }}</b
                >{{ getLanguage().infoDialog.howToPlay.guessNumbered[2][6]
                }}<b>{{ getLanguage().infoDialog.howToPlay.guessNumbered[2][7] }}</b>
                {{ getLanguage().infoDialog.howToPlay.guessNumbered[2][8] }}<br />
                {{ getLanguage().infoDialog.howToPlay.guessNumbered[2][9]
                }}<b>{{ getLanguage().infoDialog.howToPlay.guessNumbered[2][10] }}</b
                >{{ getLanguage().infoDialog.howToPlay.guessNumbered[2][11] }}
            </p>
            <p class="my-2">
                <b>{{ getLanguage().infoDialog.howToPlay.guessNumbered[3][1] }}</b
                >{{ getLanguage().infoDialog.howToPlay.guessNumbered[3][2]
                }}<b>{{ getLanguage().infoDialog.howToPlay.guessNumbered[3][3] }}</b
                >{{ getLanguage().infoDialog.howToPlay.guessNumbered[3][4]
                }}<b>{{ getLanguage().infoDialog.howToPlay.guessNumbered[3][5] }}</b
                >{{ getLanguage().infoDialog.howToPlay.guessNumbered[3][6] }}<br />
                {{ getLanguage().infoDialog.howToPlay.guessNumbered[3][7] }}
                <b>{{ getLanguage().infoDialog.howToPlay.guessNumbered[3][8] }}.</b>
            </p>
            <p class="my-2">
                <b>{{ getLanguage().infoDialog.howToPlay.guessNumbered[4][1] }}</b
                >{{ getLanguage().infoDialog.howToPlay.guessNumbered[4][2]
                }}<b>{{ getLanguage().infoDialog.howToPlay.guessNumbered[4][3] }}</b
                >{{ getLanguage().infoDialog.howToPlay.guessNumbered[4][4] }}
                <b>{{ getLanguage().infoDialog.howToPlay.guessNumbered[4][5] }}</b> <br />
                {{ getLanguage().infoDialog.howToPlay.guessNumbered[4][6]
                }}<b>{{ getLanguage().infoDialog.howToPlay.guessNumbered[4][7] }}.</b>
            </p>
            <img class="card mx-auto w-[600px]" :src="'./client/assets/game-examples/game-example.png'" />

            <hr class="divider" />
            <h1 class="mb-2 text-center font-pkmEmerald text-[32px] uppercase">
                {{ getLanguage().infoDialog.habitats.habitats }}
            </h1>
            <p>
                {{ getLanguage().infoDialog.habitats.takenFrom }}
                <a class="link" href="https://bulbapedia.bulbagarden.net/wiki/List_of_Pok%C3%A9mon_by_habitat">
                    {{ getLanguage().infoDialog.habitats.bulbapedia }}
                </a>
                {{ getLanguage().infoDialog.habitats.andInclude }}
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
                {{ getLanguage().infoDialog.contribute.contribute }}
            </h1>
            <p class="leading-8">
                {{ getLanguage().infoDialog.contribute.wantToContribute }}
                <br />
                {{ getLanguage().infoDialog.contribute.helpNeeded }}
            </p>
            <ul class="ml-5 list-disc">
                <li>{{ getLanguage().infoDialog.contribute.animations }}</li>
                <li>{{ getLanguage().infoDialog.contribute.translation }}</li>
                <li>{{ getLanguage().infoDialog.contribute.styling }}</li>
                <li>{{ getLanguage().infoDialog.contribute.features }}</li>
            </ul>
            <hr class="divider" />
            <h1 class="mb-2 text-center font-pkmEmerald text-[32px] uppercase">
                {{ getLanguage().infoDialog.credit.credit }}
            </h1>
            <div class="w-full leading-8">
                <p class="mb-2">{{ getLanguage().infoDialog.credit.intro }}</p>
                <p class="mb-2">
                    {{ getLanguage().infoDialog.credit.intro }}
                    <a class="link" href="https://loldle.net/">{{ getLanguage().infoDialog.credit.loldle }}</a>
                    {{ getLanguage().infoDialog.credit.and }}
                    <a href="https://www.nytimes.com/games/wordle/index.html" class="link">{{
                        getLanguage().infoDialog.credit.wordle
                    }}</a
                    >.
                </p>
                <p>
                    <span class="font-extrabold">{{ getLanguage().infoDialog.credit.backgroundArtists }}</span>
                    <br />
                    <b>{{ getLanguage().infoDialog.credit.pokemonCollageBG }}</b>
                    {{ getLanguage().infoDialog.credit.pokemonCollageBGCreator }}
                    <br />
                    <b>{{ getLanguage().infoDialog.credit.cuboneBG }}</b
                    >{{ getLanguage().infoDialog.credit.cuboneBGCreator }}
                    <br />
                    <b>{{ getLanguage().infoDialog.credit.mtSilverBG }}</b>
                    <a class="link" href="https://www.deviantart.com/serain">{{
                        getLanguage().infoDialog.credit.mtSilverBGCreator
                    }}</a>
                </p>
                <p>{{ getLanguage().infoDialog.credit.contactForCreditOrRemoval }}</p>
            </div>
            <hr class="divider" />
            <h1 class="mb-2 text-center text-[32px] uppercase">{{ getLanguage().infoDialog.contact }}</h1>
            <p>
                {{ getLanguage().infoDialog.contact.intro }}
                <br />
                &#8594;
                <a class="link ml-2 text-[24px] font-bold" href="mailto: pokedle.contact@gmail.com">
                    {{ getLanguage().infoDialog.contact.mail }}
                </a>
                <br />
                &#8594;
                <a class="link ml-2 text-[24px] font-bold" href="https://twitter.com/pokedle">
                    {{ getLanguage().infoDialog.contact.twitter }}
                </a>
            </p>
        </div>
    </BaseDialog>
</template>

<script setup>
import BaseDialog from './BaseDialog.vue';
import SpriteAnimation from '../animations/SpriteAnimation.vue';
import { getLanguage } from '../../services/language';
const props = defineProps({
    isVisible: Boolean,
    onClose: Function,
});
</script>
