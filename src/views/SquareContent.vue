<template>
  <div class="square-content" :style="{ background: color }">
    <p>{{text}}</p>
  </div>
</template>

<script>
export default {
    props: {
        color: {
            required: true,
            type: String, //TODO: enum for true, false, partially true
        },
        text: {
            required: true,
            type: String,
        }
    },
}
</script>

<style>
.square-content {
    width: 60px;
    height: 60px;;
}

.green {
    background: greenyellow;
}

.red {
    background: red;
}
</style>
