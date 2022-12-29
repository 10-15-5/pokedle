/* eslint-disable no-undef */
import express from "express";
import * as dotenv from "dotenv";

import { someResourceRouter } from "./routes/some-resource/index.js";

dotenv.config();

const app = express();

// For parsing application/json
app.use(express.json())

app.get('/', (req, res) => {
    res.status(200);
    res.send("Welcome to root URL of Server");
});

app.use('/api',
[
    someResourceRouter,
    //Add routes here
])

app.listen(process.env.EXPRESS_PORT, (error) => {
    if (error) {
        console.log("Error occurred, server can't start", error);
    }
    console.log(`Server is Successfully Running on Port ${process.env.EXPRESS_PORT}`);
});