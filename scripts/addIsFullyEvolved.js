import { createRequire } from "module"; // Bring in the ability to create the 'require' method
const require = createRequire(import.meta.url); // construct the require method
const pokemonData = require('../src/data/pokemonData-v1.json');
import fs from 'fs';

const saveAsJsonTest = async () => {
  let lastElement = pokemonData[0];
  lastElement.isFullyEvolved = false;

  for (let i = 1; i < pokemonData.length; i++) {
    const currElement = pokemonData[i];
    if (lastElement.evolutionState >= currElement.evolutionState) {
      lastElement.isFullyEvolved = true;
    } else {
      lastElement.isFullyEvolved = false;
    }
    lastElement = currElement;
  }
  const jsonData = JSON.stringify(pokemonData);

  fs.writeFile("./src/data/pokemon-v2.json", jsonData, (err) => {
    if (err) {
      console.log("Error");
      console.log(err);
    }
  })
}

await saveAsJsonTest();
