<template>
    <button
        :disabled="disabled"
        :class="{
            'btn--primary': primary,
            'btn--nano': nano,
        }"
        class="btn"
    >
        <slot />
    </button>
</template>

<script>
export default {
    name: "Button",
    props: {
        primary: {
            type: Boolean,
            default: false,
        },
        nano: {
            type: Boolean,
            default: false,
        },
        disabled: {
            type: Boolean,
            default: false,
        },
    },
};
</script>

<style scoped>
.btn {
    font-size: 16px;
    font-family: Georgia, serif;
    padding: 11px 30px 12px;
}

.btn:disabled {
    opacity: 0.6;
    cursor: not-allowed;
}

.btn--primary {
    color: #fff;
    background-color: #24272b;
    border: 1px solid #000000;
    box-shadow: 0 1px 1px 0 #000000, 0 2px 3px 0 rgba(0, 0, 0, 0.32);
    transition: background-color 0.2s ease-in-out, color 0.2s ease-in-out;
}

.btn--primary:hover {
    /* TODO: rewrite to proper hsl brightness change */
    /*background-color: lighten(#24272b, 6%);*/
    filter: brightness(120%);
}

.btn--primary:active {
    /* TODO: rewrite to proper hsl brightness change */
    /*color: darken(#fff, 20%);*/
    /*background-color: darken(#24272b, 2%);*/
    filter: brightness(80%);
}

.btn--primary:disabled {
    background-color: #24272b;
}

.btn--nano {
    font-family: "system-ui", monospace;
    font-size: 11px;
    color: #aaa;
    border: 1px solid transparent;
    background-color: transparent;
    text-transform: uppercase;
    padding: 2px 4px;
}

.btn--nano:hover {
    color: #222;
    border-color: #222;
}

.btn--nano:disabled {
    border-color: transparent;
    color: #aaa;
}
</style>
