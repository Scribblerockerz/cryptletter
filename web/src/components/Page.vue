<template>
    <div class="page">
        <div class="page__header">
            <h1 class="page__title">
                <a class="logo" href="/">Cryptletter</a>
            </h1>
            <div class="page__slogan">
                Encrypted self-destructing messages.<br />
                Even Cruise would like to have it!
            </div>
        </div>
        <div class="page__container"><slot /></div>
    </div>
</template>

<script>
export default {
    name: "Page",
};
</script>

<style scoped>
.page {
    margin-bottom: 70px;
}

.page__header {
    text-align: center;
    line-height: 1.2em;
    margin: 70px auto;
}

.page__container {
    margin: auto;
    max-width: 840px;
    padding: 0 20px;
}
</style>
