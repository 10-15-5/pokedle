import moment from 'moment-timezone';

const guessState = {
    CorrectGuess: 'CorrectGuess',
    PartlyCorrectGuess: 'PartlyCorrectGuess',
    WrongGuess: 'WrongGuess',
    None: 'None',
};

const guessType = {
    Pokemon: 'Pokemon',
    Text: 'Text',
    Habitat: 'Habitat',
    Blackout: 'Blackout',
};

const guessFieldTitles = {
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

const DateOfFirstPokeldeGameClassic = moment('Marts 16, 2023 00:00:00 UTC');

const ClassicGuessesNeededForHintOne = 4;
const ClassicGuessesNeededForHintTwo = 6;
const ClassicGuessesNeededForHintThree = 8;

const GameModes = {
    Classic: 'Classic',
    Flavortext: 'Flavortext'
}

export {
    guessState,
    guessFieldTitles,
    Habitats,
    GuessFieldTooltips,
    guessType,
    TotalResultCardFlipDelay,
    getGuessFieldTooltipsFromIndex,
    DateOfFirstPokeldeGameClassic,
    ClassicGuessesNeededForHintOne,
    ClassicGuessesNeededForHintTwo,
    ClassicGuessesNeededForHintThree,
    GameModes
};
