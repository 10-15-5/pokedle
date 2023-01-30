import { Habitats } from "../constants"

const getHabitatImage = (habitat) => {
    switch (habitat) {
        case Habitats.Cave:
            return "/client/assets/habitats/cave-habitat.png"
        case Habitats.Forest:
            return "/client/assets/habitats/forest-habitat.png"
        case Habitats.Grassland:
            return "/client/assets/habitats/grassland-habitat.png"
        case Habitats.Mountain:
            return "/client/assets/habitats/mountain-habitat.png"
        case Habitats.Rare:
            return "/client/assets/habitats/rare-habitat.png"
        case Habitats.RoughTerrain:
            return "/client/assets/habitats/rough-terrain-habitat.png"
        case Habitats.Sea:
            return "/client/assets/habitats/sea-habitat.png"
        case Habitats.Urban:
            return "/client/assets/habitats/urban-habitat.png"
        case Habitats.WatersEdge:
            return "/client/assets/habitats/waters-edge-habitat.png"
    }
    return "";
}

export {
    getHabitatImage
}