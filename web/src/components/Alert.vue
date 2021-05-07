<template>
    <div :class="`alert--${type}`" class="alert alert--hidden"><slot /></div>
</template>

<script>
export default {
    name: "Alert",
    props: {
        type: {
            type: String,
            default: "info",
            validator: (v) => ["info", "error"].includes(v),
        },
    },
};
</script>

<style scoped>
.alert {
    position: relative;
    background-color: #fff;
    color: #000;
    padding: 20px 30px 18px;
    text-align: center;
    font-family: Georgia, serif;
    font-size: 18px;
    -webkit-box-shadow: 0 1px 1px 0 #000000, 0 2px 3px 0 rgba(0, 0, 0, 0.32);
    box-shadow: 0 1px 1px 0 #000000, 0 2px 3px 0 rgba(0, 0, 0, 0.32);
    width: 80%;
    margin-left: auto;
    margin-right: auto;
}

.alert--error {
    background-color: #ff5454;
    border-bottom: 5px solid #ff5454;
    margin-bottom: -5px;
    color: #fff;
}
</style>
