import { createRequire } from "module"; // Bring in the ability to create the 'require' method
const require = createRequire(import.meta.url); // construct the require method
const pokemonData = require('../../server/data/pokemonData-v4.json');
import fs from 'fs';
import axios from "axios";

const runScript = async () => {
    
    const updatedPokemonData = pokemonData.map((d) => ({
        number: d.number,
        name: d.name,
        type1: d.type1.toLowerCase(),
        type2: d.type2.toLowerCase(),
        generation: d.generation,
        legendary: d.legendary.toLowerCase(),
        evolutionState: d.evolutionState,
        isFullyEvolved: d.isFullyEvolved.toLowerCase(),
        habitat: d.habitat.toLowerCase(),
        color: d.color.toLowerCase()        
    }))

    const jsonData = JSON.stringify(updatedPokemonData);

    fs.writeFile("./scripts/result.json", jsonData, (err) => {
        if (err) {
            console.log("Error");
            console.log(err);
        }
    })
}

await runScript();