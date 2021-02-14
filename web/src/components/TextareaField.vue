<template>
    <textarea
        :value="modelValue"
        :rows="rows"
        :autofocus="autofocus"
        :readonly="readonly"
        :placeholder="placeholder"
        :class="{
            'input--drag-down': dragDown,
            'input--raw': raw,
            'input--url-holder': urlHolder,
        }"
        class="input input--textarea"
        autocomplete="off"
        @input="handleInput"
    />
</template>

<script>
export default {
    name: "TextareaField",
    props: {
        modelValue: {
            type: String,
            required: true,
        },
        placeholder: {
            type: String,
            default: null,
        },
        rows: {
            type: Number,
            default: 3,
        },
        // TODO: refactor name
        dragDown: {
            type: Boolean,
            default: false,
        },
        autofocus: {
            type: Boolean,
            default: false,
        },
        readonly: {
            type: Boolean,
            default: false,
        },
        raw: {
            type: Boolean,
            default: false,
        },
        // TODO: refactor name
        urlHolder: {
            type: Boolean,
            default: false,
        },
    },
    emits: ["update:modelValue"],
    methods: {
        handleInput(e) {
            this.$emit("update:modelValue", e.target.value);
        },
    },
};
</script>

<style scoped>
.input {
    font-family: "Courier New", monospace;
    width: 100%;
    display: block;
    border: 0;
    font-size: 14px;
    outline: none;
}

.input--textarea {
    padding-bottom: 30px;
    resize: vertical;
}

.input--raw {
    padding: 31px 30px 29px;
    border: 3px solid #eee;
}

.input--drag-down {
    margin-bottom: -30px;
}

.input--url-holder {
    font-size: 13px;
    padding: 9px 20px 8px 0;
    border-bottom: 2px solid #b6d8ff;
    text-overflow: ellipsis;
}
</style>
