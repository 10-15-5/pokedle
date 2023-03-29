import { Habitats } from '../constants';

const getHabitatImage = (habitat) => {
    switch (habitat) {
        case Habitats.Cave:
            return '/client/assets/habitats/cave.png';
        case Habitats.Forest:
            return '/client/assets/habitats/forest.png';
        case Habitats.Grassland:
            return '/client/assets/habitats/grassland.png';
        case Habitats.Mountain:
            return '/client/assets/habitats/mountain.png';
        case Habitats.Rare:
            return '/client/assets/habitats/rare.png';
        case Habitats.RoughTerrain:
            return '/client/assets/habitats/rough-terrain.png';
        case Habitats.Sea:
            return '/client/assets/habitats/sea.png';
        case Habitats.Urban:
            return '/client/assets/habitats/urban.png';
        case Habitats.WatersEdge:
            return '/client/assets/habitats/waters-edge.png';
    }
    return '';
};

const getRandomPokemonAnimation = () => {
    const random = Math.floor(Math.random() * 8);

    switch (random) {
        case 1:
            return '/client/assets/moving-sprites/charmander.png';
        case 2:
            return '/client/assets/moving-sprites/elekid.png';
        case 3:
            return '/client/assets/moving-sprites/gulpin.png';
        case 4:
            return '/client/assets/moving-sprites/mudkip.png';
        case 5:
            return '/client/assets/moving-sprites/ponytar.png';
        case 6:
            return '/client/assets/moving-sprites/ralts.png';
        case 7:
            return '/client/assets/moving-sprites/squirdle.png';
        case 8:
            return '/client/assets/moving-sprites/torchic.png';
        default:
            return '/client/assets/moving-sprites/whooper.png';
    }
};

export { getHabitatImage, getRandomPokemonAnimation };
