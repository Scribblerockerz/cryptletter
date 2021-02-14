<template>
    <div
        :class="{
            'letter--striped': striped,
            'letter--spaced': spaced,
        }"
        class="letter"
    >
        <div v-if="$slots.header" class="letter__header">
            <slot name="header" />
        </div>
        <div class="letter__body"><slot /></div>
        <div
            v-if="$slots.footer"
            :class="{ 'letter__footer--separator': footSeparator }"
            class="letter__footer"
        >
            <slot name="footer" />
        </div>
    </div>
</template>

<script>
export default {
    name: "Letter",
    props: {
        striped: {
            type: Boolean,
            default: false,
        },
        spaced: {
            type: Boolean,
            default: false,
        },
        footSeparator: {
            type: Boolean,
            default: false,
        },
    },
};
</script>

<style scoped>
.letter {
    position: relative;
    background-color: #fff;
    color: #000;
    border-radius: 4px;
    padding: 30px 50px;
}

.letter--striped {
    background-image: url("../assets/stripe.svg");
    background-position: right top;
    background-repeat: repeat-y;
    background-size: 20px;
}

.letter--spaced {
    margin-bottom: 40px;
}

.letter__header {
    position: absolute;
    top: 20px;
    right: 50px;
    left: 50px;
    text-align: right;
}

.letter__body {
    padding: 20px 0;
}

.letter__footer {
    padding-top: 30px;
}

.letter__footer--separator {
    border-top: 2px solid #efefef;
    margin-top: 10px;
}
</style>
