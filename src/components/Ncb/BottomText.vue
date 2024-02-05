<template>
    <div>
        <v-flex xl8>
            <p class="text text--secondary" >{{ disclaimer }}</p>
        </v-flex>
    </div>
</template>

<script>
export default {
    props: {
        disclaimer: String
    }
}
</script>

<style scoped>
.text {
    font-size: .8em;
}
</style>