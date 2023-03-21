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

export { getHabitatImage };
