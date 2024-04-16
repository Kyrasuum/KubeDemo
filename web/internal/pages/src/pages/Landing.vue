<script setup lang="ts">
import layout from 'src/components/Layout.vue';

let pathname = window.location.pathname;
let index = pathname.lastIndexOf("/");
let host_url = pathname.substring(0, index);
</script>

<template>
    <layout>
        <template #default>
            <iframe title="Message of the Day" style="background-color: white;" :src="host_url + '/api/test.json'" width="100%" height="100%" type="text/json"></iframe>
        </template>
    </layout>
</template>

<style scoped lang="scss">
</style>
