<script setup lang="ts">
import { useQuasar } from "quasar";

const q = useQuasar();
let init = 0;
if (init < 1) {
	q.dark.toggle();
	init++;
}

function toggleDarkMode() {
	q.dark.toggle();
}
</script>

<template>
	<router-view />
</template>

<style scoped></style>
