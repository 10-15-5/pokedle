import moment from 'moment-timezone';

const GuessState = {
    CorrectGuess: 'CorrectGuess',
    PartlyCorrectGuess: 'PartlyCorrectGuess',
    WrongGuess: 'WrongGuess',
    None: 'None',
};

const GuessType = {
    Pokemon: 'Pokemon',
    Text: 'Text',
    Habitat: 'Habitat',
    Blackout: 'Blackout',
};

const GuessFieldTitles = {
    Pokemon: 'pokemon',
    Type1: 'type1',
    Type2: 'type2',
    EvolutionLevel: 'evolutionLevel',
    Evolutions: 'evolutions',
    Color: 'color',
    Habitat: 'habitat',
    Generation: 'generation',
};

const Habitats = {
    WatersEdge: 'waters-edge',
    Cave: 'cave',
    Rare: 'rare',
    Sea: 'sea',
    Urban: 'urban',
    Grassland: 'grassland',
    Mountain: 'mountain',
    Forest: 'forest',
    RoughTerrain: 'rough-terrain',
};

const TotalResultCardFlipDelay = 2750;

const DateOfFirstPokeldeGameClassic = moment.utc('16/03/2023', 'DD/MM/YYYY');
const DateOfFirstPokeldeGameFlavortext = moment.utc('27/03/2023', 'DD/MM/YYYY');
const DateOfFirstPokeldeGameSilhouette = moment.utc('27/03/2023', 'DD/MM/YYYY');

const ClassicGuessesNeededForHintOne = 4;
const ClassicGuessesNeededForHintTwo = 6;
const ClassicGuessesNeededForHintThree = 8;

const FlavortextGuessesNeededForHintOne = 2;
const FlavortextGuessesNeededForHintTwo = 3;
const FlavortextGuessesNeededForHintThree = 4;

const SilhouetteGuessesNeededForHintOne = 1;
const SilhouetteGuessesNeededForHintTwo = 2;

const GameModes = {
    Classic: 'Classic',
    Flavortext: 'Flavortext',
    Silhouette: 'Silhouette',
};

export {
    GuessState,
    GuessFieldTitles,
    Habitats,
    GuessType,
    TotalResultCardFlipDelay,
    DateOfFirstPokeldeGameClassic,
    DateOfFirstPokeldeGameFlavortext,
    ClassicGuessesNeededForHintOne,
    ClassicGuessesNeededForHintTwo,
    ClassicGuessesNeededForHintThree,
    FlavortextGuessesNeededForHintOne,
    FlavortextGuessesNeededForHintTwo,
    FlavortextGuessesNeededForHintThree,
    DateOfFirstPokeldeGameSilhouette,
    SilhouetteGuessesNeededForHintOne,
    SilhouetteGuessesNeededForHintTwo,
    GameModes,
};
