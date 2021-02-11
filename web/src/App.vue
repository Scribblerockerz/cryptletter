<template>
    <img alt="Vue logo" class="logo" src="./assets/logo.svg" />
    <LetterNew />
</template>

<script>
import "./assets/global.css";

import LetterNew from "./pages/LetterNew";

export default {
    name: "App",

    components: {
        LetterNew,
    },
};
</script>

<style>
#app {
    font-family: Avenir, Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    text-align: center;
    color: #2c3e50;
    margin-top: 60px;
}
</style>
