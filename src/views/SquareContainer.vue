<template>
    <div class="content">
        <SquareContent :color="'green'" :text="'fire'" />
        <SquareContent :color="'red'" :text="'water'" />
    </div>
</template>

<script>
import SquareContent from './SquareContent.vue';

export default {
    components:{
        SquareContent,
    }
}

</script>

<style>

.content {
    display: flex;
    gap: 10px;
}

</style>