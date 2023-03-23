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
    Pokemon: 'Pokemon',
    Type1: 'Type 1',
    Type2: 'Type 2',
    EvolutionLevel: 'Evolution Level',
    Evolutions: 'Fully Evolved',
    Color: 'Color',
    Habitat: 'Habitat',
    Generation: 'Generation',
};

const GuessFieldTooltips = {
    Pokemon: '',
    Type1: 'Primary Type',
    Type2: 'Secondary Type',
    EvolutionLevel: 'Current Evolution Stage; 1, 2 or 3',
    Evolutions: 'Yes or No',
    Color: 'Red, Green, Blue etc.',
    Habitat: 'Sea, Cave, Grasslands etc.',
    Generation: 'Generation 1, 2 or 3',
};

const getGuessFieldTooltipsFromIndex = (index) => {
    const values = Object.values(GuessFieldTooltips);
    return values.at(index);
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
const FlavortextGuessesNeededForHintTwo = 4;
const FlavortextGuessesNeededForHintThree = 6;

const SilhouetteGuessesNeededForHintOne = 2;
const SilhouetteGuessesNeededForHintTwo = 3;
const SilhouetteGuessesNeededForHintThree = 4;

const GameModes = {
    Classic: 'Classic',
    Flavortext: 'Flavortext',
    Silhouette: 'Silhouette',
};

export {
    GuessState,
    GuessFieldTitles,
    Habitats,
    GuessFieldTooltips,
    GuessType,
    TotalResultCardFlipDelay,
    getGuessFieldTooltipsFromIndex,
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
    SilhouetteGuessesNeededForHintThree,
    GameModes,
};
