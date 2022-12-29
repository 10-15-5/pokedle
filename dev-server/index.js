/* eslint-disable no-undef */
import express from "express";
import * as dotenv from "dotenv";
dotenv.config();

const app = express();

app.listen(process.env.EXPRESS_PORT, (error) => {
    if (!error) {
        console.log(`Server is Successfully Running on Port ${process.env.EXPRESS_PORT}`);
    } else {
        console.log("Error occurred, server can't start", error);
    }
}
);