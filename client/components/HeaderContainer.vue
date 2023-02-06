<template>
    <div class="container">
        <v-btn class="icon"
               icon="mdi-fire-circle"
               size="large"></v-btn>
        <v-btn class="icon"
               icon="mdi-cog"
               size="large"></v-btn>
        <v-btn class="icon"
               icon="mdi-poll"
               size="large"></v-btn>
        <v-btn class="icon"
               icon="mdi-information-outline"
               size="large"></v-btn>
    </div>
</template>

<script>
</script>

<style scoped>
.icon {
    width: 32px;
    height: 32px;
    margin: 4px;
}

.container {
    display: flex;
    flex-direction: row;
    justify-content: center;
    align-items: center;
}
</style>