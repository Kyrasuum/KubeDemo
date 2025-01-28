<script setup lang="ts">
import { useQuasar } from "quasar";

export const call = (f:any) => typeof f === 'function'? f(): f;

export const parallel = (...fns:any[])  => () => Promise.all(fns.map(call));

export const series = (...fns:any[]) => async () => {
	let res = [];

	for (let f of fns)
		res.push(await call(f));

	return res;
};

export function until(conditionFunction: any) {
	const poll = (resolve: any) => {
		if(conditionFunction()) resolve();
		else setTimeout(_ => poll(resolve), 400);
	}
	return new Promise(poll);
}

	
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
